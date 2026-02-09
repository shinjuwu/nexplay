package model

type AutoRiskControlSetting struct {
	GameUserCoinInAndOutRequestPerMinuteLimit int64   `json:"game_user_coin_in_and_out_request_per_minute_limit"` // 玩家每分鐘可請求上下分最多次數(超過該次數玩家會無法上下分，需要等待至下一分鐘才可繼續使用)
	GameUserApiRequestPerSecondLimit          int64   `json:"game_user_api_request_per_second_limit"`             // 玩家每秒可請求api最多次數(超過該次數玩家會被禁用，需要後台解鎖)
	GameUserCoinInAndOutDiffLimit             float64 `json:"game_user_coin_in_and_out_diff_limit"`               // 玩家上分與下分的最大差額(超過該數值玩家會禁止下分，需要後台解鎖)
	GameUserWinRateLimit                      float64 `json:"game_user_win_rate_limit"`                           // 玩家勝率限制(超過該比例玩家會禁止下注，需要後台解鎖)
	IsEnabled                                 bool    `json:"is_enabled"`                                         // 開關
}

// default
/*
{
  "game_user_coin_in_and_out_request_per_minute_limit": 5,
  "game_user_api_request_per_second_limit": 10,
  "game_user_coin_in_and_out_diff_limit":100000,
  "game_user_win_rate_limit": 0.7,
  "is_enabled": false,
}
*/
func NewAutoRiskControlSetting() *AutoRiskControlSetting {
	return &AutoRiskControlSetting{
		GameUserCoinInAndOutRequestPerMinuteLimit: 5,
		GameUserApiRequestPerSecondLimit:          10,
		GameUserCoinInAndOutDiffLimit:             100000,
		GameUserWinRateLimit:                      0.7,
		IsEnabled:                                 false,
	}
}
