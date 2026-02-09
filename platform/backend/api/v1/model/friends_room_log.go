package model

import (
	"definition"
	"time"
)

type GetFriendRoomLogListRequest struct {
	AgentId  int    `json:"agent_id"` // 代理編號
	GameId   int    `json:"game_id"`  // 遊戲編號
	RoomId   string `json:"room_id"`  // 房間編號
	Username string `json:"username"` // 玩家名稱
	DateTimeTableRequest
}

func (g *GetFriendRoomLogListRequest) CheckParams() int {
	if g.AgentId < definition.AGENT_ID_ALL ||
		g.GameId < definition.GAME_ID_ALL {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	if code := g.CheckDateTimeTableRequest(); code != definition.ERROR_CODE_SUCCESS {
		return code
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetFriendRoomLogListResponse struct {
	Id         string    `json:"id"`          // 訂單號
	AgentId    int       `json:"agent_id"`    // 代理編號
	AgentName  string    `json:"agent_name"`  // 代理名稱
	GameId     int       `json:"game_id"`     // 遊戲編號
	RoomId     string    `json:"room_id"`     // 房間編號
	UserId     int       `json:"user_id"`     // 玩家編號
	Username   string    `json:"username"`    // 玩家名稱
	Tax        float64   `json:"tax"`         // 房間抽水
	Taxpercent float64   `json:"taxpercent"`  // 房間抽水%數
	CreateTime time.Time `json:"create_time"` // 建立時間
	EndTime    time.Time `json:"end_time"`    // 結束時間
	Detail     string    `json:"detail"`      // 詳細資訊json格式(各遊戲不同)
}
