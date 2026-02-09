package model

import "time"

type MarqueeSettingDataResponse struct {
	Id    string `json:"id"`    // uuid
	Lang  string `json:"lang"`  // 語系
	Type  int    `json:"type"`  // 跑馬燈類型(1: 系統, 2: 活動)
	Order int    `json:"order"` // 順序
	Freq  int    `json:"freq"`  // 播放頻率(每?秒播放一次)
	// IsEnabled  bool      `json:"is_enabled"`  // 是否啟動
	// IsOpen     bool      `json:"is_open"`     // 是否開啟
	Content   string    `json:"content"`    // 內文
	StartTime time.Time `json:"start_time"` // 開始時間
	EndTime   time.Time `json:"end_time"`   // 結束時間
	// CreateTime time.Time `json:"create_time"` // 創建時間
	// UpdateTime time.Time `json:"update_time"` // 更新時間
}
