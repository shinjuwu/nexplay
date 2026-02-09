package model

import "definition"

type GetRealtimeGameRatioRequest struct {
	GameId int `json:"game_id"` // 遊戲id 一次查詢一整個遊戲
}

func (p *GetRealtimeGameRatioRequest) CheckParams() int {
	if p.GameId <= 0 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	return definition.ERROR_CODE_SUCCESS
}

type RealtimeGameRatio struct {
	Ya          float64 `json:"ya"`           // 總壓注
	De          float64 `json:"de"`           // 總得分
	Tax         float64 `json:"tax"`          // 抽水
	PlayCount   int     `json:"play_count"`   // 遊戲局數
	St          int     `json:"st"`           // 資料計算開始時間(utc+0, timestamp)
	Et          int     `json:"et"`           // 資料計算結束時間(utc+0, timestamp)
	PlayerCount int     `josn:"player_count"` // 目前用戶數量
}

// use roomid as map index
// map[roomId]*RealtimeGameRatio
type GetRealtimeGameRatioResponse struct {
	Data map[int]*RealtimeGameRatio `json:"data"`
}
