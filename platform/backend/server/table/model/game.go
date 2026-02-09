package model

import "time"

type Game struct {
	Id             int       `json:"id"`               // 遊戲id(PK)
	ServerInfoCode string    `json:"server_info_code"` // 遊戲的server code
	Name           string    `json:"name"`             // 遊戲名稱
	Code           string    `json:"code"`             // 遊戲代碼
	Type           int       `json:"type"`             // 遊戲類型
	State          int16     `json:"state"`            // 遊戲狀態
	CalState       int16     `json:"cal_state"`        // 遊戲狀態
	Image          string    `json:"image"`            // 遊戲圖片base64
	H5Link         string    `json:"h5_link"`          // 前端遊戲位置
	RoomNumber     int       `json:"room_number"`      // 房間數量
	TableNumber    int       `json:"table_number"`     // 桌子數量
	CreateTime     time.Time `json:"create_time"`      // 創建時間
	UpdateTime     time.Time `json:"update_time"`      // 更新時間
}

func (p *Game) GameResponseOutput() map[string]interface{} {
	temp := make(map[string]interface{}, 0)

	temp["id"] = p.Id
	temp["name"] = p.Name
	temp["code"] = p.Code
	temp["state"] = p.State
	// temp["image"] = p.Image

	return temp
}

type GameSlice []*Game

func (a GameSlice) Len() int {
	return len(a)
}
func (a GameSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a GameSlice) Less(i, j int) bool {
	return a[j].Id > a[i].Id
}
