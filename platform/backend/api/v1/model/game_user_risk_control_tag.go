package model

type GetGameUserRiskControlTagRequest struct {
	Id int `json:"id"` // 用戶id
}

type GetGameUserRiskControlTagResponse struct {
	Id             int    `json:"id"`               // 用戶id
	Username       string `json:"username"`         // 用戶名(第三方平台原始帳號)
	AgentId        int    `json:"agent_id"`         // 代理商id
	AgentName      string `json:"agent_name"`       // 代理商名稱
	RiskControlTag string `json:"risk_control_tag"` // 風控處置標示(4bit) "1000":禁止登入 "0100":禁止下注 "0010":禁止上分 "0001":禁止下分
}

type SetGameUserRiskControlTagRequest struct {
	Id             int    `json:"id"`               // 用戶id
	RiskControlTag string `json:"risk_control_tag"` // 風控處置標示(4bit) "1000":禁止登入 "0100":禁止下注 "0010":禁止上分 "0001":禁止下分
}
