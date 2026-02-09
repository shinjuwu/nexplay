package model

import "definition"

type GameSetting struct {
	GameId                  int     `json:"gameId"`                  // 遊戲id
	MatchGameRTP            float64 `json:"MatchGameRTP"`            // 控牌RTP
	MatchGameKillRate       float64 `json:"MatchGameKillRate"`       // 控牌殺率
	MatchGames              int     `json:"MatchGames"`              // 控牌場次
	NormalMatchGameRTP      float64 `json:"NormalMatchGameRTP"`      // 單間RTP
	NormalMatchGameKillRate float64 `json:"NormalMatchGameKillRate"` // 單間殺率
	LowBoundRTP             float64 `json:"LowBoundRTP"`             // 保底RTP
	LimitOdds               float64 `json:"LimitOdds"`               // 限制倍率
	LimitMoney              float64 `json:"LimitMoney"`              // 禁開分數
}

func (g *GameSetting) CheckParams() int {
	if g.MatchGameRTP < 0 || g.MatchGameRTP > 1 ||
		g.MatchGameKillRate < 0 || g.MatchGameKillRate > 1 ||
		g.MatchGames < 0 || g.MatchGames > 9999 ||
		g.NormalMatchGameRTP < 0 || g.NormalMatchGameRTP > 1 ||
		g.NormalMatchGameKillRate < 0 || g.NormalMatchGameKillRate > 1 ||
		g.LowBoundRTP < 0 || g.LowBoundRTP > 1 ||
		g.LimitOdds < 0 || g.LimitOdds > 9999 ||
		g.LimitMoney < 0 || g.LimitMoney > 999999 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}

func (g *GameSetting) IsEqual(obj *GameSetting) bool {
	if g == obj {
		return true
	}

	if g == nil || obj == nil {
		return false
	}

	if g.GameId != obj.GameId ||
		g.MatchGameRTP != obj.MatchGameRTP ||
		g.MatchGameKillRate != obj.MatchGameKillRate ||
		g.MatchGames != obj.MatchGames ||
		g.NormalMatchGameRTP != obj.NormalMatchGameRTP ||
		g.NormalMatchGameKillRate != obj.NormalMatchGameKillRate ||
		g.LowBoundRTP != obj.LowBoundRTP ||
		g.LimitOdds != obj.LimitOdds ||
		g.LimitMoney != obj.LimitMoney {
		return false
	}

	return true
}

type GetGameSettingResponse []*GameSetting

type SetGameSettingRequest []GameSetting

func (gs SetGameSettingRequest) CheckParams() int {
	for _, g := range gs {
		if code := g.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
			return code
		}
	}

	return definition.ERROR_CODE_SUCCESS
}
