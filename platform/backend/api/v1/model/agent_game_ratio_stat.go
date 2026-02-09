package model

import (
	"backend/pkg/utils"
	"definition"
	"time"
)

type GetAgentGameRatioStatListRequest struct {
	AgentId  int `json:"agent_id"`  // 代理編號
	GameId   int `json:"game_id"`   // 遊戲編號
	RoomType int `json:"room_type"` // 房間類型編號
	DateTimeTableRequest
	ServerSideTableRequest
}

func (p *GetAgentGameRatioStatListRequest) CheckParams(reportTimeRange, reportTimeBeforeDays int) int {
	if p.AgentId < definition.AGENT_ID_ALL ||
		p.GameId < definition.GAME_ID_ALL ||
		p.RoomType < definition.ROOM_TYPE_ALL {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	if code := p.CheckDateTimeTableRequest(); code != definition.ERROR_CODE_SUCCESS {
		return code
	}

	if code := p.CheckServerSideTableRequest(); code != definition.ERROR_CODE_SUCCESS {
		return code
	}

	if p.EndTime.Sub(p.StartTime) > time.Duration(reportTimeRange)*24*time.Hour {
		return definition.ERROR_CODE_ERROR_TIME_RANGE
	}

	today := utils.GetTimeNowUTCTodayTime()
	beforeDays := time.Duration(reportTimeBeforeDays) * 24 * time.Hour
	if today.Sub(p.StartTime) > beforeDays || today.Sub(p.EndTime) > beforeDays {
		return definition.ERROR_CODE_ERROR_TIME_BEFORE_DAYS
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetAgentGameRatioStatListResponse struct {
	DataTablesResponse
	TotalPlatformWinloseScore float64 `json:"total_platform_winlose_score"` // 平台總輸贏
}

type GetAgentGameRatioStatResponse struct {
	AgentId       int     `json:"agent_id"`       // 代理編號
	AgentName     string  `json:"agent_name"`     // 代理名稱
	GameId        int     `json:"game_id"`        // 遊戲編號
	RoomType      int     `json:"room_type"`      // 房間類型編號
	PlayCount     int     `json:"play_count"`     // 注單數
	DeScore       float64 `json:"de_score"`       // 總玩家得分
	YaScore       float64 `json:"ya_score"`       // 總投注
	Tax           float64 `json:"tax"`            // 總遊戲抽水
	PlayerWinlose float64 `json:"player_winlose"` // 總玩家輸贏分數
	TaxedRtp      float64 `json:"taxed_rtp"`      // 抽水後rtp
	Bonus         float64 `json:"bonus"`          // 紅利
}
