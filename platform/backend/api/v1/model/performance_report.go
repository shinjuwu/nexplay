package model

import (
	"backend/pkg/utils"
	"definition"
	"time"
)

/**************************************************/

type GetPerformanceReportListRequest struct {
	AgentId        int       `json:"agent_id"`                               // 代理id
	StartTime      time.Time `json:"start_time"`                             // 指定開始時間
	EndTime        time.Time `json:"end_time"`                               // 指定結束時間
	TimezoneOffset int       `json:"timezone_offset" form:"timezone_offset"` // UTC+0時間與本地時間差(分鐘) default(0)
}

func NewGetPerformanceReportListRequest() *GetPerformanceReportListRequest {
	return &GetPerformanceReportListRequest{}
}

func (p *GetPerformanceReportListRequest) CheckParams(reportTimeMinuteIncrement, reportTimeRange, reportTimeBeforeDays int) int {

	if p.AgentId < definition.AGENT_ID_ALL ||
		p.StartTime.Second() > 0 ||
		p.EndTime.Second() > 0 ||
		p.StartTime.Minute()%reportTimeMinuteIncrement != 0 ||
		p.EndTime.Minute()%reportTimeMinuteIncrement != 0 ||
		p.TimezoneOffset < -840 || p.TimezoneOffset > 720 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	} else if p.EndTime.Sub(p.StartTime) > time.Duration(reportTimeRange)*24*time.Hour {
		return definition.ERROR_CODE_ERROR_TIME_RANGE
	} else {
		todayUTC := utils.GetTimeNowUTCTodayTime().Add(time.Duration(p.TimezoneOffset * int(time.Minute)))
		beforeUTCDays := todayUTC.AddDate(0, 0, -reportTimeBeforeDays)
		if p.StartTime.Before(beforeUTCDays) {
			return definition.ERROR_CODE_ERROR_TIME_BEFORE_DAYS
		}
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetPerformanceReportListResponse struct {
	AgentId               int     `json:"agentId"`                   // 代理id
	LevelCode             string  `json:"level_code"`                //
	AgentName             string  `json:"agent_name"`                // 代理名稱
	AgentCommission       int     `json:"agent_commission"`          // 代理分成
	Currency              string  `json:"currency"`                  // 幣種
	ToCoin                float64 `json:"to_coin"`                   // 幣種換算比值
	BetUser               int     `json:"bet_user"`                  // 總投注人數
	BetCount              int     `json:"bet_count"`                 // 注單數
	JackpotUser           int     `json:"jackpot_user"`              // jackpot中獎總人數
	JackpotCount          int     `json:"jackpot_count"`             // jackpot中獎總注單數
	SumYa                 float64 `json:"sum_ya"`                    // 總投注
	SumValidYa            float64 `json:"sum_valid_ya"`              // 有效投注
	SumDe                 float64 `json:"sum_de"`                    // 總派獎
	SumBonus              float64 `json:"sum_bonus"`                 // 紅利
	SumTax                float64 `json:"sum_tax"`                   // 抽水
	SumJpInjectWaterScore float64 `json:"sum_jp_inject_water_score"` // jp注水分數
	SumJpPrizeScore       float64 `json:"sum_jp_prize_score"`        // jp中獎分數
}

func NewGetPerformanceReportListResponse() *GetPerformanceReportListResponse {
	return &GetPerformanceReportListResponse{}
}

func NewGetPerformanceReportListResponseSlice() []*GetPerformanceReportListResponse {
	return make([]*GetPerformanceReportListResponse, 0)
}

/**************************************************/

type GetPerformanceReportRequest struct {
	AgentId        int       `json:"agent_id"`                               // 代理id
	StartTime      time.Time `json:"start_time"`                             // 指定開始時間
	EndTime        time.Time `json:"end_time"`                               // 指定結束時間
	TimezoneOffset int       `json:"timezone_offset" form:"timezone_offset"` // UTC+0時間與本地時間差(分鐘) default(0)
}

func NewGetPerformanceReportRequest() *GetPerformanceReportRequest {
	return &GetPerformanceReportRequest{}
}

func (p *GetPerformanceReportRequest) CheckParams(reportTimeMinuteIncrement, reportTimeRange, reportTimeBeforeDays int) int {

	if p.AgentId <= definition.AGENT_ID_ALL ||
		p.StartTime.Second() > 0 ||
		p.EndTime.Second() > 0 ||
		p.StartTime.Minute()%reportTimeMinuteIncrement != 0 ||
		p.EndTime.Minute()%reportTimeMinuteIncrement != 0 ||
		p.TimezoneOffset < -840 || p.TimezoneOffset > 720 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	} else if p.EndTime.Sub(p.StartTime) > time.Duration(reportTimeRange)*24*time.Hour {
		return definition.ERROR_CODE_ERROR_TIME_RANGE
	} else {
		todayUTC := utils.GetTimeNowUTCTodayTime().Add(time.Duration(p.TimezoneOffset * int(time.Minute)))
		beforeUTCDays := todayUTC.AddDate(0, 0, -reportTimeBeforeDays)
		if p.StartTime.Before(beforeUTCDays) {
			return definition.ERROR_CODE_ERROR_TIME_BEFORE_DAYS
		}
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetPerformanceReportResponse struct {
	LogTime               string  `json:"log_time"`                  // 紀錄時間 (唯一)
	AgentId               int     `json:"agentId"`                   // 代理id
	AgentName             string  `json:"agent_name"`                // 代理名稱
	BetUser               int     `json:"bet_user"`                  // 總投注人數
	BetCount              int     `json:"bet_count"`                 // 注單數
	JackpotUser           int     `json:"jackpot_user"`              // jackpot中獎總人數
	JackpotCount          int     `json:"jackpot_count"`             // jackpot中獎總注單數
	SumYa                 float64 `json:"sum_ya"`                    // 總投注
	SumValidYa            float64 `json:"sum_valid_ya"`              // 有效投注
	SumDe                 float64 `json:"sum_de"`                    // 總派獎
	SumBonus              float64 `json:"sum_bonus"`                 // 紅利
	SumTax                float64 `json:"sum_tax"`                   // 抽水
	SumJpInjectWaterScore float64 `json:"sum_jp_inject_water_score"` // jp注水分數
	SumJpPrizeScore       float64 `json:"sum_jp_prize_score"`        // jp中獎分數
}

func NewGetPerformanceReportResponse() *GetPerformanceReportResponse {
	return &GetPerformanceReportResponse{}
}

func NewGetPerformanceReportResponseSlice() []*GetPerformanceReportResponse {
	return make([]*GetPerformanceReportResponse, 0)
}
