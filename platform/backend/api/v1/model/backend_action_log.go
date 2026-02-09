package model

import (
	"backend/pkg/utils"
	"definition"
	"time"
)

type GetBackendActionLogListRequest struct {
	AgentId        int   `json:"agent_id"`                               // 代理id
	ActionType     int   `json:"action_type"`                            // 操作類型
	FeatureCodes   []int `json:"feature_codes"`                          // api feature code陣列
	TimezoneOffset int   `json:"timezone_offset" form:"timezone_offset"` // UTC+0時間與本地時間差(分鐘) default(0)
	DateTimeTableRequest
}

func (r *GetBackendActionLogListRequest) CheckParams(reportTimeMinuteIncrement, reportTimeRange, reportTimeBeforeDays int) int {
	if r.AgentId <= definition.AGENT_ID_ALL ||
		r.ActionType < definition.ACTION_LOG_TYPE_ALL {
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

type GetBackendActionLogResponse struct {
	Id          string    `json:"id"`           // 操作紀錄id
	CreateTime  time.Time `json:"create_time"`  // 操作時間
	Username    string    `json:"username"`     // 後台帳號
	ActionType  int       `json:"action_type"`  // 操作類型
	ActionLog   string    `json:"action_log"`   // 操作紀錄(json格式)
	FeatureCode int       `json:"feature_code"` // 操作api代碼
}
