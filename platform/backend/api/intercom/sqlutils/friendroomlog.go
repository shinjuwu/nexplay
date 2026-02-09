package sqlutils

import (
	"backend/api/intercom/model"
	"backend/pkg/utils"
	"backend/server/global"
	"database/sql"
	"definition"
	"fmt"
	"strconv"

	sq "github.com/Masterminds/squirrel"
)

func CreateFriendRoomLog(db *sql.DB, req model.FriendRoomLogRequest) error {
	var err error

	createTimeI64 := strconv.FormatInt(req.CreateTime, 10)
	createTime := utils.ToTimeUnixMilliUTC(createTimeI64, 0)

	endTimeTimeI64 := strconv.FormatInt(req.EndTime, 10)
	endTime := utils.ToTimeUnixMilliUTC(endTimeTimeI64, 0)

	agentId := -1
	levelCode := ""

	agent := global.AgentCache.Get(req.AgentId)
	if agent != nil {
		agentId = agent.Id
		levelCode = agent.LevelCode
	}

	// roomId, agentId, gameId, userId, username, createTime.UnixMilli()
	salt := fmt.Sprintf("%s_%d_%d_%d_%s_%d", req.RoomId, agentId, req.GameId, req.UserId, req.Username, createTime.UnixMilli())
	id := utils.CreatreOrderIdByOrderTypeAndSalt(definition.ORDER_TYPE_FRIEND_ROOM_LOG, salt, createTime)

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert("friend_room_log").
		Columns(
			"id", "agent_id", "level_code", "game_id", "room_id",
			"user_id", "username", "tax", "taxpercent", "detail",
			"create_time", "end_time",
		).
		Values(
			id, agentId, levelCode, req.GameId, req.RoomId,
			req.UserId, req.Username, req.Tax, req.TaxPercentage, req.Detail,
			createTime, endTime,
		).
		ToSql()

	_, err = db.Exec(query, args...)

	return err
}
