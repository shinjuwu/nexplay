package sqlutils

import (
	"backend/api/intercom/model"
	"backend/pkg/redis"
	"backend/pkg/utils"
	"backend/server/global"
	"database/sql"
	"definition"
	"fmt"
	"strconv"
	"time"
)

func CreateUserPlayLog(db *sql.DB, redisClient redis.IRedisCliect, userPlayLogTableName, lognumber, playlog string,
	gameId, roomType, deskId int, exchange float64, betTime time.Time) (map[int]*PlayerStatData, []*model.JackpotTokenLog, error) {

	playerStatDataList := make(map[int]*PlayerStatData, 0)
	jackpotTokenList := make([]*model.JackpotTokenLog, 0)

	seatId := -1
	killLevel := -1
	realPlayer := -1
	tempMap := utils.ToMap([]byte(playlog))

	sqlStr := fmt.Sprintf(`INSERT INTO
		"public"."%s" ("lognumber", "agent_id", "user_id", "game_id", "room_type",
		"desk_id", "seat_id", "exchange", "de_score", "ya_score",
		"valid_score", "start_score", "end_score", "create_time", "is_robot",
		"is_big_win", "is_issue", "bet_time", "tax", "level_code",
		"username", "bet_id", "kill_type", "bonus", "jp_inject_water_rate",
		"jp_inject_water_score", "wallet_ledger_id", "kill_prob", "kill_level", "real_players",
		"room_id") VALUES `, userPlayLogTableName)

	sqlVauleStr := `($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),` //31

	paramsCount := 31

	playerDataTemp, ok := tempMap["playerlog"].([]interface{})
	if !ok {
		return playerStatDataList, jackpotTokenList, fmt.Errorf("playerData is not array")
	}

	friendRoomInfo, isFriendRoom := tempMap["friend_room_info"].(map[string]interface{})

	vals := []interface{}{}

	now := time.Now().UTC()

	idx := 1

	for _, val := range playerDataTemp {

		valMap, ok := val.(map[string]interface{})
		if ok {
			isRobotF, ok := valMap["is_robot"].(float64)
			if !ok {
				isRobotF = .0
			}
			isRobot := int(isRobotF)

			// 不新增機器人的個人遊戲紀錄
			if isRobot == 1 {
				continue
			}

			sqlStr += fmt.Sprintf(sqlVauleStr, idx, idx+1, idx+2, idx+3, idx+4,
				idx+5, idx+6, idx+7, idx+8, idx+9,
				idx+10, idx+11, idx+12, idx+13, idx+14,
				idx+15, idx+16, idx+17, idx+18, idx+19,
				idx+20, idx+21, idx+22, idx+23, idx+24,
				idx+25, idx+26, idx+27, idx+28, idx+29,
				idx+30)
			idx += paramsCount

			userid := utils.ToInt(valMap["user_id"])
			username := utils.ToString(valMap["username"])

			// for singlewallet
			walletLedgerId := utils.ToString(valMap["wallet_ledger_id"])

			agentId := 0
			levelCode := ""
			valAgent := global.AgentDataOfGameUserCache.Get(userid)
			if valAgent == nil {
				agentId = -1 // 表示需要修復
				levelCode = ""
			} else {
				agentId = valAgent.AgentId
				levelCode = valAgent.LevelCode
			}

			bonus, ok := valMap["bonus"].(float64)
			if !ok {
				bonus = .0
			}

			tax, ok := valMap["tax"].(float64)
			if !ok {
				tax = .0
			}

			seatIdI, ok := valMap["seatId"]
			if ok {
				seatId = int(seatIdI.(float64))
			}

			deScore, ok := valMap["de_score"].(float64)
			if !ok {
				deScore = .0
			}

			yaScore, ok := valMap["ya_score"].(float64)
			if !ok {
				yaScore = .0
			}

			vaildYaScore, ok := valMap["valid_score"].(float64)
			if !ok {
				vaildYaScore = .0
			}

			isBigWin, ok := valMap["is_big_win"].(bool)
			if !ok {
				isBigWin = false
			}

			killProb, ok := valMap["kill_prob"].(float64)
			if !ok {
				killProb = .0
			}

			killLevelI, ok := valMap["kill_level"]
			if ok {
				killLevel = int(killLevelI.(float64))
			}

			realPlayerI, ok := valMap["real_players"]
			if ok {
				realPlayer = int(realPlayerI.(float64))
			}

			roomId := ""
			if isFriendRoom {
				roomId = friendRoomInfo["room_id"].(string)
			}

			bigWinCount := 0
			// 贏的金額大於壓注10倍
			if deScore > 0 && yaScore > 0 {
				if isBigWin || (deScore >= yaScore*10) {
					bigWinCount = 1
				}
			}

			isWin := deScore > yaScore
			var winCount, loseCount int
			if isWin {
				winCount = 1
			} else {
				loseCount = 1
			}

			isKillDive := utils.ToInt(valMap["kill_type"])

			jpInjectWaterRate := utils.ToFloat64(valMap["jp_commission"])
			jpInjectWaterScore := utils.ToFloat64(valMap["jp_commission _score"])

			// lognumber, agentId, gameId, roomType, deskId, seatId, username, time.Now().UnixMilli()
			salt := fmt.Sprintf("%s_%d_%d_%d_%d_%d_%s_%d", lognumber, agentId, gameId, roomType, deskId, seatId, username, betTime.UnixMilli())
			betId := utils.CreatreOrderIdByOrderTypeAndSalt(definition.ORDER_TYPE_USER_PLAY_LOG, salt, betTime)

			vals = append(vals,
				lognumber, agentId, valMap["user_id"], gameId, roomType,
				deskId, seatId, exchange, valMap["de_score"], valMap["ya_score"],
				valMap["valid_score"], valMap["start_score"], valMap["end_score"], now, isRobot,
				false, false, betTime, tax, levelCode,
				username, betId, isKillDive, bonus, jpInjectWaterRate,
				jpInjectWaterScore, walletLedgerId, killProb, killLevel, realPlayer,
				roomId)

			// 即時統計
			// 這裡只產生數據，不作實際的統計紀錄
			tmpData, ok := playerStatDataList[userid]
			if !ok {
				playerStatDataList[userid] = &PlayerStatData{
					LevelCode:    levelCode,
					Username:     username,
					AgentId:      agentId,
					GameId:       gameId,
					GameType:     gameId / 1000,
					RoomType:     roomType,
					UserId:       userid,
					DeScore:      deScore,
					YaScore:      yaScore,
					VaildYaScore: vaildYaScore,
					Tax:          tax,
					Bonus:        bonus,
					OrderCount:   1,
					WinCount:     winCount,
					LoseCount:    loseCount,
					BigWinCount:  bigWinCount,
					BetId:        betId,
				}
			} else {
				tmpData.DeScore = utils.DecimalAdd(tmpData.DeScore, deScore)
				tmpData.YaScore = utils.DecimalAdd(tmpData.YaScore, yaScore)
				tmpData.VaildYaScore = utils.DecimalAdd(tmpData.VaildYaScore, vaildYaScore)
				tmpData.Tax = utils.DecimalAdd(tmpData.Tax, tax)
				tmpData.Bonus = utils.DecimalAdd(tmpData.Bonus, bonus)
				tmpData.OrderCount += 1
				tmpData.WinCount += winCount
				tmpData.LoseCount += loseCount
				tmpData.BigWinCount += bigWinCount
			}

			// token 紀錄
			// 這裡只做產生數據，不做實際的token紀錄
			/*
				type JackpotToken struct {
					Id         string  `json:"Id"`
					CreateTime int64   `json:"CreateTime"`
					Bet        float64 `json:"bet"`
				}
			*/
			token, ok := valMap["jp_token"].(map[string]interface{})
			if ok {
				tokenId := token["Id"].(string)
				tokenJpBet := token["bet"].(float64)
				tokenCreateTimeInt64 := strconv.FormatInt(int64(token["CreateTime"].(float64)), 10)
				tokenCreateTime := utils.ToTimeUnixMilliUTC(tokenCreateTimeInt64, 0)

				jackpotTokenList = append(jackpotTokenList, &model.JackpotTokenLog{
					TokenId:         tokenId,
					AgentId:         agentId,
					LevelCode:       levelCode,
					UserId:          userid,
					Username:        username,
					JpBet:           tokenJpBet,
					TokenCreateTime: tokenCreateTime,
					SourceGameId:    gameId,
					SourceLognumber: lognumber,
					SourceBetId:     betId,
				})
			}
		}
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	//prepare the statement
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return playerStatDataList, jackpotTokenList, err
	}
	defer stmt.Close()

	if stmt != nil {
		//format all vals at once
		if _, err := stmt.Exec(vals...); err != nil {
			return playerStatDataList, jackpotTokenList, err
		}
	}

	return playerStatDataList, jackpotTokenList, nil
}

func ParseUserPlayLog(playlog string, gameId, roomType, deskId int) (map[int]*PlayerStatData, error) {

	playerStatDataList := make(map[int]*PlayerStatData, 0)
	var err error
	// seatId := -1
	tempMap := utils.ToMap([]byte(playlog))

	playerDataTemp, ok := tempMap["playerlog"].([]interface{})
	if !ok {
		return playerStatDataList, err
	}

	for _, val := range playerDataTemp {

		valMap, ok := val.(map[string]interface{})
		if ok {
			isRobotF, ok := valMap["is_robot"].(float64)
			if !ok {
				isRobotF = .0
			}
			isRobot := int(isRobotF)

			// 不新增機器人的個人遊戲紀錄
			if isRobot == 1 {
				continue
			}

			userid := utils.ToInt(valMap["user_id"])
			username := valMap["username"].(string)

			agentId := 0
			levelCode := ""
			valAgent := global.AgentDataOfGameUserCache.Get(userid)
			if valAgent == nil {
				agentId = -1 // 表示需要修復
				levelCode = ""
			} else {
				agentId = valAgent.AgentId
				levelCode = valAgent.LevelCode
			}

			bonus, ok := valMap["bonus"].(float64)
			if !ok {
				bonus = .0
			}

			tax, ok := valMap["tax"].(float64)
			if !ok {
				tax = .0
			}

			// seatIdI, ok := valMap["seatId"]
			// if ok {
			// 	seatId = int(seatIdI.(float64))
			// }

			deScore, ok := valMap["de_score"].(float64)
			if !ok {
				deScore = .0
			}

			yaScore, ok := valMap["ya_score"].(float64)
			if !ok {
				yaScore = .0
			}

			vaildYaScore, ok := valMap["valid_score"].(float64)
			if !ok {
				vaildYaScore = .0
			}

			isBigWin, ok := valMap["is_big_win"].(bool)
			if !ok {
				isBigWin = false
			}

			bigWinCount := 0
			// 贏的金額大於壓注10倍
			if deScore > 0 && yaScore > 0 {
				if isBigWin || (deScore >= yaScore*10) {
					bigWinCount = 1
				}
			}

			isWin := deScore > yaScore
			var winCount, loseCount int
			if isWin {
				winCount = 1
			} else {
				loseCount = 1
			}

			// isKillDive := utils.ToInt(valMap["kill_type"])

			// 即時統計
			// 這裡只產生數據，不作實際的統計紀錄
			tmpData, ok := playerStatDataList[userid]
			if !ok {
				playerStatDataList[userid] = &PlayerStatData{
					LevelCode:    levelCode,
					Username:     username,
					AgentId:      agentId,
					GameId:       gameId,
					GameType:     gameId / 1000,
					RoomType:     roomType,
					UserId:       userid,
					DeScore:      deScore,
					YaScore:      yaScore,
					VaildYaScore: vaildYaScore,
					Tax:          tax,
					Bonus:        bonus,
					OrderCount:   1,
					WinCount:     winCount,
					LoseCount:    loseCount,
					BigWinCount:  bigWinCount,
					BetId:        "",
				}
			} else {
				tmpData.DeScore = utils.DecimalAdd(tmpData.DeScore, deScore)
				tmpData.YaScore = utils.DecimalAdd(tmpData.YaScore, yaScore)
				tmpData.VaildYaScore = utils.DecimalAdd(tmpData.VaildYaScore, vaildYaScore)
				tmpData.Tax = utils.DecimalAdd(tmpData.Tax, tax)
				tmpData.Bonus = utils.DecimalAdd(tmpData.Bonus, bonus)
				tmpData.OrderCount += 1
				tmpData.WinCount += winCount
				tmpData.LoseCount += loseCount
				tmpData.BigWinCount += bigWinCount
			}

		}
	}

	return playerStatDataList, err
}
