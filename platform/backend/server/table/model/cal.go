package model

type RpRecordData struct {
	LogTime               string  `json:"log_time"`                  // 紀錄時間 (唯一)
	AgentId               int     `json:"agent_id"`                  // 代理id
	LevelCode             string  `json:"level_code"`                // 層級碼
	BetUser               int     `json:"bet_user"`                  // 總投注人數
	BetCount              int     `json:"bet_count"`                 // 注單數
	JackpotUser           int     `json:"jacpot_user"`               // 總Jackpot中獎人數
	JackpotCount          int     `json:"jacpot_count"`              // Jackpot中獎注單數
	SumYa                 float64 `json:"sum_ya"`                    // 總投注
	SumValidYa            float64 `json:"sum_valid_ya"`              // 有效投注
	SumDe                 float64 `json:"sum_de"`                    // 總派獎
	SumBonus              float64 `json:"sum_bonus"`                 // 紅利
	SumTax                float64 `json:"sum_tax"`                   // 抽水
	SumJpInjectWaterScore float64 `json:"sum_jp_inject_water_score"` // jp注水分數
	SumJpPrizeScore       float64 `json:"sum_jp_prize_score"`        // jp中獎分數
}

type RtRealtimeGameHourData struct {
	LogTime      string  `json:"log_time"`       // 紀錄時間 (PK)
	AgentId      int     `json:"agent_id"`       // 代理編號 (PK)
	GameId       int     `json:"game_id"`        // 遊戲id(PK)
	BetUser      int     `json:"bet_user"`       // 不重複投注人數
	BetCount     int     `json:"bet_count"`      // 注單數
	YaScore      float64 `json:"ya_score"`       // 投注
	ValidYaScore float64 `json:"valid_ya_score"` // 有效投注
	DeScore      float64 `json:"de_score"`       // 得分
	Bonus        float64 `json:"bonus"`          // 紅利獎金
	Tax          float64 `json:"tax"`            // 抽水
}
