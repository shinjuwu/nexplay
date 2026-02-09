package model

type OperateCoinRequest struct {
	OrderId   string  `json:"orderid"`    //訂單號(流水號(格式:代理編號+yyyyMMddHHmmssSSS+account,長度不能超過 100 字符串))
	AgentId   int     `json:"agent_id"`   //代理編號
	AgentCode string  `json:"agent_code"` //代理識別碼
	Username  string  `json:"username"`   //會員帳號
	Action    string  `json:"action"`     //操作類型(加款:add, 扣款:sub)
	Coin      float64 `json:"coin"`       //金額
	OpCode    int     `json:"op_code"`    //操作代碼
	Reason    string  `json:"reason"`     //原因備註
}

func (p *OperateCoinRequest) CheckParams() (errMsg string, isSuccess bool) {
	isSuccess = false
	if p.Coin <= 0 {
		errMsg = "no changed, because coin is zero"
		return
	}

	if p.OrderId == "" || p.AgentCode == "" || p.Username == "" || p.Action == "" || p.OpCode <= 0 {
		errMsg = "required information is empty"
		return
	}

	isSuccess = true

	return
}

type OperateCoinResponse struct {
	OrderId    string  `jsonL:"orderid"`    //訂單號(流水號(格式:代理編號+yyyyMMddHHmmssSSS+ account,長度不能超過 100 字符串))
	AgentCode  string  `json:"agent_code"`  //代理識別碼
	Username   string  `json:"username"`    //會員帳號
	Action     string  `json:"action"`      //操作類型(加款:add, 扣款:sub)
	Coin       float64 `json:"coin"`        //帳變金額
	BeforeCoin float64 `json:"before_coin"` //帳變前金額
	AfterCoin  float64 `json:"after_coin"`  //帳變後金額
	OpCode     int     `json:"op_code"`     //操作代碼
	Reason     string  `json:"reason"`      //原因備註
}
