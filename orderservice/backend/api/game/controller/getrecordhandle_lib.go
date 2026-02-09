package controller

import (
	"backend/api/game/model"
	"backend/internal/ginweb"
	"backend/pkg/redis"
	"backend/pkg/utils"
	"database/sql"
	"time"
)

const (
	TIMEDRUATION_INTERVALS_MIN = 15 * time.Minute
)

func getGameData(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, agentId int, currency string, startTime, endTime time.Time) ([]model.GameData, error) {
	tmps := make([]model.GameData, 0)

	query := `SELECT bet_id, lognumber, username, game_id, room_type, 
	desk_id, seat_id, de_score, ya_score, valid_score,
	tax, bonus, bet_time, start_time, end_time
	FROM user_play_log 
	WHERE bet_time >= $1 AND bet_time < $2 AND agent_id = $3`

	rows, err := db.Query(query, startTime, endTime, agentId)
	if err != nil {
		return tmps, err
	}

	defer rows.Close()

	for rows.Next() {

		var tmp model.GameData

		var dbDeScore, dbYaScore float64

		if err := rows.Scan(&tmp.BetId, &tmp.OrderId, &tmp.Account, &tmp.GameId, &tmp.RoomId,
			&tmp.DeskId, &tmp.SeatId, &dbDeScore, &dbYaScore, &tmp.Validbet,
			&tmp.Revenue, &tmp.Bonus, &tmp.GameBetTime, &tmp.GameStartTime, &tmp.GameEndTime); err != nil {
			return tmps, err
		}

		tmp.Bet = dbYaScore
		tmp.Win = utils.DecimalSub(dbDeScore, utils.DecimalAdd(dbYaScore, tmp.Revenue))
		tmp.Currency = currency
		tmps = append(tmps, tmp)
	}

	return tmps, nil
}
