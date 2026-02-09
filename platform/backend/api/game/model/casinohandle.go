package model

type CasinoHandleRequest struct {
	Agent     string  `json:"agent"`      // 代理編號（平台提供）
	AgentCode string  `json:"agent_code"` // 代理代碼
	LevelCode string  `json:"level_code"` // 代理層級代碼
	Account   string  `json:"account"`
	Password  string  `json:"password"`
	Coin      string  `json:"coin"`
	Currency  string  `json:"currency"` // 幣種
	ToCoin    float64 `json:"to_coin"`  // 轉換比值
	Lang      string  `json:"lang"`     // 遊戲前端語系(串接時傳入)
}

type CasinoHandleResponse struct {
	S int         `json:"s"` // 子操作類型
	M string      `json:"m"` // 主操作類型
	D interface{} `json:"d"` // 數據結果
}

func CreateCasinoHandleResponse(subOper int, data interface{}) *CasinoHandleResponse {
	var tmp CasinoHandleResponse
	tmp.Response(subOper, data)
	return &tmp
}

func (p *CasinoHandleResponse) Response(subOper int, data interface{}) {

	p.M = CasinoHandle_Path

	// 原命令+100
	p.S = subOper + ChannelHandle_ResponseBase
	// 帶入資料
	p.D = data
}
