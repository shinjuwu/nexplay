package model

import "time"

type JackpotTokenLog struct {
	TokenId         string    `json:"token_id"`         // token id
	AgentId         int       `json:"agent_id"`         // 代理id
	LevelCode       string    `json:"agent_level_code"` // 代理層級碼
	UserId          int       `json:"user_id"`          // 用戶id
	Username        string    `json:"username"`         // 用戶名稱
	JpBet           float64   `json:"jp_bet"`           // jp bet
	TokenCreateTime time.Time `json:"create_time"`      // token 建立時間
	SourceGameId    int       `json:"source_game_id"`   // 來源遊戲id
	SourceLognumber string    `json:"source_lognumber"` // 來源局號
	SourceBetId     string    `json:"source_bet_id"`    // 來源訂單號
}

type JackpotLogRequest struct {
	Lognumber       string  `json:"lognumber"`         // 局號
	AgentId         int     `json:"agent_id"`          // 代理id
	UserId          int     `json:"user_id"`           // 用戶id
	Username        string  `json:"username"`          // 用戶名稱
	TokenId         string  `json:"token_id"`          // token id
	JpBet           float64 `json:"jp_bet"`            // jp bet
	TokenCreateTime int64   `json:"token_create_time"` // token 建立時間
	PrizeScore      float64 `json:"prize_score"`       // JP中獎分數
	PrizeItem       int     `json:"prize_item"`        // 中獎項目(4:參加獎、3:小獎、2:中獎、1:大獎)
	WinningTime     int64   `json:"winning_time"`      // 中獎時間
	ShowPool        float64 `json:"show_pool"`         // 公告獎池
	RealPool        float64 `json:"real_pool"`         // 真實獎池
	IsRobot         int     `json:"is_robot"`          // 1:是 0:否
}
