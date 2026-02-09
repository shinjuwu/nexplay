package model

import "time"

type UserPlaylog struct {
	BetId      string    `json:"bet_id"`      // 注單號
	AgentId    int       `json:"agent_id"`    // 代理編號
	OrderId    string    `json:"order_id"`    // 遊戲局號
	Username   string    `json:"username"`    // 玩家帳號
	GameId     int       `json:"game_id"`     // 遊戲id
	RoomType   int       `json:"room_type"`   // 房間類型
	DeskId     int       `json:"desk_id"`     // 桌子號
	SeatId     int       `json:"seat_id"`     // 椅子號
	DeScore    float64   `json:"de_score"`    // 得分
	YaScore    float64   `json:"ya_score"`    // 壓分
	VaildScore float64   `json:"vaild_score"` // 有效投注
	Tax        float64   `json:"tax"`         // 抽水
	Bonus      float64   `json:"bonus"`       // 紅利
	BetTime    time.Time `json:"bet_time"`    // 遊戲開獎時間
	StartTime  time.Time `json:"start_time"`  // 遊戲開始時間
	EndTime    time.Time `json:"end_time"`    // 遊戲結束時間
}

// tip: json tag 跟其他tag風格不一樣是要跟其他家廠商一樣
type GameData struct {
	BetId         string    `json:"betId"`         // 注單號
	OrderId       string    `json:"orderId"`       // 遊戲局號
	Account       string    `json:"account"`       // 玩家帳號
	GameId        int       `json:"gameId"`        // 遊戲id(對應遊戲見附錄)
	RoomId        int       `json:"roomId"`        // 房間類型
	DeskId        int       `json:"deskId"`        // 桌子號
	SeatId        int       `json:"seatId"`        // 椅子號
	Validbet      float64   `json:"validBet"`      // 有效下注
	Bet           float64   `json:"bet"`           // 下注
	Win           float64   `json:"win"`           // 盈利
	Revenue       float64   `json:"revenue"`       // 抽水
	Bonus         float64   `json:"bonus"`         // 紅利
	Currency      string    `json:"currency"`      // 幣別
	GameBetTime   time.Time `json:"gameBetTime"`   // 遊戲開獎時間
	GameStartTime time.Time `json:"gameStartTime"` // 遊戲開始時間
	GameEndTime   time.Time `json:"gameEndTime"`   // 遊戲結束時間
}

func NewGameData() *GameData {
	return &GameData{}
}

func (p *GameData) TransUserPlayLog(data map[string]interface{}) {

}
