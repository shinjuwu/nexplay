package model

import "definition"

type GetAgentGameListRequest struct {
	ServerSideTableRequest
	AgentId int   `json:"agent_id"` // 代理id
	GameId  int   `json:"game_id"`  // 遊戲id
	State   int16 `json:"state"`    // 代理遊戲狀態
}

func (r *GetAgentGameListRequest) CheckParams() int {
	if r.AgentId < definition.AGENT_ID_ALL ||
		r.GameId < definition.GAME_ID_ALL ||
		r.State < definition.GAME_STATE_ALL ||
		r.State > definition.GAME_STATE_MAINTAIN {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetAgentGameResponse struct {
	AgentId        int    `json:"agent_id"`         // 代理id
	AgentName      string `json:"agent_name"`       // 代理名稱
	AgentLevelCode string `json:"agent_level_code"` // 代理層級代碼
	GameId         int    `json:"game_id"`          // 遊戲id
	GameCode       string `json:"game_code"`        // 遊戲代碼
	State          int16  `json:"state"`            // 代理遊戲狀態
}

type SetAgentGameStateGameRequest struct {
	AgentId int `json:"agent_id"` // 代理id
	GameId  int `json:"game_id"`  // 遊戲id
}

type SetAgentGameStateRequest struct {
	List  []*SetAgentGameStateGameRequest `json:"list" form:"list"`
	State int16                           `json:"state" form:"state"` // 代理遊戲狀態
}

func (r *SetAgentGameStateRequest) CheckParams() int {
	if len(r.List) == 0 ||
		r.State < definition.GAME_STATE_OFFLINE ||
		r.State > definition.GAME_STATE_MAINTAIN {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	return definition.ERROR_CODE_SUCCESS
}
