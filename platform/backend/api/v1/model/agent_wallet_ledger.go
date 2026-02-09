package model

import (
	"backend/pkg/utils"
	"definition"
	"time"
)

// 代理分數紀錄
type GetAgentWalletLedgerListRequest struct {
	DateTimeTableRequest
	Id             string `json:"id" form:"id"`                           // 訂單號碼
	AgentId        int16  `json:"agent_id" form:"agent_id"`               // 代理編號 default(0)
	TimezoneOffset int    `json:"timezone_offset" form:"timezone_offset"` // UTC+0時間與本地時間差(分鐘) default(0)
}

func (r *GetAgentWalletLedgerListRequest) CheckParams(reportTimeMinuteIncrement, reportTimeRange, reportTimeBeforeDays int) int {
	if r.AgentId < definition.AGENT_ID_ALL {
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

type GetAgentWalletLedgerResponse struct {
	Id         string    `json:"id"`          // 訂單號
	AgentName  string    `json:"agent_name"`  // 代理識別編碼
	UpdateTime time.Time `json:"update_time"` // 帳變時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
	Kind       int       `json:"kind"`        // 帳變類型
	ChangeSet  string    `json:"changeset"`   // 帳變內容
	Creator    string    `json:"creator"`     // 操作人
	Info       string    `json:"info"`        // 備註
}
