package model

type GetGameServerStateResponse struct {
	State int `json:"state"` // state: 狀態 (1: 開啟, 2: 關閉; 預設: 關閉)
}
