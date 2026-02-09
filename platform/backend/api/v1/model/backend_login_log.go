package model

import (
	"backend/pkg/utils"
	"definition"
	"time"
)

type GetBackendLoginLogListRequest struct {
	AgentId        int    `json:"agent_id" form:"agent_id"`               // 代理id
	UserName       string `json:"username" form:"username"`               // 代理玩家名稱
	TimezoneOffset int    `json:"timezone_offset" form:"timezone_offset"` // UTC+0時間與本地時間差(分鐘) default(0)
	DateTimeTableRequest
}

func (r *GetBackendLoginLogListRequest) CheckParams(reportTimeMinuteIncrement, reportTimeRange, reportTimeBeforeDays int) int {
	if r.AgentId <= definition.AGENT_ID_ALL {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	if code := r.CheckDateTimeTableRequest(); code != definition.ERROR_CODE_SUCCESS {
		return code
	}

	if r.StartTime.Second() > 0 ||
		r.EndTime.Second() > 0 ||
		r.StartTime.Minute()%reportTimeMinuteIncrement != 0 ||
		r.EndTime.Minute()%reportTimeMinuteIncrement != 0 ||
		r.TimezoneOffset < -840 || r.TimezoneOffset > 720 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	} else if r.EndTime.Sub(r.StartTime) > time.Duration(reportTimeRange)*24*time.Hour {
		return definition.ERROR_CODE_ERROR_TIME_RANGE
	} else {
		todayUTC := utils.GetTimeNowUTCTodayTime().Add(time.Duration(r.TimezoneOffset * int(time.Minute)))
		beforeUTCDays := todayUTC.AddDate(0, 0, -reportTimeBeforeDays)
		if r.StartTime.Before(beforeUTCDays) {
			return definition.ERROR_CODE_ERROR_TIME_BEFORE_DAYS
		}
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetBackendLoginLogListResponse struct {
	AgentId   int       `json:"agent_id"`   // 代理id
	UserName  string    `json:"username"`   // 後台帳號
	ErrorCode int       `json:"error_code"` // 錯誤代碼
	Ip        string    `json:"ip"`         // 登入IP
	LoginTime time.Time `json:"login_time"` // 登入時間
}
