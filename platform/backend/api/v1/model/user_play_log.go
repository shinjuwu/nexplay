package model

import (
	"backend/pkg/utils"
	"definition"
	"time"
)

// 輸贏報表
type GetUserPlayLogListRequest struct {
	DateTimeTableRequest
	AgentId        int    `json:"agent_id" form:"agent_id"`                 // 代理編號 default(0)
	GameId         int    `json:"game_id" form:"game_id"`                   // 遊戲編號 default(0)
	RoomType       int    `json:"room_type" form:"room_type"`               // 房間類型 default(0)
	UserName       string `json:"username" form:"username"`                 // 代理玩家名稱
	LogNumber      string `json:"lognumber" form:"lognumber"`               // 局號
	SingleWalletId string `json:"single_wallet_id" form:"single_wallet_id"` // 單一錢包識別碼
	BetId          string `json:"bet_id" form:"bet_id"`                     // 注單號
	RoomId         string `json:"room_id" form:"room_id"`                   // 房間編號(邀請碼)
	TimezoneOffset int    `json:"timezone_offset" form:"timezone_offset"`   // UTC+0時間與本地時間差(分鐘) default(0)
}

func (r *GetUserPlayLogListRequest) CheckParams(reportTimeMinuteIncrement, reportTimeRange, reportTimeBeforeDays int) int {
	if r.BetId == "" &&
		r.LogNumber == "" &&
		r.SingleWalletId == "" &&
		r.RoomId == "" {
		if code := r.CheckDateTimeTableRequest(); code != definition.ERROR_CODE_SUCCESS {
			return code
		}

		if r.StartTime.Second() > 0 ||
			r.EndTime.Second() > 0 ||
			r.StartTime.Minute()%reportTimeMinuteIncrement != 0 ||
			r.EndTime.Minute()%reportTimeMinuteIncrement != 0 ||
			r.TimezoneOffset < -840 || r.TimezoneOffset > 720 {
			return definition.ERROR_CODE_ERROR_REQUEST_DATA
		} else if r.EndTime.Sub(r.StartTime) > time.Duration(reportTimeRange)*time.Minute {
			return definition.ERROR_CODE_ERROR_TIME_RANGE
		} else {
			todayUTC := utils.GetTimeNowUTCTodayTime().Add(time.Duration(r.TimezoneOffset * int(time.Minute)))
			beforeUTCDays := todayUTC.AddDate(0, 0, -reportTimeBeforeDays)
			if r.StartTime.Before(beforeUTCDays) {
				return definition.ERROR_CODE_ERROR_TIME_BEFORE_DAYS
			}
		}

		if r.AgentId < definition.AGENT_ID_ALL || r.GameId < definition.GAME_ID_ALL ||
			r.RoomType < definition.ROOM_TYPE_ALL || r.RoomType > definition.ROOM_TYPE_PERFECTION {
			return definition.ERROR_CODE_ERROR_REQUEST_DATA
		}
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetUserPlayLogResponse struct {
	AgentId            int       `json:"agent_id"`              // 代理編號
	AgentName          string    `json:"agent_name"`            // 代理名稱
	BetTime            time.Time `json:"bet_time"`              // 遊戲結算時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
	UserName           string    `json:"username"`              // 代理用戶帳號
	GameId             int       `json:"game_id"`               // 遊戲編號
	RoomType           int       `json:"room_type"`             // 房間類型編號
	DeskId             int       `json:"desk_id"`               // 桌子編號
	SeatId             int       `json:"seat_id"`               // 座位編號
	ValidScore         float64   `json:"valid_score"`           // 有效投注
	StartScore         float64   `json:"start_score"`           // 玩家壓注前遊戲分
	EndScore           float64   `json:"end_score"`             // 玩家壓注後遊戲分
	DeScore            float64   `json:"de_score"`              // 總得遊戲分
	YaScore            float64   `json:"ya_score"`              // 總壓遊戲分
	LogNumber          string    `json:"lognumber"`             // 局號
	BetId              string    `json:"bet_id"`                // 注單號
	KillType           int       `json:"kill_type"`             // 殺放狀態(0. 一般、1. 基礎追殺、2. 放水、3. 控牌追殺、4. 單間追殺、5. 限制倍率、6. 禁開分數)
	Tax                float64   `json:"tax"`                   // 遊戲抽水
	Bonus              float64   `json:"bonus"`                 // 紅利
	JpInjectWaterRate  float64   `json:"jp_inject_water_rate"`  // jp注水%數
	JpInjectWaterScore float64   `json:"jp_inject_water_score"` // jp注水分數
	WalletLedgerId     string    `json:"wallet_ledger_id"`      // 單一錢包識別碼
	KillProb           float64   `json:"kill_prob"`             // 殺放設定機率
	KillLevel          int       `json:"kill_level"`            // 殺放層級 (預設值-1)
	RealPlayers        int       `json:"real_players"`          // 真實玩家人數 (預設值-1)
	RoomId             string    `json:"room_id"`               // 房間編號(邀請碼)
}

type GetBatchUserPlayLogListRequest struct {
	BetslipStatus int `json:"betslip_status"` // 訂單狀態(0:全部,1:正常,2:異常)
	GetUserPlayLogListRequest
	ServerSideTableRequest
}

func (r *GetBatchUserPlayLogListRequest) CheckParams(reportTimeMinuteIncrement, reportTimeRange, reportTimeBeforeDays int) int {
	if r.BetslipStatus < 0 || r.BetslipStatus > 2 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	if code := r.CheckServerSideTableRequest(); code != definition.ERROR_CODE_SUCCESS {
		return code
	}

	if code := r.GetUserPlayLogListRequest.CheckParams(
		reportTimeMinuteIncrement,
		reportTimeRange,
		reportTimeBeforeDays); code != definition.ERROR_CODE_SUCCESS {
		return code
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetBatchUserPlayLogListResponse struct {
	DataTablesResponse
	TotalKillGames            int     `json:"total_kill_games"`             // 總殺局數
	TotalDiveGames            int     `json:"total_dive_games"`             // 總放局數
	TotalPlayerWinGames       int     `json:"total_player_win_games"`       // 總玩家贏局數
	TotalValidScore           float64 `json:"total_valid_score"`            // 總有效投注
	TotalPlatformWinloseScore float64 `json:"total_platform_winlose_score"` // 平台總輸贏
	TotalJpInjectWaterScore   float64 `json:"total_jp_inject_water_score"`  // 總jp注水分數
}

//-------------------------------------------------------------------

type GetUserCreditLogListRequest struct {
	DateTimeTableRequest
	AgentId        int    `json:"agent_id" form:"agent_id"`               // 代理編號 default(0)
	GameId         int    `json:"game_id" form:"game_id"`                 // 遊戲編號 default(0)
	UserName       string `json:"username" form:"username"`               // 代理玩家名稱
	TimezoneOffset int    `json:"timezone_offset" form:"timezone_offset"` // UTC+0時間與本地時間差(分鐘) default(0)
}

func (r *GetUserCreditLogListRequest) CheckParams(reportTimeMinuteIncrement, reportTimeRange, reportTimeBeforeDays int) int {
	if code := r.CheckDateTimeTableRequest(); code != definition.ERROR_CODE_SUCCESS {
		return code
	}

	if r.StartTime.Second() > 0 ||
		r.EndTime.Second() > 0 ||
		r.StartTime.Minute()%reportTimeMinuteIncrement != 0 ||
		r.EndTime.Minute()%reportTimeMinuteIncrement != 0 ||
		r.TimezoneOffset < -840 || r.TimezoneOffset > 720 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	} else if r.EndTime.Sub(r.StartTime) > time.Duration(reportTimeRange)*time.Minute {
		return definition.ERROR_CODE_ERROR_TIME_RANGE
	} else {
		todayUTC := utils.GetTimeNowUTCTodayTime().Add(time.Duration(r.TimezoneOffset * int(time.Minute)))
		beforeUTCDays := todayUTC.AddDate(0, 0, -reportTimeBeforeDays)
		if r.StartTime.Before(beforeUTCDays) {
			return definition.ERROR_CODE_ERROR_TIME_BEFORE_DAYS
		}
	}

	if r.AgentId < definition.AGENT_ID_ALL || r.GameId < definition.GAME_ID_ALL {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	return definition.ERROR_CODE_SUCCESS
}

// 玩家帳變紀錄回傳資料
type GetUserCreditLogListResponse struct {
	AgentId          int       `json:"agent_id"`           // 代理編號 (輸贏報表欄位)
	AgentName        string    `json:"agent_name"`         // 代理名稱 (輸贏報表欄位)
	UserName         string    `json:"username"`           // 代理用戶帳號 (輸贏報表欄位)
	GameId           int       `json:"game_id"`            // 遊戲編號 (輸贏報表欄位)
	StartScore       float64   `json:"start_score"`        // 玩家壓注前遊戲分 (輸贏報表欄位)
	EndScore         float64   `json:"end_score"`          // 玩家壓注後遊戲分 (輸贏報表欄位)
	DeScore          float64   `json:"de_score"`           // 總得遊戲分 (輸贏報表欄位)
	YaScore          float64   `json:"ya_score"`           // 總壓遊戲分 (輸贏報表欄位)
	WalletKind       int       `json:"wallet_kind"`        // 帳變類型 (帳變欄位)
	WalletCoinAmount float64   `json:"wallet_coin_amount"` // 執行分數 (帳變欄位)
	WalletStatus     int       `json:"wallet_status"`      // 訂單狀態 (帳變欄位)
	WalletErrorCode  int       `json:"wallet_error_code"`  // 錯誤訊息代碼 (帳變欄位)
	WalletChangeSet  string    `json:"wallet_changeset"`   // 帳變內容 (帳變欄位)
	CreditId         string    `json:"credit_id"`          // 遊戲局號 / 上下分訂單號 (合併欄位)
	CreditTime       time.Time `json:"credit_time"`        // 遊戲結算時間 / 訂單時間(UTC+0, format: 2006-01-02T15:04:05.000Z) (合併欄位)
	CreditCode       string    `json:"credit_code"`        // 單一錢包識別碼 / JP 代幣ID (合併欄位)
	JPScore          float64   `json:"jp_score"`           // JP中獎分數
	JPItem           int       `json:"jp_item"`            // 中獎項目(4:參加獎、3:小獎、2:中獎、1:大獎)
	KillType         int       `json:"kill_type"`          // 殺放狀態 (輸贏報表欄位)
}

func (p *GetUserCreditLogListResponse) TranGameData(g *GetUserPlayLogResponse) {
	p.AgentId = g.AgentId
	p.AgentName = g.AgentName
	p.UserName = g.UserName
	p.GameId = g.GameId
	p.DeScore = g.DeScore
	p.YaScore = g.YaScore
	p.StartScore = g.StartScore
	p.EndScore = g.EndScore
	p.KillType = g.KillType
	p.CreditId = g.LogNumber
	p.CreditTime = g.BetTime
	p.CreditCode = "-"
	p.WalletKind = -1
	p.WalletCoinAmount = .0
	p.WalletStatus = -1
	p.WalletErrorCode = -1
	p.WalletChangeSet = "{}"
	p.JPScore = .0
	p.JPItem = -1
}

func (p *GetUserCreditLogListResponse) TranWalletData(g *GetWalletLedgerResponse) {
	p.AgentId = utils.ToInt(g.AgentId)
	p.AgentName = g.AgentName
	p.UserName = g.UserName
	p.GameId = -1
	p.DeScore = .0
	p.YaScore = .0
	p.StartScore = .0
	p.EndScore = .0
	p.KillType = -1
	p.CreditId = g.Id
	p.CreditTime = g.CreateTime
	p.CreditCode = g.SingleWalletId
	p.WalletKind = g.Kind
	p.WalletCoinAmount = g.CoinAmount
	p.WalletStatus = g.Status
	p.WalletErrorCode = g.ErrorCode
	p.WalletChangeSet = g.ChangeSet
	p.JPScore = .0
	p.JPItem = -1
}

func (p *GetUserCreditLogListResponse) TranJPData(g *GetJackpotListResponse) {
	p.AgentId = g.AgentId
	p.AgentName = g.AgentName
	p.UserName = g.Username
	p.GameId = -1
	p.DeScore = .0
	p.YaScore = .0
	p.StartScore = .0
	p.EndScore = .0
	p.KillType = -1
	p.CreditId = g.Lognumber
	p.CreditTime = g.WinningTime
	p.CreditCode = g.TokenId
	p.WalletKind = -1
	p.WalletCoinAmount = .0
	p.WalletStatus = -1
	p.WalletErrorCode = -1
	p.WalletChangeSet = "{}"
	p.JPScore = g.PrizeScore
	p.JPItem = g.PrizeItem
}
