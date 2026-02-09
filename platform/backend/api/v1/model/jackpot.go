package model

import (
	"backend/pkg/utils"
	"definition"
	"time"
)

type GetJackpotTokenListRequest struct {
	AgentId        int    `json:"agent_id"`                               // 代理id
	Username       string `json:"username"`                               // 玩家名稱
	TokenId        string `json:"token_id"`                               // 代幣id
	Lognumber      string `json:"lognumber"`                              // 局號
	TimezoneOffset int    `json:"timezone_offset" form:"timezone_offset"` // UTC+0時間與本地時間差(分鐘) default(0)
	DateTimeTableRequest
}

func (r *GetJackpotTokenListRequest) CheckParams(reportTimeMinuteIncrement, reportTimeRange, reportTimeBeforeDays int) int {
	if r.TokenId == "" && r.Lognumber == "" {
		if code := r.CheckDateTimeTableRequest(); code != definition.ERROR_CODE_SUCCESS {
			return code
		}

		if r.StartTime.Second() > 0 ||
			r.EndTime.Second() > 0 ||
			r.StartTime.Minute()%reportTimeMinuteIncrement != 0 ||
			r.EndTime.Minute()%reportTimeMinuteIncrement != 0 ||
			r.TimezoneOffset < -840 || r.TimezoneOffset > 720 {
			return definition.ERROR_CODE_ERROR_REQUEST_DATA
		} else if r.EndTime.Sub(r.StartTime) > time.Duration(reportTimeRange)*time.Hour {
			return definition.ERROR_CODE_ERROR_TIME_RANGE
		} else {
			todayUTC := utils.GetTimeNowUTCTodayTime().Add(time.Duration(r.TimezoneOffset * int(time.Minute)))
			beforeUTCDays := todayUTC.AddDate(0, 0, -reportTimeBeforeDays)
			if r.StartTime.Before(beforeUTCDays) {
				return definition.ERROR_CODE_ERROR_TIME_BEFORE_DAYS
			}
		}

		if r.AgentId < definition.AGENT_ID_ALL {
			return definition.ERROR_CODE_ERROR_REQUEST_DATA
		}
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetJackpotTokenListResponse struct {
	Id              string    `json:"id"`                // 紀錄id
	TokenId         string    `json:"token_id"`          // 代幣id
	AgentId         int       `json:"agent_id"`          // 代理id
	AgentName       string    `json:"agent_name"`        // 代理名稱
	Username        string    `json:"username"`          // 玩家名稱
	SourceLognumber string    `json:"source_lognumber"`  // 代幣來源局號
	SourceBetId     string    `json:"source_bet_id"`     // 代幣來源玩家遊戲訂單號
	JpBet           float64   `json:"jp_bet"`            // 代幣分數
	UsageCount      int       `json:"usage_count"`       // 使用次數
	Creator         string    `json:"creator"`           // 建立者
	Info            string    `json:"info"`              // 備註
	Status          int       `json:"status"`            // 訂單狀態
	ErrorCode       int       `json:"error_code"`        // 錯誤碼
	TokenCreateTime time.Time `json:"token_create_time"` // 代幣產生時間
}

/*---------------------------------------*/

type CreateJackpotTokenRequest struct {
	AgentId  int     `json:"agent_id"` // 代理id
	Username string  `json:"username"` // 玩家名稱
	JpBet    float64 `json:"jp_bet"`   // 代幣分數
	Info     string  `json:"info"`     // 備註
}

func (r *CreateJackpotTokenRequest) CheckParams() int {
	if r.AgentId <= definition.AGENT_ID_ALL ||
		r.Username == "" ||
		r.JpBet < 0 || r.JpBet > 1000000 ||
		r.Info == "" || utils.WordLength(r.Info) > 100 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	return definition.ERROR_CODE_SUCCESS
}

/*---------------------------------------*/

type GetJackpotListRequest struct {
	AgentId        int    `json:"agent_id"`                               // 代理id
	Username       string `json:"username"`                               // 玩家名稱
	Lognumber      string `json:"lognumber"`                              // 局號
	TokenId        string `json:"token_id"`                               // 代幣id
	TimezoneOffset int    `json:"timezone_offset" form:"timezone_offset"` // UTC+0時間與本地時間差(分鐘) default(0)
	DateTimeTableRequest
}

func (r *GetJackpotListRequest) CheckParams(reportTimeMinuteIncrement, reportTimeRange, reportTimeBeforeDays int) int {
	if r.TokenId == "" && r.Lognumber == "" {
		if code := r.CheckDateTimeTableRequest(); code != definition.ERROR_CODE_SUCCESS {
			return code
		}

		if r.StartTime.Second() > 0 ||
			r.EndTime.Second() > 0 ||
			r.StartTime.Minute()%reportTimeMinuteIncrement != 0 ||
			r.EndTime.Minute()%reportTimeMinuteIncrement != 0 ||
			r.TimezoneOffset < -840 || r.TimezoneOffset > 720 {
			return definition.ERROR_CODE_ERROR_REQUEST_DATA
		} else if r.EndTime.Sub(r.StartTime) > time.Duration(reportTimeRange)*time.Hour {
			return definition.ERROR_CODE_ERROR_TIME_RANGE
		} else {
			todayUTC := utils.GetTimeNowUTCTodayTime().Add(time.Duration(r.TimezoneOffset * int(time.Minute)))
			beforeUTCDays := todayUTC.AddDate(0, 0, -reportTimeBeforeDays)
			if r.StartTime.Before(beforeUTCDays) {
				return definition.ERROR_CODE_ERROR_TIME_BEFORE_DAYS
			}
		}

		if r.AgentId < definition.AGENT_ID_ALL {
			return definition.ERROR_CODE_ERROR_REQUEST_DATA
		}
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetJackpotListResponse struct {
	BetId           string    `json:"bet_id"`            // 訂單號
	Lognumber       string    `json:"lognumber"`         // 局號
	TokenId         string    `json:"token_id"`          // 代幣id
	AgentId         int       `json:"agent_id"`          // 代理id
	AgentName       string    `json:"agent_name"`        // 代理名稱
	Username        string    `json:"username"`          // 玩家名稱
	JpBet           float64   `json:"jp_bet"`            // 代幣分數
	TokenCreateTime time.Time `json:"token_create_time"` // 代幣產生時間
	PrizeScore      float64   `json:"prize_score"`       // JP中獎分數
	PrizeItem       int       `json:"prize_item"`        // 中獎項目(4:參加獎、3:小獎、2:中獎、1:大獎)
	WinningTime     time.Time `json:"winning_time"`      // 中獎時間
	ShowPool        float64   `json:"show_pool"`         // 公告獎池
	RealPool        float64   `json:"real_pool"`         // 真實獎池
	IsRobot         int       `json:"is_robot"`          // 1:是 0:否
}

/*---------------------------------------*/

type GetJackpotPoolDataResponse struct {
	ShowPool          float64 `json:"show_pool"`            // 公告獎池
	RealPool          float64 `json:"real_pool"`            // 真實獎池
	ReservePool       float64 `json:"reserve_pool"`         // 預備獎池
	JpShowPool        float64 `json:"jp_show_pool"`         // jp公告獎池基礎值
	JpInjectWaterRate float64 `json:"jp_inject_water_rate"` // jp注水%數
}

/*---------------------------------------*/

type GetJackpotLeaderboardRequest struct {
	AgentId int `json:"agent_id"` // 代理id
}

func (r *GetJackpotLeaderboardRequest) CheckParams() int {
	if r.AgentId < definition.AGENT_ID_ALL {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetJackpotLeaderboardResponse struct {
	AgentId   int     `json:"agent_id"`   // 代理id
	AgentName string  `json:"agent_name"` // 代理名稱
	UserId    int     `json:"user_id"`    // 玩家id
	Username  string  `json:"username"`   // 玩家帳號
	PlayNum   int     `json:"play_num"`   // 總遊戲場次
	TotalBet  float64 `json:"total_bet"`  // 總有效投注
	Win       int     `json:"win"`        // 總贏場次
	Lose      int     `json:"lose"`       // 總輸場次
	Score     float64 `json:"score"`      // 貢獻度
}
