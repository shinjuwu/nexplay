package model

import (
	"time"
)

type GameUsers struct {
	Id               int64     `json:"id"`                // user id(與game server 同步)
	AgentId          int       `json:"agent_id"`          // 代理編號
	OriginalUsername string    `json:"original_username"` // 用戶原平台帳號
	Username         string    `json:"username"`          // 用戶帳號(after mapping)
	UserMetadata     string    `json:"user_metadata"`     // 用戶基本資料
	TemporaryCoin    float64   `json:"temporary_coin"`    // 暫存遊戲幣(update in runtime)
	SumCoinIn        float64   `json:"sum_coin_in"`       // 累積遊戲幣轉入
	SumCoinOut       float64   `json:"sum_coin_out"`      // 累積遊戲幣轉出
	IsEnable         bool      `json:"is_enable"`         // 是否開啟
	CreateTime       time.Time `json:"create_time"`       // 創建時間
	UpdateTime       time.Time `json:"update_time"`       // 更新時間
	DisabledTime     time.Time `json:"disabled_time"`     // 關閉時間(可預先設定帳號關閉時間)
}
