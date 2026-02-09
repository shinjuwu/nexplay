package model

type AgentGame struct {
	AgentId int   `json:"agent_id"`
	GameId  int   `json:"game_id"`
	State   int16 `json:"state"`
}
