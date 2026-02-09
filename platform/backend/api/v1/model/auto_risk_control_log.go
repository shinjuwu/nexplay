package model

import (
	"backend/pkg/utils"
	"definition"
	"time"
)

type GetAutoRiskControlLogListRequest struct {
	DateTimeTableRequest
	AgentId        int    `json:"agent_id" form:"agent_id"`               // 代理編號 default(0)
	UserName       string `json:"username" form:"username"`               // 代理玩家名稱
	TimezoneOffset int    `json:"timezone_offset" form:"timezone_offset"` // UTC+0時間與本地時間差(分鐘) default(0)
}

func (r *GetAutoRiskControlLogListRequest) CheckParams(reportTimeMinuteIncrement, reportTimeRange, reportTimeBeforeDays int) int {
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

	if r.AgentId < definition.AGENT_ID_ALL {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetAutoRiskControlLogResponse struct {
	AgentId    int       `json:"agent_id"`    // 代理編號
	AgentName  string    `json:"agent_name"`  // 代理名稱
	CreateTime time.Time `json:"create_time"` // 紀錄時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
	UserId     int       `json:"user_id"`     // 代理用戶編號
	UserName   string    `json:"username"`    // 代理用戶帳號
	RiskCode   int       `json:"risk_code"`   // 風險代碼(1:玩家每秒請求api次數過多 2:玩家上下分的比例過高 3:玩家RTP過高 4:玩家勝率過高)
}
