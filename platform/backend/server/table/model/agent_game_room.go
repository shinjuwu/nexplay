package model

type AgentGameRoom struct {
	AgentId    int   `json:"agent_id"`
	GameRoomId int   `json:"game_room_id"`
	State      int16 `json:"state"`
}
