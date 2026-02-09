package model

import (
	"backend/pkg/utils"
	"definition"
	"time"
)

type GetGameUsersStatHourListRequest struct {
	AgentId  int    `json:"agent_id"` // 代理編號
	Username string `json:"username"` // 玩家名稱
	DateTimeTableRequest
}

func (r *GetGameUsersStatHourListRequest) CheckParams(reportTimeRange, reportTimeBeforeDays int) int {
	if r.AgentId <= definition.AGENT_ID_ALL ||
		r.Username == "" {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	if code := r.CheckDateTimeTableRequest(); code != definition.ERROR_CODE_SUCCESS {
		return code
	}

	if r.EndTime.Sub(r.StartTime) > time.Duration(reportTimeRange)*24*time.Hour {
		return definition.ERROR_CODE_ERROR_TIME_RANGE
	}

	today := utils.GetTimeNowUTCTodayTime()
	beforeDays := time.Duration(reportTimeBeforeDays) * 24 * time.Hour
	if today.Sub(r.StartTime) > beforeDays || today.Sub(r.EndTime) > beforeDays {
		return definition.ERROR_CODE_ERROR_TIME_BEFORE_DAYS
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetGameUsersStatHourResponse struct {
	Username  string  `json:"username"`   // 玩家名稱
	UserId    int     `json:"user_id"`    // 玩家帳號編號
	AgentId   int     `json:"agent_id"`   // 代理編號
	AgentName string  `json:"agent_name"` // 代理名稱
	GameId    string  `json:"game_id"`    // 遊戲編號
	PlayCount int     `json:"play_count"` // 注單數
	DeScore   float64 `json:"de_score"`   // 總玩家得分
	YaScore   float64 `json:"ya_score"`   // 總投注
	Tax       float64 `json:"tax"`        // 總遊戲抽水
	Bonus     float64 `json:"bonus"`      // 紅利
}
