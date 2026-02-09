package model

import "time"

type GameRoom struct {
	Id         int       `json:"id"`          // 房間id(PK)
	Name       string    `json:"name"`        // 房間名稱
	State      int16     `json:"state"`       // 房間狀態
	GameId     int       `json:"game_id"`     // 遊戲id(FK)
	CreateTime time.Time `json:"create_time"` // 創建時間
	UpdateTime time.Time `json:"update_time"` // 更新時間
	RoomType   int       `json:"room_type"`   // 房間類型
}
