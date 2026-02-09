package sqlutils

import (
	"backend/internal/ginweb"
	"backend/internal/notification"
	"backend/pkg/utils"
	"backend/server/global"
	table_model "backend/server/table/model"
	"database/sql"
	"definition"
	"fmt"
	"time"

	api_game_model "backend/api/game/model"
	"backend/api/intercom/model"
)

// single wallet 上分開始
// 1. 代理錢包扣款及統計
// 2. 創建上下分紀錄
// 3. 回傳 json 格式，code:結果(0:成功，其他:錯誤)
func SWCoinInOutStart(logger ginweb.ILogger, db *sql.DB, agentId, agentCooperation, userId, kind int,
	coin, toCoin float64, currency, orderId, username, agentLevelCode, info, creator, request, agentname, walletLedgerId string) (code int, err error) {
	status := definition.WALLET_LEDGER_STATUS_CREATED

	addAgentWalletAmount := float64(0)
	addAgentWalletSumCoinOut := coin

	if kind == definition.WALLET_LEDGER_KIND_SINGLE_WALLET_UP {
		// 代理合作模式為買分需要扣款
		if agentCooperation == definition.AGENT_COOPERATION_BUY_POINT {
			status = definition.WALLET_LEDGER_STATUS_AGENT_DEDUCTED
			addAgentWalletAmount = coin
		}

		// 先建立訂單
		code, err = global.UdfGameUserStartCoinIn(db, orderId, userId, username, agentId,
			agentLevelCode, kind, status, info, creator,
			request, addAgentWalletAmount, addAgentWalletSumCoinOut, walletLedgerId)
		code = global.TranfromUdfGameUserStartCoinInCodeToModelErrorCode(code)
	} else if kind == definition.WALLET_LEDGER_KIND_SINGLE_WALLET_DOWN {
		// 下分不需要扣款
		// 先建立訂單
		code, err = global.UdfGameUserStartCoinOut(db, orderId, userId, username, agentId,
			agentLevelCode, kind, status, info, creator,
			request, walletLedgerId)
		code = global.TranfromUdfGameUserStartCoinOutCodeToModelErrorCode(code)
	} else {
		code = api_game_model.Response_GameServeDepositTypeFailed_Error
		err = fmt.Errorf("unknow kind of wallet ledger")
	}
	return
}

// addAgentWalletAmount: 代理額度扣款
// addAgentWalletSumCoinOut: 統計
// addUserSumCoinIn
func SWCoinInOutFinish(logger ginweb.ILogger, db *sql.DB, agentId, agentCooperation, userId, kind, apiErrCode int,
	coin, toCoin float64, currency, orderId, username, agentname string,
	resp map[string]interface{}) (code int, err error) {
	status := definition.WALLET_LEDGER_STATUS_SUCCESS

	addAgentWalletAmount := float64(0)
	addAgentWalletSumCoin := float64(0)
	addUserSumCoin := float64(0)
	changeset := make(map[string]interface{}, 0)
	// 更新訂單
	if kind == definition.WALLET_LEDGER_KIND_SINGLE_WALLET_UP {
		if apiErrCode != definition.ERROR_CODE_SUCCESS {
			// 如果任何原因扣款失敗，代理餘額不扣款
			addAgentWalletAmount = coin
			addAgentWalletSumCoin = coin
		} else {
			addUserSumCoin = coin
		}
		changeset = global.CreateWalletChangeset(0, coin, 0, toCoin, currency)
		changeset["sw_resp"] = resp
		code, err = global.UdfGameUserFinishCoinIn(
			db, orderId, utils.ToJSON(changeset), status, apiErrCode,
			agentId, addAgentWalletAmount, addAgentWalletSumCoin, userId, addUserSumCoin)
		code = global.TranfromUdfGameUserFinishCoinInCodeToModelErrorCode(code)
	} else if kind == model.SingleWalletRequest_Command_CancelAddScore_resend {
		if apiErrCode == definition.ERROR_CODE_SUCCESS {
			// 只要取消上分，代理餘額扣款就要加回
			addAgentWalletAmount = coin
			addAgentWalletSumCoin = coin
			status = definition.WALLET_LEDGER_STATUS_CANCEL
		} else {
			status = definition.WALLET_LEDGER_STATUS_CANCEL_FAILED
		}

		changeset = global.CreateWalletChangeset(0, coin, 0, toCoin, currency)
		changeset["sw_resp"] = resp

		code, err = global.UdfGameUserFinishCoinInCancel(
			db, orderId, utils.ToJSON(changeset), status, apiErrCode,
			agentId, addAgentWalletAmount, addAgentWalletSumCoin, userId, addUserSumCoin)
		code = global.TranfromUdfGameUserFinishCoinInCodeToModelErrorCode(code)
	} else if kind == definition.WALLET_LEDGER_KIND_SINGLE_WALLET_DOWN {
		addCoin := -coin
		if apiErrCode == definition.ERROR_CODE_SUCCESS {
			addUserSumCoin = coin
			addAgentWalletSumCoin = coin
			if agentCooperation == definition.AGENT_COOPERATION_BUY_POINT {
				addAgentWalletAmount = coin
			}
		}
		changeset = global.CreateWalletChangeset(0, addCoin, 0, toCoin, currency)

		changeset["sw_resp"] = resp
		code, err = global.UdfGameUserFinishCoinOut(
			db, orderId, utils.ToJSON(changeset), status, apiErrCode,
			agentId, addAgentWalletAmount, addAgentWalletSumCoin, userId, addUserSumCoin)
		code = global.TranfromUdfGameUserFinishCoinOutCodeToModelErrorCode(code)
	} else {
		code = api_game_model.Response_GameServeDepositTypeFailed_Error
		err = fmt.Errorf("unknow kind of wallet ledger")
	}
	if addUserSumCoin >= 5000 {
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

	return
}
