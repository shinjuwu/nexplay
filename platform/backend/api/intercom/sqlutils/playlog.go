package sqlutils

import (
	"backend/api/intercom/model"
	"backend/internal/ginweb"
	"backend/pkg/redis"
	"backend/pkg/utils"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func CreateGamePlayLog(db *sql.DB, redisClient redis.IRedisCliect, logger ginweb.ILogger, playLogTableName string, req model.PlayLogRequest, betTime, startTime, endTime time.Time) error {
	now := utils.TimeNowUTC()

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert(playLogTableName).
		Columns("lognumber", "game_id", "room_type", "desk_id", "exchange",
			"playlog", "de_score", "ya_score", "valid_score", "create_time",
			"is_big_win", "is_issue", "bet_time", "tax", "bonus",
			"start_time", "end_time").
		Values(req.Lognumber, req.GameId, req.RoomType, req.DeskId, req.Exchange,
			req.Playlog, req.DeScore, req.YaScore, req.ValidScore, now,
			req.IsBigWin, req.IsIssue, betTime, req.Tax, req.Bonus,
			startTime, endTime).
		ToSql()
	result, err := db.Exec(query, args...)
	if err != nil {
		logger.Printf("CreateGamePlayLog: db exec failed, query = %v, err = %v", query, err)
		return err
	}

	if count, err := result.RowsAffected(); err == nil && count <= 0 {
		logger.Printf("CreateGamePlayLog: db exec result has no rows be change, query = %v", query)
	}
	return nil
}
