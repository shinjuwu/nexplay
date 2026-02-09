package model

type FriendRoomLogRequest struct {
	RoomId        string  `json:"room_id"`     // 房間編號
	GameId        int     `json:"game_id"`     // 遊戲id
	AgentId       int     `json:"agent_id"`    // 房主的代理id
	UserId        int     `json:"user_id"`     // 房主的玩家id
	Username      string  `json:"username"`    // 房主的玩家名稱
	CreateTime    int64   `json:"create_time"` // 創建時間 timestamp milliseconds
	EndTime       int64   `json:"end_time"`    // 解桌時間 timestamp milliseconds
	Tax           float64 `json:"tax"`         // 總抽水
	TaxPercentage float64 `json:"taxpercent"`  // 遊戲抽水%數
	Detail        string  `json:"detail"`      // 房間詳情(json字串，依照各遊戲設定)
}
