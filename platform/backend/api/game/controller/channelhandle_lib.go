package controller

import (
	"backend/api/game/model"
	"backend/internal/ginweb"
	"backend/internal/notification"
	"backend/pkg/redis"
	"backend/pkg/utils"
	"backend/server/global"
	"database/sql"
	"definition"
	"fmt"
	"strconv"
	"strings"
	"time"

	table_model "backend/server/table/model"

	"github.com/shopspring/decimal"
)

// 創建新帳號處理
func createNewAccount(logger ginweb.ILogger, db *sql.DB, userId, agentId, agentCooperation int,
	coin, toCoin float64, currency, orderId, original_username, username, agentLevelCode, info, creator, request, agentname string,
	createNew bool) (gold float64, code int, err error) {
	code = definition.ERROR_CODE_SUCCESS

	kind := definition.WALLET_LEDGER_KIND_API_UP
	// status := definition.WALLET_LEDGER_STATUS_CREATED

	// addAgentWalletAmount := float64(0)
	// addAgentWalletSumCoinOut := coin

	if createNew {
		var apiCode int
		apiCode, err = notification.SendCreateAccount("", 0, agentId, int64(userId), original_username, username)
		if err != nil {
			logger.Info("SendCreateAccount has error, err: %v", err)
		}

		code = notification.TransformApiCodeToModelErrorCode(apiCode)
	}

	// 創建帳號成功才加款
	if code == model.Response_Success {
		if coin > 0 {
			gold, code, err = coinIn(logger, db, agentId, agentCooperation, userId,
				kind, coin, toCoin, currency, orderId,
				username, agentLevelCode, info, creator, request,
				agentname)
		}
	}

	return
}

// 上分
func coinIn(logger ginweb.ILogger, db *sql.DB, agentId, agentCooperation, userId, kind int,
	coin, toCoin float64, currency, orderId, username, agentLevelCode, info, creator, request, agentname string) (gold float64, code int, err error) {
	status := definition.WALLET_LEDGER_STATUS_CREATED

	addAgentWalletAmount := float64(0)
	addAgentWalletSumCoinOut := coin

	// 代理合作模式為買分需要扣款
	if agentCooperation == definition.AGENT_COOPERATION_BUY_POINT {
		status = definition.WALLET_LEDGER_STATUS_AGENT_DEDUCTED
		addAgentWalletAmount = coin
	}

	// 先建立訂單
	code, err = global.UdfGameUserStartCoinIn(db,
		orderId, userId, username, agentId, agentLevelCode,
		kind, status, info, creator, request,
		addAgentWalletAmount, addAgentWalletSumCoinOut, "")
	if code != definition.ERROR_CODE_SUCCESS {
		return
	}

	// 通知 game server 處理上分
	var afterCoin float64
	var apiCode int
	afterCoin, apiCode, err = notification.SendDeposit("up", coin, int64(userId))
	if err != nil {
		logger.Info("SendDeposit failed. orderId: %s, code: %d, err: %v", orderId, code, err)
	}

	gold = afterCoin
	apiCode = notification.TransformApiCodeToModelErrorCode(apiCode)

	addUserSumCoinIn := float64(0)

	var beforeCoin, addCoin float64
	if apiCode == definition.ERROR_CODE_SUCCESS {
		addCoin = coin
		beforeCoin = utils.DecimalSub(afterCoin, addCoin)

		// 成功則原先的扣款金額及統計不需要rollback
		addAgentWalletAmount = 0
		addAgentWalletSumCoinOut = 0
		addUserSumCoinIn = coin
	}

	changeset := global.CreateWalletChangeset(beforeCoin, addCoin, afterCoin, toCoin, currency)
	status = definition.WALLET_LEDGER_STATUS_SUCCESS

	// 更新訂單
	code, err = global.UdfGameUserFinishCoinIn(db, orderId, utils.ToJSON(changeset), status, apiCode,
		agentId, addAgentWalletAmount, addAgentWalletSumCoinOut, userId, addUserSumCoinIn)

	if coin >= 5000 {
		// 加入 上下分監控
		conninfo := make(map[string]interface{}, 0)
		addr, ok := global.ServerInfoCache.Load("monitor")
		if ok {
			addres, ok := addr.(table_model.ServerInfo)
			if ok {
				conninfo = utils.ToMap(addres.AddressesBytes)
			}

			d := notification.NewTempCollectorCoinInOutRequest()
			d.ID = global.DEF_PLATFORM

			tmp := new(notification.TempWalletLedger)
			tmp.ID = orderId
			tmp.AgentID = agentId
			tmp.AgentName = agentname
			tmp.ChangeSet = utils.ToJSON(changeset)
			tmp.CreateTime = time.Now()
			tmp.Kind = kind
			tmp.Status = status
			tmp.UserID = userId
			tmp.Username = username

			d.Data = append(d.Data, tmp)

			ok, _ = notification.SendCollectorCoinInOutToMonitorService(conninfo, d)
			if !ok {
				if _, err := global.MonitorServiceCache.AddSendCollectorCoinInOutData(d); err != nil {
					logger.Error("Channelhandle CoinIn SendCollectorCoinInOutToMonitorService fail, data: %s", utils.ToJSON(d))
				}
			}
		}
	}

	// apiCode如果有問題要優先回傳告知
	if apiCode != definition.ERROR_CODE_SUCCESS {
		code = apiCode
	}

	return
}

// 下分
func coinOut(logger ginweb.ILogger, db *sql.DB, agentId, agentCooperation, userId, kind int,
	coin, toCoin float64, currency, orderId, username, agentLevelCode, info, creator, request, agentname string) (gold float64, code int, err error) {
	status := definition.WALLET_LEDGER_STATUS_CREATED

	// 先建立訂單
	code, err = global.UdfGameUserStartCoinOut(db,
		orderId, userId, username, agentId, agentLevelCode,
		kind, status, info, creator, request,
		"")
	if code != definition.ERROR_CODE_SUCCESS {
		return
	}

	// 通知 game server處理下分
	var afterCoin float64
	var apiCode int
	afterCoin, apiCode, err = notification.SendDeposit("down", coin, int64(userId))
	if err != nil {
		logger.Info("SendDeposit failed. orderId: %s, code: %d, err: %v", orderId, code, err)
	}

	gold = afterCoin
	apiCode = notification.TransformApiCodeToModelErrorCode(apiCode)

	addAgentWalletAmount := float64(0)
	addAgentWalletSumCoinIn := float64(0)
	addUserSumCoinOut := float64(0)

	var beforeCoin, addCoin float64
	if apiCode == definition.ERROR_CODE_SUCCESS {
		addCoin = -coin
		beforeCoin = utils.DecimalSub(afterCoin, addCoin)

		addUserSumCoinOut = coin
		addAgentWalletSumCoinIn = coin
		if agentCooperation == definition.AGENT_COOPERATION_BUY_POINT {
			addAgentWalletAmount = coin
		}
	}

	changeset := global.CreateWalletChangeset(beforeCoin, addCoin, afterCoin, toCoin, currency)
	status = definition.WALLET_LEDGER_STATUS_SUCCESS

	// 更新訂單
	code, err = global.UdfGameUserFinishCoinOut(db, orderId, utils.ToJSON(changeset), definition.WALLET_LEDGER_STATUS_SUCCESS, apiCode,
		agentId, addAgentWalletAmount, addAgentWalletSumCoinIn, userId, addUserSumCoinOut)

	if coin >= 5000 {
		// 加入 上下分監控
		conninfo := make(map[string]interface{}, 0)
		addr, ok := global.ServerInfoCache.Load("monitor")
		if ok {
			addres, ok := addr.(table_model.ServerInfo)
			if ok {
				conninfo = utils.ToMap(addres.AddressesBytes)
			}

			d := notification.NewTempCollectorCoinInOutRequest()
			d.ID = global.DEF_PLATFORM

			tmp := new(notification.TempWalletLedger)
			tmp.ID = orderId
			tmp.AgentID = agentId
			tmp.AgentName = agentname
			tmp.ChangeSet = utils.ToJSON(changeset)
			tmp.CreateTime = time.Now()
			tmp.Kind = kind
			tmp.Status = status
			tmp.UserID = userId
			tmp.Username = username

			d.Data = append(d.Data, tmp)

			ok, _ = notification.SendCollectorCoinInOutToMonitorService(conninfo, d)
			if !ok {
				if _, err := global.MonitorServiceCache.AddSendCollectorCoinInOutData(d); err != nil {
					logger.Error("Channelhandle CoinOut SendCollectorCoinInOutToMonitorService fail, data: %s", utils.ToJSON(d))
				}
			}
		}
	}

	// apiCode如果有問題要優先回傳告知
	if apiCode != definition.ERROR_CODE_SUCCESS {
		code = apiCode
	}

	return
}

// 取得上次登入時間
func getLastLoginTinme(logger ginweb.ILogger, db *sql.DB, userId int) (lastLoginTime, lastLoginOutTime string, err error) {

	query := `SELECT login_time, logout_time FROM user_login_log WHERE game_user_id=$1 ORDER BY login_time DESC LIMIT 1;`
	err = db.QueryRow(query, userId).Scan(&lastLoginTime, &lastLoginOutTime)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

// 檢查小數點位數
func CheckFloatPlaces(dec string, place int) bool {
	rebate, err := decimal.NewFromString(dec)
	if err != nil {
		return false
	}
	if rebate.Round(int32(place)).String() != rebate.String() {
		return false
	}

	return true
}

// 偵測api request是否異常
func checkGameUserApiRequest(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, agentId, userId int, username, levelCode string, threshold int64) (isUserEnabled bool) {
	count, err := global.AutoRiskControlStatCache.IncrGameUserApiRequestCount(userId)
	if err != nil {
		logger.Info("AutoRiskControlStatCache IncrGameUserApiRequestCount has error: %v", err)
		return true
	}

	if count > threshold {
		return !disableGameUserByAutoRiskControl(logger, db, rdb, agentId, userId, username, levelCode)
	}

	return true
}

// 封停玩家
func disableGameUserByAutoRiskControl(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, agentId, userId int, username, levelCode string) bool {
	if msg, code, err := notification.SendSetBlockPlayer([]int{userId}, true); err != nil {
		logger.Info("disableGameUser() SendSetPlayerBlock has error, msg: %s, code: %d, err: %v", msg, code, err)
		return false
	}

	query := `UPDATE game_users SET is_enabled = false WHERE id=$1;`
	if _, err := db.Exec(query, userId); err != nil {
		logger.Info("disableGameUser() db exec has error, err: %v", err)
		return false
	}

	reqUserId := utils.IntToString(userId)
	serverInfoCode, err := rdb.LoadHValue(global.REDIS_IDX_LOGIN_INFO, global.REDIS_HASH_INGAME_USER, reqUserId)
	if err == nil && serverInfoCode != "" {
		_ = rdb.DeleteHValue(global.REDIS_IDX_LOGIN_INFO, global.REDIS_HASH_INGAME_USER, reqUserId)
		_ = rdb.DeleteHValue(
			global.REDIS_IDX_LOGIN_INFO,
			global.GenRedisHashName(global.REDIS_HASH_INGAME_USER, serverInfoCode),
			reqUserId)
	}

	if errorCode := global.InsertAutoRiskControlLog(db, agentId, userId, definition.AUTO_RISK_CONTROL_RISK_CODE_GAME_USER_API_REQUEST, username, levelCode); errorCode != model.Response_Success {
		logger.Info("disableGameUser() InsertAutoRiskControlLog failed, errorCode: %d", errorCode)
	}

	imNotifys := make([]string, 0)
	imNotifys = append(imNotifys,
		fmt.Sprintf("代理id: %v, 用戶: %v, 封停玩家: 頻繁呼叫 api",
			agentId,
			username))

	global.IMAlertTelegram(imNotifys)

	return true
}

// 偵測上下分是否太頻繁
func checkGameUserCoinInAndOutRequest(logger ginweb.ILogger, db *sql.DB, agentId, userId int, username string, threshold int64) (isTooManyRequest bool) {
	count, err := global.AutoRiskControlStatCache.IncrGameUserCoinInAndOutRequestCount(userId)
	if err != nil {
		logger.Info("AutoRiskControlStatCache IncrGameUserCoinInAndOutRequestCount has error: %v", err)
		return false
	}

	check := count > threshold

	if check {
		imNotifys := make([]string, 0)
		imNotifys = append(imNotifys,
			fmt.Sprintf("代理id: %v, 用戶: %v, 上下分太頻繁，目前次數: %v, 限制: %v",
				agentId,
				username,
				count,
				threshold))

		global.IMAlertTelegram(imNotifys)
	}

	return check

}

// 偵測下分是否超過限制
func checkGameUserCoinOutRequest(logger ginweb.ILogger, db *sql.DB, agentId, userId int, username, levelCode string, coinOut, threshold float64) (isUserAllowCoinOut bool) {
	totalCoinIn, totalCoinOut, err := global.AutoRiskControlStatCache.GetGameUserTotalCoinInAndTotalCoinOut(userId)
	if err != nil {
		logger.Info("AutoRiskControlStatCache GetGameUserTotalCoinInAndTotalCoinOut has error: %v", err)
		return true
	}

	newTotalCoinOut := totalCoinOut + coinOut
	if newTotalCoinOut-totalCoinIn > threshold {

		result := !disableGameUserCoinOutByAutoRiskControl(logger, db, agentId, userId, username, levelCode)

		imNotifys := make([]string, 0)
		imNotifys = append(imNotifys,
			fmt.Sprintf("代理id: %v, 用戶: %v, 玩家下分差距過大，原始上分:%v, 累積下分: %v, 本次下分: %v 限制:%v",
				agentId,
				username,
				totalCoinIn,
				totalCoinOut,
				coinOut,
				threshold))

		global.IMAlertTelegram(imNotifys)

		return result
	}

	return true
}

// 禁止玩家下分
func disableGameUserCoinOutByAutoRiskControl(logger ginweb.ILogger, db *sql.DB, agentId, userId int, username, levelCode string) bool {
	if errorCode := global.UpdateGameUserRiskControlStatusByIndex(db, userId, definition.RISK_CONTROL_COIN_OUT_IDX, definition.RISK_CONTROL_STATUS_ENABLED); errorCode != model.Response_Success {
		logger.Info("disableGameUserCoinOut() UpdateGameUserRiskControlStatusByIndex() failed, errorCode: %d", errorCode)
		return false
	}

	if errorCode := global.InsertAutoRiskControlLog(db, agentId, userId, definition.AUTO_RISK_CONTROL_RISK_CODE_GAME_USER_COIN_IN_AND_COIN_OUT_DIFF, username, levelCode); errorCode != model.Response_Success {
		logger.Info("disableGameUserCoinOut() InsertAutoRiskControlLog failed, errorCode: %d", errorCode)
	}

	return true
}

// 建立登入lobby switch info字串
func createLobbySwitchInfoString(loginGameId, agentLobbySwitchInfo int) string {
	paddingLeftZero := func(str string, length int) string {
		if len(str) >= length {
			return str
		}
		return strings.Repeat("0", length-len(str)) + str

	}
	formatInfo := func(info int) string {
		return paddingLeftZero(strconv.FormatInt(int64(info), 2), definition.LOBBY_COUNT)
	}

	lobbySwitchInfo := formatInfo(agentLobbySwitchInfo)

	loginSwitchFlag, find := definition.GameIdToLobbySwitch[loginGameId]
	if !find {
		loginSwitchFlag = 0
	}
	loginSwitchInfo := formatInfo(loginSwitchFlag)

	return fmt.Sprintf("%s|%s", loginSwitchInfo, lobbySwitchInfo)
}
