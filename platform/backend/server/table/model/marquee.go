package model

import "time"

type Marquee struct {
	Id         string    `json:"id"`          // uuid
	Lang       int       `json:"lang"`        // 語系
	Type       int       `json:"type"`        // 跑馬燈類型(1: 系統, 2: 活動)
	Order      int       `json:"order"`       // 順序
	Freq       int       `json:"freq"`        // 播放頻率(每?秒播放一次)
	IsEnabled  bool      `json:"is_enabled"`  // 是否啟動
	IsOpen     bool      `json:"is_open"`     // 是否開啟
	Content    string    `json:"content"`     // 內文
	StartTime  time.Time `json:"start_time"`  // 開始時間
	EndTime    time.Time `json:"end_time"`    // 結束時間
	CreateTime time.Time `json:"create_time"` // 創建時間
	UpdateTime time.Time `json:"update_time"` // 更新時間
}

func NewEmptyMarquee() *Marquee {
	return &Marquee{}
}

func (p *Marquee) CheckParams() bool {

	// 開始時間不能在結束時間之後
	// int 型態參數一律不能小於1 (0 為 DB 預設值)
	// 順序(order)數字範圍為 1~9999
	// 頻率單位為秒(sec)
	// 內容不可為空
	if !p.EndTime.After(p.StartTime) || p.Lang < 0 || p.Type <= 0 || p.Type > 3 || p.Order < 1 ||
		p.Order > 9999 || p.Freq < 1 || p.Content == "" {
		return false
	}

	return true
}
