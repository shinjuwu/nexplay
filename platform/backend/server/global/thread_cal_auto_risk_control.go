package global

import (
	"backend/internal/notification"
	"backend/pkg/logger"
	"backend/pkg/utils"
	"database/sql"
	"definition"
	"fmt"
	"time"
)

var (
	execAutoRiskControlTimeForCheck = 30 * time.Second
)

func T_CalAutoRiskControl(logger *logger.RuntimeGoLogger, db *sql.DB) {
	c := time.NewTicker(execAutoRiskControlTimeForCheck)
	for {
		_execGameUserRiskControlGameUsers(logger, db)
		<-c.C
	}
}

func _execGameUserRiskControlGameUsers(logger *logger.RuntimeGoLogger, db *sql.DB) {
	// 固定檢測時間
	now := time.Now().UnixNano()
	ftTimeNow := now / int64(time.Millisecond)

	logger.Printf("_execGameUserRiskControlGameUsers run, ftTimeNow: %v", ftTimeNow)

	autoRiskControlSetting := AutoRiskControlSettingCache.Get()
	if autoRiskControlSetting == nil || !autoRiskControlSetting.IsEnabled {
		return
	}

	needSendSoftBlockPlayer := false
	userIds := make([]int, 0)
	imNotifys := make([]string, 0)
	isSoftBlocks := make([]int, 0)

	f := func(key, value interface{}) bool {
		userId := key.(int)
		winRateMap := value.(map[string]interface{})
		username := winRateMap["username"].(string)
		agentId := utils.ToInt(winRateMap["agent_id"])
		levelCode := winRateMap["level_code"].(string)

		adjustWinRate := winRateMap["adjust_win_rate"].(float64)
		if adjustWinRate > autoRiskControlSetting.GameUserWinRateLimit {
			needSendSoftBlockPlayer = true
			userIds = append(userIds, userId)
			imNotifys = append(imNotifys, fmt.Sprintf("代理id: %v, 用戶: %v, 勝率過高，目前勝率: %v, 限制勝率: %v",
				agentId,
				username,
				adjustWinRate,
				autoRiskControlSetting.GameUserWinRateLimit))
			isSoftBlocks = append(isSoftBlocks, utils.ToInt(definition.RISK_CONTROL_STATUS_ENABLED))
		}

		needUpdate := winRateMap["need_update"].(float64)
		if needUpdate > 0 {
			// errCode != 0，表示db或網路有問題就先中斷處理下次再繼續
			errCode := UpdateGameUserRiskControlStatusByIndex(db, userId, definition.RISK_CONTROL_BET_IDX, definition.RISK_CONTROL_STATUS_ENABLED)
			if errCode != 0 {
				logger.Printf("_execGameUserRiskControlGameUsers UpdateGameUserRiskControlStatusByIndex failed, userId= %v", userId)
				return false
			}

			errCode = InsertAutoRiskControlLog(db, agentId, userId, definition.AUTO_RISK_CONTROL_RISK_CODE_GAME_USER_WIN_RATE, username, levelCode)
			if errCode != 0 {
				logger.Printf("_execGameUserRiskControlGameUsers InsertAutoRiskControlLog failed, userId= %v", userId)
				return false
			}

			winRateMap["need_update"] = float64(0)
		}

		GameUsersAutoRiskWinRateList.Store(userId, winRateMap)
		return true
	}
	GameUsersAutoRiskWinRateList.Range(f)

	// sent notify to gameserver
	if needSendSoftBlockPlayer {
		apiResult, errCode, err := notification.SendSoftBlockPlayer(userIds, isSoftBlocks)
		if err != nil {
			logger.Printf("_execGameUserRiskControlGameUsers SendSoftBlockPlayer failed, userIds: %v, apiResult= %v, errCode= %v, err= %v", userIds, apiResult, errCode, err)
		} else {
			// add im notify start.
			IMAlertTelegram(imNotifys)
			// add im notify end.

			// 刪除已處理完成的玩家
			f = func(key, value interface{}) bool {
				userId := key.(int)
				winRateMap := value.(map[string]interface{})

				if winRateMap["need_update"].(float64) == 0 {
					GameUsersAutoRiskWinRateList.Delete(userId)
				}

				return true
			}
			GameUsersAutoRiskWinRateList.Range(f)
		}
	}
}
