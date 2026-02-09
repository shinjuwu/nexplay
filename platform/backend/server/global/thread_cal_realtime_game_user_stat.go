package global

import (
	"backend/internal/notification"
	"backend/internal/statistical"
	"backend/pkg/logger"
	"backend/pkg/redis"
	"backend/pkg/utils"
	table_model "backend/server/table/model"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
)

var (
	execRealtimeGameUserStatTimeForCheck = 60 * time.Second
)

func T_CalRealtimeGameUserStat(logger *logger.RuntimeGoLogger, db *sql.DB, rdb redis.IRedisCliect) {
	c := time.NewTicker(execRealtimeGameUserStatTimeForCheck)
	for {
		_execRealtimeGameUserStat(logger, db, rdb)
		<-c.C
	}
}

func _execRealtimeGameUserStat(logger *logger.RuntimeGoLogger, db *sql.DB, rdb redis.IRedisCliect) {
	// 固定檢測時間
	now := time.Now()
	ftTimeNow := now.UnixNano() / int64(time.Millisecond)

	logger.Info("_execRealtimeGameUserStat run, ftTimeNow: %v", ftTimeNow)

	// default index is lobby
	// lobby idx: 0
	gg := GameCache.Get(SERVER_INFO_DEFAULT)

	conninfo := make(map[string]interface{}, 0)
	addr, ok := ServerInfoCache.Load(gg.ServerInfoCode)
	if ok {
		addres, ok := addr.(table_model.ServerInfo)
		if ok {
			conninfo = utils.ToMap(addres.AddressesBytes)
		}

		// check game server is life first.
		success, _ := notification.SendNotifyToPing(conninfo)
		if !success {
			return
		}
	}

	gameUsersInCache, err := rdb.LoadHAllValue(REDIS_IDX_LOGIN_INFO, REDIS_HASH_INGAME_USER)
	if err != nil {
		logger.Error("_execRealtimeGameUserStat get in game users from redis fail, err: %v", err)
		return
	}

	if len(gameUsersInCache) <= 0 {
		return
	}

	gameUserIds := make([]int, len(gameUsersInCache))
	for gameUserIdStr := range gameUsersInCache {
		gameUserIds = append(gameUserIds, utils.ToInt(gameUserIdStr))
	}

	logtime := utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_MINUTE, now)

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert("agent_game_users_stat_min").
		Columns("log_time", "agent_id", "level_code", "online_game_user_count").
		Select(
			sq.Select(logtime, "gs.agent_id", "a.level_code", `COUNT(gs."id")`).
				From("game_users AS gs").
				InnerJoin(`agent AS a ON a."id" = gs."agent_id"`).
				Where(sq.Eq{`gs."id"`: gameUserIds}).
				GroupBy("gs.agent_id", "a.level_code"),
		).
		ToSql()

	_, err = db.Exec(query, args...)
	if err != nil {
		logger.Error("_execRealtimeGameUserStat insert stat to db fail, game user ids: %v,err: %v", gameUserIds, err)
		return
	}
}
