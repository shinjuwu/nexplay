package module

import (
	game_model "backend/api/game/model"
	"backend/api/intercom/model"
	"backend/api/intercom/sqlutils"

	"backend/pkg/logger"
	"backend/pkg/utils"
	"backend/server/global"
	"database/sql"
	"definition"
	"encoding/json"
	"net/http"
	"time"
)

/*
api 重送機制
*/

var (
	// 單位時間檢查是否需要啟動重送
	resendSingleWalletRequestTimeInterval = time.Duration(1) * time.Minute
	// 重送時間點(分鐘)
	resendSingleWalletRequestTimeCheckPerMin = 5
)

func T_ResendSingleWalletRequest(logger *logger.RuntimeGoLogger, db *sql.DB) {
	ticker := time.NewTicker(resendSingleWalletRequestTimeInterval)
	for t := range ticker.C {
		min := t.Minute()
		if min%resendSingleWalletRequestTimeCheckPerMin == 0 {
			execStartTime := time.Now().UTC()
			_execResendSingleWalletRequest(logger, db)
			logger.Printf("T_ResendSingleWalletRequest exec spend time: %v", time.Now().UTC().Sub(execStartTime))
		}
	}
}

func _execResendSingleWalletRequest(logger *logger.RuntimeGoLogger, db *sql.DB) {
	// 檢測時間
	now := time.Now().UnixNano()
	ftTimeNow := now / int64(time.Millisecond)
	logger.Printf("_execResendSingleWalletRequest run, ftTimeNow: %v", ftTimeNow)

	tmps := global.SelectReSendData(logger, db, global.RESEND_TYPE_SINGLEWALLET, 10)

	logger.Printf("SelectReSendData() run, data: %v", tmps)

	if len(tmps) > 0 {
		/** 加入 API 重送**/

		//重送結果ID紀錄
		finishedOrderIds, unFinishedOrderIds := reSendSingleWalletData(logger, db, tmps)

		// 送出成功的API 要更改資料表內的旗標
		global.UpdateReSendData(logger, db, true, global.RESEND_TYPE_SINGLEWALLET, finishedOrderIds)
		global.UpdateReSendData(logger, db, false, global.RESEND_TYPE_SINGLEWALLET, unFinishedOrderIds)
	}
}

// 加入 API 重送
func reSendSingleWalletData(logger *logger.RuntimeGoLogger, db *sql.DB, reSendData []*global.ReSendData) ([]string, []string) {

	finishedOrderIds := make([]string, 0)
	unFinishedOrderIds := make([]string, 0)

	if len(reSendData) <= 0 {
		return finishedOrderIds, unFinishedOrderIds
	}

	for _, dat := range reSendData {
		body, httpCode, err := utils.GetBasicAuthAPIWithHttpCode(dat.RequestToURL.String(), dat.ApiKey, "")
		// 實際發生error or 平台方回應 httpCode = 500 強制啟動重送
		if err != nil || httpCode == http.StatusInternalServerError {
			unFinishedOrderIds = append(unFinishedOrderIds, dat.Id)
		} else {
			type ResponseData struct {
				Code int         `json:"code"`
				Msg  string      `json:"msg"`
				Data interface{} `json:"data"`
			}
			var resBodyData ResponseData
			err = json.Unmarshal([]byte(body), &resBodyData)
			apiErrCode := game_model.Response_Success
			// 平台方回應 api code = INTERCOME_ERROR_CODE_SINGLE_WALLET_FORCED_RESENDING(699) 強制啟動重送
			if resBodyData.Code == definition.INTERCOME_ERROR_CODE_SINGLE_WALLET_FORCED_RESENDING {
				unFinishedOrderIds = append(unFinishedOrderIds, dat.Id)
				continue
			} else if err != nil || resBodyData.Code > 0 {
				apiErrCode = game_model.Response_GameServeDepositFailed_Error
			}
			// 無論成功或是失敗都要跑結束上下分的流程
			agentObj := global.AgentCache.Get(dat.AgentId)
			exchangeData := global.ExchangeDataCache.Get(agentObj.Currency)
			toCoin := exchangeData.ToCoin
			kind := 0
			switch dat.RequestFromParam.Command {
			case model.SingleWalletRequest_Command_AddScore: // 上分重送，種類要改成取消上分
				kind = model.SingleWalletRequest_Command_CancelAddScore_resend // 取消上分 command
			case model.SingleWalletRequest_Command_MinusScore:
				kind = definition.WALLET_LEDGER_KIND_SINGLE_WALLET_DOWN
			}
			_, _ = sqlutils.SWCoinInOutFinish(logger, db, agentObj.Id, agentObj.Cooperation, dat.RequestFromParam.UserId,
				kind, apiErrCode, dat.RequestFromParam.Point, toCoin, agentObj.Currency, dat.Id, dat.RequestFromParam.Username, agentObj.Name, utils.ToMap(body))

			finishedOrderIds = append(finishedOrderIds, dat.Id)
		}
	}

	return finishedOrderIds, unFinishedOrderIds
}
