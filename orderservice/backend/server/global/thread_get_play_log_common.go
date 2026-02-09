package global

import (
	"backend/api/game/model"
	"backend/pkg/database"
	"backend/pkg/logger"
	"backend/pkg/utils"
	"database/sql"
	"fmt"
	"time"
)

var (
	execTimeForCheck = time.Duration(1) * time.Second

	// table name
	TN_USER_PLAY_LOG   = "user_play_log"
	TN_PLAY_LOG_COMMON = "play_log_common"

	// 批量處理筆數
	SQL_BATCH_SIZE = 2000
)

func T_execTimeForGetPlaylogCommon(logger *logger.RuntimeGoLogger, idb database.IDatabase) {
	ticker := time.NewTicker(execTimeForCheck)
	for t := range ticker.C {
		second := t.Second()
		if second == 0 || second == 30 {
			execStartTime := time.Now().UTC()
			_execTimeForGetPlaylogCommon(logger, idb)
			logger.Printf("T_execTimeForGetPlaylogCommon exec spend time: %v", time.Now().UTC().Sub(execStartTime))
		}
	}
}

// 固定時間取注單，並存入 DB
// 每次取注單的時候，取目前時間前一分鐘的資料
// 0秒時取注單，並清除上一次取的資料紀錄，並記錄本次取的注單 id
// 30秒時取注單並比對0秒時取的注單列表，是否有新增的資料
func _execTimeForGetPlaylogCommon(logger *logger.RuntimeGoLogger, idb database.IDatabase) {
	// dbBetId := ""
	dbLognumber := ""
	dbPlaylog := ""
	dbGameId := 0
	dbRoomType := 0
	dbDeskId := 0
	var dbBetTime, dbStartTime, dbEndTime time.Time

	unSyncFlag := false

	query := `select lognumber, playlog, game_id, room_type, desk_id, bet_time, start_time, end_time
	FROM play_log_common 
	where sync=$1 
	ORDER BY create_time ASC 
	LIMIT $2;`
	rows, err := idb.GetDB(DB_IDX_GAME).Query(query, unSyncFlag, SQL_BATCH_SIZE)
	//.Scan(&dbLognumber, &dbPlaylog, &dbGameId, &dbRoomType, &dbDeskId, dbBetTime)
	if err != nil {
		logger.Printf("select from DB_IDX_GAME has error is %v", err)
	}

	defer rows.Close()

	tmps := make([]model.UserPlaylog, 0)

	for rows.Next() {
		if err := rows.Scan(&dbLognumber, &dbPlaylog, &dbGameId, &dbRoomType, &dbDeskId,
			&dbBetTime, &dbStartTime, &dbEndTime); err != nil {
			rows.Close()
			return
		}

		playlogMap := utils.ToMap([]byte(dbPlaylog))
		playerDataMap, ok := playlogMap["playerlog"].([]interface{})
		if ok {
			for _, val := range playerDataMap {
				valMap, ok := val.(map[string]interface{})
				if ok {
					var tmp model.UserPlaylog
					// 遊戲局號
					tmp.OrderId = dbLognumber
					// 玩家帳號
					tmp.Username = valMap["username"].(string)
					// 房間類型
					tmp.RoomType = dbRoomType
					// 遊戲id
					tmp.GameId = dbGameId
					// 桌號
					tmp.DeskId = dbDeskId
					// 座位
					seatId := -1
					seatIdI, ok := valMap["seatId"]
					if ok {
						seatId = int(seatIdI.(float64))
					}
					tmp.SeatId = seatId
					// 有效投注
					vaildYaScore, ok := valMap["valid_score"].(float64)
					if !ok {
						vaildYaScore = .0
					}
					tmp.VaildScore = vaildYaScore
					// 全部投注
					yaScore, ok := valMap["ya_score"].(float64)
					if !ok {
						yaScore = .0
					}
					tmp.YaScore = yaScore
					// 營利 (得分-壓分)
					deScore, ok := valMap["de_score"].(float64)
					if !ok {
						deScore = .0
					}
					tmp.DeScore = deScore
					// 抽水
					tax, ok := valMap["tax"].(float64)
					if !ok {
						tax = .0
					}
					tmp.Tax = tax
					// 紅利
					bonus, ok := valMap["bonus"].(float64)
					if !ok {
						bonus = .0
					}
					tmp.Bonus = bonus
					// 開獎時間
					tmp.BetTime = dbBetTime
					// 遊戲開始時間
					tmp.StartTime = dbStartTime
					// 遊戲結束時間
					tmp.EndTime = dbEndTime
					// 代理號
					agentId := 0
					userId := utils.ToInt(valMap["user_id"])
					valAgent := AgentDataOfGameUserCache.Get(userId)
					if valAgent == nil {
						agentId = -1 // 表示需要修復
					} else {
						agentId = valAgent.AgentId
					}
					tmp.AgentId = agentId
					// 注單號
					// lognumber, agentId, gameId, roomType, deskId, seatId, username, time.Now().UnixMilli()
					salt := fmt.Sprintf("%s_%d_%d_%d_%d_%d_%s_%d", dbLognumber, agentId, tmp.GameId, dbRoomType, tmp.DeskId, tmp.SeatId, tmp.Username, dbBetTime.UnixMilli())
					// 4 is definition.ORDER_TYPE_USER_PLAY_LOG
					tmp.BetId = utils.CreatreOrderIdByOrderTypeAndSalt(4, salt, dbBetTime)
					tmps = append(tmps, tmp)
				}
			}
		}
	}

	if len(tmps) > 0 {
		// TODO: insert order data
		// successOrderList, failOrderList,
		successOrderList, err := createUserPlaylog(logger, idb.GetDB(DB_IDX_ORDER), TN_USER_PLAY_LOG, tmps)
		if err != nil {
			logger.Printf("select from DB_IDX_ORDER, create data has error : %v", err)
		}

		if len(successOrderList) > 0 {
			updatePlaylogCommonSync(logger, idb.GetDB(DB_IDX_GAME), TN_PLAY_LOG_COMMON, successOrderList)
		}
	}
}

func createUserPlaylog(logger *logger.RuntimeGoLogger, db *sql.DB, tablename string, datas []model.UserPlaylog) (map[string]bool, error) {
	retSuccOrderIdList := make(map[string]bool, 0)

	sqlStr := fmt.Sprintf(`INSERT INTO
		"public"."%s" ("bet_id", "agent_id", "lognumber", "username", "game_id",
		"room_type", "desk_id", "seat_id", "de_score", "ya_score",
		"valid_score", "tax", "bonus", "bet_time", "create_time",
		"start_time", "end_time") VALUES `, tablename)

	sqlVauleStr := `($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),` //17

	now := time.Now().UTC()
	vals := []interface{}{}
	idx := 1

	for _, data := range datas {
		sqlStr += fmt.Sprintf(sqlVauleStr,
			idx, idx+1, idx+2, idx+3, idx+4,
			idx+5, idx+6, idx+7, idx+8, idx+9,
			idx+10, idx+11, idx+12, idx+13, idx+14,
			idx+15, idx+16)
		idx += 17

		vals = append(vals,
			data.BetId, data.AgentId, data.OrderId, data.Username, data.GameId,
			data.RoomType, data.DeskId, data.SeatId, data.DeScore, data.YaScore,
			data.VaildScore, data.Tax, data.Bonus, data.BetTime, now,
			data.StartTime, data.EndTime)

		retSuccOrderIdList[data.OrderId] = true
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	//prepare the statement
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return retSuccOrderIdList, err
	}
	defer stmt.Close()

	if stmt != nil {
		//format all vals at once
		if _, err := stmt.Exec(vals...); err != nil {
			return retSuccOrderIdList, err
		}
	}

	return retSuccOrderIdList, nil
}

func updatePlaylogCommonSync(logger *logger.RuntimeGoLogger, db *sql.DB, tablename string, datas map[string]bool) error {
	inParamPart := ""

	for orderId := range datas {
		inParamPart += "'" + orderId + "',"
	}
	if inParamPart != "" {
		inParamPart = inParamPart[0 : len(inParamPart)-1]
	}

	query := fmt.Sprintf(`UPDATE %s SET sync=true WHERE lognumber IN (%s);`, tablename, inParamPart)

	result, err := db.Exec(query)
	if err != nil {
		return err
	}

	count, _ := result.RowsAffected()
	logger.Printf("updatePlaylogCommonSync: query exec count = %v", count)

	return nil
}
