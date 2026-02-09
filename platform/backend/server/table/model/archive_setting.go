package model

import "time"

type ArchiveSetting struct {
	LastArchiveTime                     time.Time `json:"last_archive_time"`                          // DB 上次備份完成處理時間
	ArchiveFireTime                     int       `json:"archive_fire_time"`                          // DB 每日備份工作小時(0-23)
	DeleteRetainDataBatchCount          int       `json:"delete_retain_data_batch_count"`             // DB 備份資料分批刪除數量
	AgentGameUsersStatMinRetainDataDays int       `json:"agent_game_users_stat_min_retain_data_days"` // DB agent_game_users_stat_min 資料表保留天數
}
