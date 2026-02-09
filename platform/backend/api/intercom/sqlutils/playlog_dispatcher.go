package sqlutils

import (
	"backend/internal/ginweb"
	"backend/internal/notification"
	"backend/pkg/redis"
	"backend/pkg/utils"
	"backend/server/global"
	table_model "backend/server/table/model"
	"database/sql"
	"definition"
	"fmt"
	"time"
)

// 依照遊戲id 分別產生不同格式的遊戲紀錄
func DispatchUserPlayLogFromGameId(db *sql.DB, redisClient redis.IRedisCliect, logger ginweb.ILogger, tableName, lognumber, playlog string,
	gameId, roomType, deskId int, exchange float64, betTime time.Time) error {

	// 只要DB裡面有遊戲資料設定就會解析 user play log
	table := global.GameCache.Get(gameId)
	if table == nil {
		return fmt.Errorf("DispatchPlayLogFromGameId() game id is not exist, game id is: %v", gameId)
	}

	// playerStatDataList 提供遊戲即時統計使用
	playerStatDataList, jackpotTokenList, err := CreateUserPlayLog(db, redisClient, tableName, lognumber, playlog,
		gameId, roomType, deskId, exchange, betTime)

	if err == nil && len(jackpotTokenList) > 0 {
		err = CreateJackpotTokenLog(db, jackpotTokenList)
	}

	// 有紀錄且沒錯誤才跑統計
	// 統計資料從
	if err == nil && len(playerStatDataList) > 0 {
		CalRTGameUsersStat(db, playerStatDataList, betTime)
		CalRPAgentGameRatioStat(db, playerStatDataList, betTime)

		autoRiskControlSetting := global.AutoRiskControlSettingCache.Get()

		for userId, playData := range playerStatDataList {
			if playData.GameType > definition.GAME_TYPE_BAIREN {
				valMap, ok := global.GameUsersKillDiveValueList.Load(userId)
				if ok {
					// 有追殺額度才處理
					killDiveValMap := valMap.(map[string]float64)
					if killDiveValMap["kill_dive_value"] > 0 {
						finalScore := utils.DecimalSub(playData.YaScore, playData.DeScore)
						newVal := utils.DecimalSub(killDiveValMap["kill_dive_value"], finalScore)
						if newVal <= 0 {
							newVal = 0
						}
						killDiveValMap["need_update"] = float64(1)
						killDiveValMap["kill_dive_value"] = newVal
						global.GameUsersKillDiveValueList.Store(userId, killDiveValMap)
					}
				}
			}

			// autoRiskControlSetting如果是nil，表示redis有問題，略過不處理
			// agentId如果是-1，表示是有問題的紀錄，略過不處理
			if autoRiskControlSetting != nil && autoRiskControlSetting.IsEnabled && playData.AgentId != -1 {
				var totalGame int64
				var totalWin int64
				var err error

				isWin := playData.DeScore-playData.YaScore-playData.Tax > 0

				totalGame, err = global.AutoRiskControlStatCache.IncrGameUserTotalGame(playData.UserId)
				if err == nil && isWin {
					totalWin, err = global.AutoRiskControlStatCache.IncrGameUserTotalWin(playData.UserId)
				}

				if err == nil && isWin {
					adjustWinRate := float64(totalWin) / (float64(totalGame + 10))
					if adjustWinRate > autoRiskControlSetting.GameUserWinRateLimit {
						global.GameUsersAutoRiskWinRateList.Store(userId, map[string]interface{}{
							"adjust_win_rate": adjustWinRate,
							"need_update":     float64(1),
							"agent_id":        playData.AgentId,
							"username":        playData.Username,
							"level_code":      playData.LevelCode,
						})
					}
				}
			}
		}
	}

	// 異常輸贏資料收集
	if len(playerStatDataList) > 0 {
		conninfo := make(map[string]interface{}, 0)
		addr, ok := global.ServerInfoCache.Load("monitor")
		if ok {
			addres, ok := addr.(table_model.ServerInfo)
			if ok {
				conninfo = utils.ToMap(addres.AddressesBytes)
			}
		}

		awl := notification.NewTmpCollectorAbnormalWinAndLoseRequest()
		awl.ID = global.DEF_PLATFORM

		rtps := notification.NewTmpCollectorRTPStatRequest()
		rtps.ID = global.DEF_PLATFORM

		for _, playData := range playerStatDataList {
			if playData.AgentId != -1 {
				if playData.BigWinCount > 0 {
					var gamename, agentname, roomname string
					if gObject := global.GameCache.Get(gameId); gObject != nil {
						gamename = gObject.Name
					}
					if aObject := global.AgentCache.Get(playData.AgentId); aObject != nil {
						agentname = aObject.Name
					}

					grc := global.GameRoomCache.Get(gameId*10 + roomType)
					if grc != nil {
						roomname = grc.Name
					}

					tmp := new(notification.TmpUserPlayLog)
					tmp.AgentID = playData.AgentId
					tmp.AgentName = agentname
					tmp.BetID = playData.BetId
					tmp.BetTime = betTime
					tmp.Bonus = playData.Bonus
					tmp.CreateTime = time.Now()
					tmp.DeScore = playData.DeScore
					tmp.GameID = gameId
					tmp.GameName = gamename
					tmp.LogNumber = lognumber
					tmp.RoomType = roomType
					tmp.RoomName = roomname
					tmp.UserID = playData.UserId
					tmp.Username = playData.Username

					awl.Data = append(awl.Data, tmp)
				}

				tmp2 := new(notification.TmpCollectorRTPStat)
				tmp2.GameId = playData.GameId
				tmp2.Bonus = playData.Bonus
				tmp2.De = playData.DeScore
				tmp2.GameType = playData.GameType
				tmp2.PlayCount = 1
				tmp2.RoomType = playData.RoomType
				tmp2.Tax = playData.Tax
				tmp2.UpdateTime = betTime
				tmp2.VaildYa = playData.VaildYaScore
				tmp2.Ya = playData.YaScore

				rtps.Data = append(rtps.Data, tmp2)
			}
		}

		if len(awl.Data) > 0 {
			ok, _ := notification.SendCollectorAbnormalWinAndLoseToMonitorService(conninfo, awl)
			if !ok {
				if _, err := global.MonitorServiceCache.AddSendCollectorAbnormalWinAndLoseData(awl); err != nil {
					logger.Error("DispatchUserPlayLogFromGameId SendCollectorAbnormalWinAndLoseToMonitorService fail, data: %s", utils.ToJSON(awl))
				}
			}
		}

		if len(rtps.Data) > 0 {
			ok, _ := notification.SendCollectorRTPStatToMonitorService(conninfo, rtps)
			if !ok {
				if _, err := global.MonitorServiceCache.AddSendCollectorRTPStatData(rtps); err != nil {
					logger.Error("DispatchUserPlayLogFromGameId SendCollectorRTPStatToMonitorService fail, data: %s", utils.ToJSON(rtps))
				}
			}
		}
	}

	return err
}
