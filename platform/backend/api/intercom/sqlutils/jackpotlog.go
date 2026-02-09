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

func CreateJackpotTokenLog(db *sql.DB, jackpotTokenList []*model.JackpotTokenLog) error {
	var err error

	builder := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert("jackpot_token_log").
		Columns(
			"token_id", "agent_id", "level_code", "user_id", "username",
			"jp_bet", "token_create_time", "source_game_id", "source_lognumber", "source_bet_id",
			"status",
		)

	for _, jpToken := range jackpotTokenList {
		builder = builder.Values(jpToken.TokenId, jpToken.AgentId, jpToken.LevelCode, jpToken.UserId, jpToken.Username,
			jpToken.JpBet, jpToken.TokenCreateTime, jpToken.SourceGameId, jpToken.SourceLognumber, jpToken.SourceBetId,
			definition.JACKPOT_TOKEN_LOG_STATUS_SUCCESS)
	}

	query, args, _ := builder.ToSql()

	_, err = db.Exec(query, args...)

	return err
}

func CreateJackpotLog(db *sql.DB, req model.JackpotLogRequest) error {
	var err error

	tokenCreateTimeI64 := strconv.FormatInt(req.TokenCreateTime, 10)
	tokenCreateTime := utils.ToTimeUnixMilliUTC(tokenCreateTimeI64, 0)

	winningTimeI64 := strconv.FormatInt(req.WinningTime, 10)
	winningTime := utils.ToTimeUnixMilliUTC(winningTimeI64, 0)

	agentId := -1
	levelCode := ""

	agent := global.AgentCache.Get(req.AgentId)
	if agent != nil {
		agentId = agent.Id
		levelCode = agent.LevelCode
	}

	// lognumber, agentId, username, jpBet, prizeScore, winningTime.UnixMilli()
	salt := fmt.Sprintf("%s_%d_%s_%f_%f_%d", req.Lognumber, agentId, req.Username, req.JpBet, req.PrizeScore, winningTime.UnixMilli())
	betId := utils.CreatreOrderIdByOrderTypeAndSalt(definition.ORDER_TYPE_JACKPOT_LOG, salt, winningTime)

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert("jackpot_log").
		Columns(
			"bet_id", "lognumber", "token_id", "agent_id", "level_code",
			"user_id", "username", "jp_bet", "token_create_time", "prize_score",
			"prize_item", "winning_time", "show_pool", "real_pool", "is_robot",
		).
		Values(
			betId, req.Lognumber, req.TokenId, agentId, levelCode,
			req.UserId, req.Username, req.JpBet, tokenCreateTime, req.PrizeScore,
			req.PrizeItem, winningTime, req.ShowPool, req.RealPool, req.IsRobot,
		).
		ToSql()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}

	query, args, _ = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("jackpot_token_log").
		SetMap(sq.Eq{
			"usage_count": sq.Expr("usage_count + 1"),
			"update_time": sq.Expr("now()"),
		}).
		Where(sq.Eq{"token_id": req.TokenId}).
		ToSql()

	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}

	err = tx.Commit()

	return err
}
