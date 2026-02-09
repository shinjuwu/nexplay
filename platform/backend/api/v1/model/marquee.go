package model

import (
	"backend/pkg/utils"
	"definition"
	"time"
)

type GetMarqueeResponse struct {
	Id         string    `json:"id"`          // uuid
	Lang       string    `json:"lang"`        // 語系
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

func NewEmptyGetMarqueeResponse() *GetMarqueeResponse {
	return &GetMarqueeResponse{}
}

/* ----------------------------------------------------------- */

type GetMarqueeRequest struct {
	Id string `json:"id"`
}

func NewEmptyGetMarqueeRequest() *GetMarqueeRequest {
	return &GetMarqueeRequest{}
}

func (p *GetMarqueeRequest) CheckParams() bool {
	return p.Id != ""
}

/* ----------------------------------------------------------- */

type CreateMarqueeRequest struct {
	Lang      string    `json:"lang"`       // 語系
	Type      int       `json:"type"`       // 跑馬燈類型(1: 系統, 2: 活動)
	Order     int       `json:"order"`      // 順序
	Freq      int       `json:"freq"`       // 播放頻率(每?秒播放一次)
	Content   string    `json:"content"`    // 內文
	StartTime time.Time `json:"start_time"` // 開始時間
	EndTime   time.Time `json:"end_time"`   // 結束時間
}

func NewEmptyCreateMarqueeRequest() *CreateMarqueeRequest {
	return &CreateMarqueeRequest{}
}

func (p *CreateMarqueeRequest) CheckParams() bool {
	// 開始時間不能在結束時間之後
	// int 型態參數一律不能小於1 (0 為 DB 預設值)
	// 順序(order)數字範圍為 1~9999
	// 頻率單位為秒(sec)
	// 內容不可為空
	if !p.EndTime.After(p.StartTime) ||
		p.Lang == "" ||
		p.Type <= 0 ||
		p.Type > 3 ||
		p.Order < definition.MARQUEE_ORDER_MIN ||
		p.Order > definition.MARQUEE_ORDER_MAX ||
		p.Freq < 1 ||
		p.Content == "" || utils.WordLength(p.Content) > 80 {
		return false
	}
	return true
}

/* ----------------------------------------------------------- */

type UpdateMarqueeRequest struct {
	Id        string    `json:"id"`         // uuid
	Lang      string    `json:"lang"`       // 語系
	Type      int       `json:"type"`       // 跑馬燈類型(1: 系統, 2: 活動)
	Order     int       `json:"order"`      // 順序
	Freq      int       `json:"freq"`       // 播放頻率(每?秒播放一次)
	Content   string    `json:"content"`    // 內文
	IsEnabled bool      `json:"is_enabled"` // 是否啟動
	StartTime time.Time `json:"start_time"` // 開始時間
	EndTime   time.Time `json:"end_time"`   // 結束時間
}

func NewEmptyUpdateMarqueeRequest() *UpdateMarqueeRequest {
	return &UpdateMarqueeRequest{}
}

func (p *UpdateMarqueeRequest) CheckParams() bool {
	// 開始時間不能在結束時間之後
	// int 型態參數一律不能小於1 (0 為 DB 預設值)
	// 順序(order)數字範圍為 1~9999
	// 頻率單位為秒(sec)
	// 內容不可為空
	if !p.EndTime.After(p.StartTime) ||
		p.Lang == "" ||
		p.Type <= 0 ||
		p.Type > 3 ||
		p.Order < definition.MARQUEE_ORDER_MIN ||
		p.Order > definition.MARQUEE_ORDER_MAX ||
		p.Freq < 1 ||
		p.Content == "" || utils.WordLength(p.Content) > 80 ||
		p.Id == "" {
		return false
	}
	return true
}

/* ----------------------------------------------------------- */

type DeleteMarqueeRequest struct {
	Id string `json:"id"` // uuid
}

func NewEmptyDeleteMarqueeRequest() *DeleteMarqueeRequest {
	return &DeleteMarqueeRequest{}
}

func (p *DeleteMarqueeRequest) CheckParams() bool {
	return p.Id != ""
}
