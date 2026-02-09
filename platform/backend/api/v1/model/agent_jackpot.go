package model

import (
	"definition"
	"time"
)

type GetAgentJackpotListRequest struct {
	AgentId int `json:"agent_id"` // 代理id
}

func (r *GetAgentJackpotListRequest) CheckParams() int {
	if r.AgentId < definition.AGENT_ID_ALL {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}

type GetAgentJackpotListResponse struct {
	Id               int       `json:"id"`                 // 代理id
	Name             string    `json:"name"`               // 代理名稱
	ChildAgentCount  int       `json:"child_agent_count"`  // 子代理數量
	JackpotStatus    int       `json:"jackpot_status"`     // jackpot 狀態(0:不參加, 1:參加, 2:限時參加)
	JackpotStartTime time.Time `json:"jackpot_start_time"` // jackpot 開始時間
	JackpotEndTime   time.Time `json:"jackpot_end_time"`   // jackpot 結束時間
}

type SetAgentJackpotRequest struct {
	AgentId          int       `json:"agent_id"`           // 代理id
	JackpotStatus    int       `json:"jackpot_status"`     // jackpot 狀態(0:不參加, 1:參加, 2:限時參加)
	JackpotStartTime time.Time `json:"jackpot_start_time"` // jackpot 開始時間
	JackpotEndTime   time.Time `json:"jackpot_end_time"`   // jackpot 結束時間
}

func (r *SetAgentJackpotRequest) CheckParams() int {
	if r.AgentId <= definition.AGENT_ID_ALL ||
		r.JackpotEndTime.Before(r.JackpotStartTime) ||
		r.JackpotStatus < 0 || r.JackpotStatus > 2 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	return definition.ERROR_CODE_SUCCESS
}
