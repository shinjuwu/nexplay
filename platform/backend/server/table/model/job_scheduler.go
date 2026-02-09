package model

type JobScheduler struct {
	Id           string `json:"id"`             // 功能代碼(PK)
	Spec         string `json:"spec"`           // 名稱
	Info         string `json:"info"`           // api 路徑
	TriggerFunc  string `json:"trigger_func"`   // 對應驅動
	IsEnabled    bool   `json:"is_enabled"`     // 是否開放
	ExecLimit    int    `json:"exec_limit"`     // 指定執行次數
	LastSyncDate string `json:"last_sync_date"` // 最後同步日期辨識用字串
}

func NewEmptyJobScheduler() *JobScheduler {
	return &JobScheduler{}
}
