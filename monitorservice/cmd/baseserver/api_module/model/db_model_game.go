package model

type Game struct {
	ID       string      `json:"id"` // 平台ID
	List     string      `json:"list"`
	GameList []*GameList `json:"game_list"`
}

type GameList struct {
	ID   int    `json:"id"`   // 遊戲ID
	Name string `json:"name"` // 遊戲名稱
	Code string `json:"code"` // uni識別碼
	Type int    `json:"type"` // 遊戲類型
}
