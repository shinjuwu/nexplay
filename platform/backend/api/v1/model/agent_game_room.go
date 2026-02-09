package model

type GetAgentGameRoomListRequest struct {
	AgentId int `json:"agent_id"` // 代理id
	GameId  int `json:"game_id"`  // 遊戲id
}

type GetAgentGameRoomResponse struct {
	AgentId    int    `json:"agent_id"`     // 代理id
	AgentName  string `json:"agent_name"`   // 代理名稱
	GameId     int    `json:"game_id"`      // 遊戲id
	GameRoomId int    `json:"game_room_id"` // 遊戲房間id
	RoomType   int    `json:"room_type"`    // 遊戲房間類型
	State      int16  `json:"state"`        // 代理遊戲狀態
}

type SetAgentGameRoomStateGameRoomRequest struct {
	AgentId    int   `json:"agent_id"`     // 代理id
	GameRoomId int   `json:"game_room_id"` // 遊戲房間id
	State      int16 `json:"state"`        // 代理遊戲狀態
}

type SetAgentGameRoomStateRequest struct {
	AgentId int                                     `json:"agent_id"`         // 代理id
	GameId  int                                     `json:"game_id"`          // 遊戲id
	List    []*SetAgentGameRoomStateGameRoomRequest `json:"list" form:"list"` // 代理遊戲房間狀態
}
