package model

type ServerInfo struct {
	Code           string    `json:"code" db:"code"`
	Ip             string    `json:"ip" db:"ip"`
	Addresses      Addresses `json:"addresses" db:"addresses"`
	AddressesBytes []byte    `json:"-"`
	IsEnabled      bool      `json:"is_enabled" db:"is_enabled"`
}

type Addresses struct {
	Notification string `json:"notification" db:"notification"`
}

type ServerInfoSetting struct {
	CommonReportTimeRange            int `json:"common_report_time_range" mapstructure:"common_report_time_range"`                         // 一般報表可查詢時間的範圍長度(天)
	CommonReportTimeBeforeDays       int `json:"common_report_time_before_days" mapstructure:"common_report_time_before_days"`             // 一般報表可查詢多久前的時間(天)
	CommonReportTimeMinuteIncrement  int `json:"common_report_time_minute_increment" mapstructure:"common_report_time_minute_increment"`   // 一般報表查詢的分鐘間格(分)
	WinloseReportTimeRange           int `json:"winlose_report_time_range" mapstructure:"winlose_report_time_range"`                       // 輸贏報表可查詢時間範圍長度(分鐘)
	WinloseReportTimeBeforeDays      int `json:"winlose_report_time_before_days" mapstructure:"winlose_report_time_before_days"`           // 輸贏報表可查詢多久前的時間(天)
	EarningReportTimeMinuteIncrement int `json:"earning_report_time_minute_increment" mapstructure:"earning_report_time_minute_increment"` // 業績報表查詢的分鐘間格(分)
}

// default
/*
{
  "common_report_time_range": 31,
  "winlose_report_time_range": 12,
  "common_report_time_before_days": 90,
  "winlose_report_time_before_days": 7,
  "common_report_time_minute_increment": 5,
  "earning_report_time_minute_increment": 15
}
*/
func NewServerInfoSetting() *ServerInfoSetting {
	return &ServerInfoSetting{
		CommonReportTimeRange:            31,
		WinloseReportTimeRange:           12,
		CommonReportTimeBeforeDays:       90,
		WinloseReportTimeBeforeDays:      7,
		CommonReportTimeMinuteIncrement:  5,
		EarningReportTimeMinuteIncrement: 15,
	}
}
