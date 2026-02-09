package model

import (
	"monitorservice/pkg/utils"
	"time"
)

type ServiceStatusRequest struct {
	Filter string `json:"filter"` // 指定查詢單一平台或是全部(全部: all, 其他: 平台唯一碼)
}

type ServiceStatus struct {
	Name    string `json:"name"`     // 平台碼(全部: all)
	SubName string `json:"sub_name"` // 副標(那裡的服務)
	Info    string `json:"info"`     // 說明
	Status  int    `json:"status"`   // 目前服務狀態(0:正常,1:異常)
}

type ServiceStatusResponse struct {
	StatusList []*ServiceStatus `json:"status_list"` // 狀態列表

}

//*****************************************************************************

type CoinInOutStatusRequest struct {
	Filter string `json:"filter"` // 指定查詢單一平台(平台唯一碼)
}

type CoinInOutStatus struct {
	Id         string    `json:"id"`          // 訂單號
	AgentId    int       `json:"agent_id"`    // 代理id
	AgentName  string    `json:"agent_name"`  // 代理名稱
	UserId     int       `json:"user_id"`     // 代理用戶id
	UserName   string    `json:"username"`    // 代理用戶帳號
	CreateTime time.Time `json:"create_time"` // 訂單時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
	Kind       int       `json:"kind"`        // 帳變類型
	ChangeSet  string    `json:"changeset"`   // 帳變內容
	Status     int       `json:"status"`      // 訂單狀態
	Info       string    `json:"info"`        // 備註
}

type CoinInOutStatusResponse struct {
	CoinInOutStatusList []*CoinInOutStatus `json:"coin_inout_status_list"` // 狀態列表
}

//*****************************************************************************

type AbnormalWinAndLoseStatusRequest struct {
	Filter string `json:"filter"` // 指定查詢單一平台(平台唯一碼)
}

type AbnormalWinAndLoseStatus struct {
	// Platform   string    `json:"platform"`
	LogNumber  string    `json:"lognumber"`
	AgentID    int       `json:"agent_id"`
	AgentName  string    `json:"agent_name"`
	UserID     int       `json:"user_id"`
	Username   string    `json:"username"`
	GameID     int       `json:"game_id"`
	GameName   string    `json:"game_name"`
	RoomType   int       `json:"room_type"`
	RoomName   string    `json:"room_name"`
	DeScore    float64   `json:"de_score"`
	Bonus      float64   `json:"bonus"`
	BetID      string    `json:"bet_id"`
	BetTime    time.Time `json:"bet_time"`
	CreateTime time.Time `json:"create_time"`
}

type AbnormalWinAndLoseStatusResponse struct {
	AbnormalWinAndLoseStatusList []*AbnormalWinAndLoseStatus `json:"abnormal_win_and_lose_status_list"` // 狀態列表
}

//*****************************************************************************

type PlatformRTPStatusRequest struct {
	Filter   string `json:"filter"`    // 指定查詢單一平台或是全部(平台唯一碼)
	TimeZone int    `json:"time_zone"` // 時間區域(UTC+X, X時間, 單位分鐘, ex: UTC+8 要傳入 -60*8=-480)
}

type GameRatioStat struct {
	LogTime   string  `json:"log_time"`
	GameID    int     `json:"game_id"`
	GameType  int     `json:"game_type"`
	GameName  string  `json:"game_name"`
	DE        float64 `json:"de"`
	YA        float64 `json:"ya"`
	Tax       float64 `json:"tax"`
	Bonus     float64 `json:"bonus"`
	PlayCount int     `json:"play_count"`
}

type PlatformGameRatioStat struct {
	GameRatioStatList []*GameRatioStat `json:"rtp_status_list"` // 狀態列表
}

type RTPStatus struct {
	RTPType   string  `json:"rtp_type"` // RTP 種類(month, week, day, gamename)
	Title     string  `json:"title"`    // 標題
	GameType  int     `json:"game_type"`
	CalST     string  `json:"cal_st"` // 計算開始時間(YYYYMMDDhh)，適配用戶端時間
	CalET     string  `json:"cal_et"` // 計算結束時間(YYYYMMDDhh)，適配用戶端時間
	DE        float64 `json:"de"`
	YA        float64 `json:"ya"`
	Tax       float64 `json:"tax"`
	Bonus     float64 `json:"bonus"`
	PlayCount int     `json:"play_count"`
}

func NewRTPStatus(rtpType, title string, gameType int, calST, calET string) *RTPStatus {
	return &RTPStatus{
		RTPType:   rtpType,
		Title:     title,
		GameType:  gameType,
		CalST:     calST,
		CalET:     calET,
		DE:        .0,
		YA:        .0,
		Tax:       .0,
		Bonus:     .0,
		PlayCount: 0,
	}
}

func (p *RTPStatus) CalProcess(de, ya, tax, bonus float64, playCount int) {
	p.DE = utils.DecimalAdd(p.DE, de)
	p.YA = utils.DecimalAdd(p.YA, ya)
	p.Tax = utils.DecimalAdd(p.Tax, tax)
	p.Bonus = utils.DecimalAdd(p.Bonus, bonus)
	p.PlayCount += playCount
}

type PlatformRTPStatusResponse struct {
	RTPStatusList      []*RTPStatus `json:"rtp_status_list"`       // 總和狀態列表(本日、本週、本月)
	RTPStatusDayList   []*RTPStatus `json:"rtp_status_day_list"`   // 遊戲狀態列表(本日)
	RTPStatusWeekList  []*RTPStatus `json:"rtp_status_week_list"`  // 遊戲狀態列表(本週)
	RTPStatusMonthList []*RTPStatus `json:"rtp_status_month_list"` // 遊戲狀態列表(本月)
}

//*****************************************************************************
