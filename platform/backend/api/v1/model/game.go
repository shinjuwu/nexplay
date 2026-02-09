package model

import "definition"

type GetGameResponse struct {
	Id    int    `json:"id"`    // 遊戲id(PK)
	Code  string `json:"code"`  // 遊戲代碼
	State int    `json:"state"` // 遊戲狀態
}

type SetGameStateRequest struct {
	GameId int   `json:"game_id"` // 遊戲id(PK)
	State  int16 `json:"state"`   // 遊戲狀態
}

func (r *SetGameStateRequest) CheckParams() int {
	if r.GameId <= definition.GAME_ID_ALL || r.State < definition.GAME_STATE_OFFLINE || r.State > definition.GAME_STATE_MAINTAIN {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}

// Plinko 球倍率设定相关结构

type GetPlinkoballMaxOddsRequest struct {
	AgentName string `json:"agentName" form:"agentName"` // 代理商名称
	GameId    int    `json:"gameId" form:"gameId"`       // 游戏ID，应该是3003
}

type GetPlinkoballMaxOddsResponse struct {
	AgentName string  `json:"agentName"` // 代理商名称
	GameId    int     `json:"gameId"`    // 游戏ID
	MaxOdds   float64 `json:"maxOdds"`   // 最大倍率
}

type SetPlinkoballMaxOddsRequest struct {
	AgentName string  `json:"agentName"` // 代理商名称
	GameId    int     `json:"gameId"`    // 游戏ID，应该是3003
	MaxOdds   float64 `json:"maxOdds"`   // 最大倍率
}

func (r *GetPlinkoballMaxOddsRequest) CheckParams() int {
	if r.AgentName == "" || r.GameId <= 0 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}

func (r *SetPlinkoballMaxOddsRequest) CheckParams() int {
	if r.AgentName == "" || r.GameId <= 0 || r.MaxOdds <= 0 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}
