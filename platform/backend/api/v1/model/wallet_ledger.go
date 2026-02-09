package model

import (
	"backend/pkg/utils"
	"definition"
	"time"
)

// 上下分紀錄
type GetWalletLedgerListRequest struct {
	DateTimeTableRequest
	Id             string `json:"id" form:"id"`                             // 訂單號
	SingleWalletId string `json:"single_wallet_id" form:"single_wallet_id"` // 單一錢包識別碼
	AgentId        int16  `json:"agent_id" form:"agent_id"`                 // 代理編號 default(0)
	UserName       string `json:"username" form:"username"`                 // 代理玩家名稱
	Kind           int    `json:"kind" form:"kind"`                         // 帳變類型 default(0)
	TimezoneOffset int    `json:"timezone_offset" form:"timezone_offset"`   // UTC+0時間與本地時間差(分鐘) default(0)
}

func (r *GetWalletLedgerListRequest) CheckParams(reportTimeMinuteIncrement, reportTimeRange, reportTimeBeforeDays int) int {
	if r.AgentId < definition.AGENT_ID_ALL || r.Kind < definition.WALLET_LEDGER_KIND_ALL {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	if r.Id == "" {
		if code := r.CheckDateTimeTableRequest(); code != definition.ERROR_CODE_SUCCESS {
			return code
		}

		if r.StartTime.Second() > 0 ||
			r.EndTime.Second() > 0 ||
			r.StartTime.Minute()%reportTimeMinuteIncrement != 0 ||
			r.EndTime.Minute()%reportTimeMinuteIncrement != 0 ||
			r.TimezoneOffset < -840 || r.TimezoneOffset > 720 {
			return definition.ERROR_CODE_ERROR_REQUEST_DATA
		} else if r.EndTime.Sub(r.StartTime) > time.Duration(reportTimeRange)*time.Hour {
			return definition.ERROR_CODE_ERROR_TIME_RANGE
		} else {
			todayUTC := utils.GetTimeNowUTCTodayTime().Add(time.Duration(r.TimezoneOffset * int(time.Minute)))
			beforeUTCDays := todayUTC.AddDate(0, 0, -reportTimeBeforeDays)
			if r.StartTime.Before(beforeUTCDays) {
				return definition.ERROR_CODE_ERROR_TIME_BEFORE_DAYS
			}
		}
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetWalletLedgerResponse struct {
	Id             string    `json:"id"`               // 訂單號
	AgentId        string    `json:"agent_id"`         // 代理id
	AgentName      string    `json:"agent_name"`       // 代理名稱
	UserName       string    `json:"username"`         // 代理用戶帳號
	CreateTime     time.Time `json:"create_time"`      // 訂單時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
	Kind           int       `json:"kind"`             // 帳變類型
	CoinAmount     float64   `json:"coin_amount"`      // 執行分數
	ChangeSet      string    `json:"changeset"`        // 帳變內容
	Creator        string    `json:"creator"`          // 操作人
	Status         int       `json:"status"`           // 訂單狀態
	ErrorCode      int       `json:"error_code"`       // 錯誤訊息代碼
	Info           string    `json:"info"`             // 備註
	SingleWalletId string    `json:"single_wallet_id"` // 單一錢包上下分群組識別碼(對應一次上/下分)
}

type ConfirmWalletLedgerRequest struct {
	Id string `json:"id"` // 訂單號
}

func (r *ConfirmWalletLedgerRequest) CheckParams() int {
	if r.Id == "" {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	return definition.ERROR_CODE_SUCCESS
}
