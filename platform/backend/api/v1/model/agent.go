package model

import (
	"backend/pkg/utils"
	table_model "backend/server/table/model"
	"definition"
	"time"
)

const TopLevelCodeLength = 4

type CreateAgentRequest struct {
	Account         string                      `json:"account"`           // 帳號
	Password        string                      `json:"password"`          // 密碼
	Nickname        string                      `json:"nickname"`          // 代理商名稱(後台顯示名稱)
	Commission      int                         `json:"commission"`        // 分成(萬分之X)
	IpWhitelist     string                      `json:"ip_whitelist"`      // IP 白名單
	Cooperation     int                         `json:"cooperation"`       // 合作模式(代理結帳類型, 1: 買分, 2: 信用)
	Info            string                      `json:"info"`              // 備註
	Role            string                      `json:"role"`              // 角色權限群組(uuid)
	Currency        string                      `json:"currency"`          // 幣種
	WalletType      int                         `json:"wallet_type"`       // 錢包模式(轉帳錢包[預設]:0、單一錢包:1)
	WalletUrl       string                      `json:"wallet_url"`        // wallet url for third api
	WalletConnInfo  *table_model.WalletConnInfo `json:"wallet_conninfo"`   // 錢包連線方式(json)
	LobbySwitchInfo int                         `json:"lobby_switch_info"` // 大廳開關設定
}

func NewCreateAgentRequest() *CreateAgentRequest {
	return &CreateAgentRequest{
		WalletConnInfo: new(table_model.WalletConnInfo),
	}
}

// correct: true, failed: false
func (p *CreateAgentRequest) CheckParams() bool {
	// TODO: 檢查IP格式
	if !utils.LowercaseEnglishAndNumber4To16.MatchString(p.Account) ||
		!utils.EnglishAndNumber8To16.MatchString(p.Password) ||
		!utils.EnglishAndNumber4To16.MatchString(p.Nickname) ||
		p.Cooperation <= 0 || p.Cooperation > 2 ||
		p.Commission < 0 || p.Commission > 100 ||
		p.IpWhitelist == "" || p.Currency == "" ||
		utils.WordLength(p.Info) > 100 ||
		p.LobbySwitchInfo < 0 {
		return false
	}

	if p.WalletType == definition.AGENT_WALLET_SINGLE && p.WalletUrl == "" {
		return false
	}

	return true
}

/**************************************************/

type GetAgentListRequest struct {
	Id int `json:"id" mapstructure:"id"` // 代理id
}

func NewGetAgentListRequest() *GetAgentListRequest {
	return &GetAgentListRequest{}
}

func (p *GetAgentListRequest) CheckParams() bool {
	return !(p.Id < 0)
}

type GetAgentListResponse struct {
	Id            int       `json:"id" mapstructure:"id"`                           // 代理id
	OwnerName     string    `json:"owner_name" mapstructure:"owner_name"`           // 開線人帳號(最上層id)
	AdminUserName string    `json:"admin_user_name" mapstructure:"admin_user_name"` // 代理後台帳號
	Name          string    `json:"name" mapstructure:"name"`                       // 代理商名稱
	TopAgentName  string    `json:"top_agent_name" mapstructure:"top_agent_name"`   // 上層代理名稱
	Code          string    `json:"code" mapstructure:"code"`                       // 識別碼
	Commission    int       `json:"commission" mapstructure:"commission"`           // 占成(萬分之n)
	Info          string    `json:"info" mapstructure:"info"`                       // 備註
	IsEnabled     int       `json:"is_enabled" mapstructure:"is_enabled"`           // 是否開放
	UpdateTime    time.Time `json:"update_time" mapstructure:"update_time"`         // 最後更新時間
	CreateTime    time.Time `json:"create_time" mapstructure:"create_time"`         // 創建時間
	MemberCount   int       `json:"member_count" mapstructure:"member_count"`       // 會員數量
	TopAgentId    int       `json:"top_agent_id" mapstructure:"top_agent_id"`       // 上層代理ID
	Cooperation   int       `json:"cooperation" mapstructure:"cooperation"`         // 合作模式
	CoinLimit     float64   `json:"coin_limit" mapstructure:"coin_limit"`           // 可使用額度上限
	CoinUse       float64   `json:"coin_use" mapstructure:"coin_use"`               // 已使用額度
	Amount        float64   `json:"amount" mapstructure:"amount"`                   // 餘額
	Currency      string    `json:"currency" mapstructure:"currency"`               // 幣種
	LevelCode     string    `json:"level_code" mapstructure:"level_code"`           // 會員層級碼
	RoleName      string    `json:"role_name" mapstructure:"role_name"`             // 角色名稱
}

func NewGetAgentListResponse() *GetAgentListResponse {
	return &GetAgentListResponse{}
}

func NewGetAgentListResponseSlice() []*GetAgentListResponse {
	return make([]*GetAgentListResponse, 0)
}

func (p *GetAgentListResponse) TransVal(obj *table_model.Agent) {
	// p.AdminUserName = obj.Name
	// p.TopAgentId = obj.TopAgentId
	p.Id = obj.Id
	p.OwnerName = obj.Creator
	p.AdminUserName = obj.AdminUsername
	p.Name = obj.Name
	p.TopAgentId = obj.TopAgentId
	p.Code = obj.Code
	p.Commission = obj.Commission
	p.Info = obj.Info
	p.IsEnabled = obj.IsEnabled
	p.UpdateTime = obj.UpdateTime
	p.CreateTime = obj.CreateTime
	p.MemberCount = obj.MemberCount
	p.CoinLimit = obj.CoinLimit
	p.CoinUse = obj.CoinUse
	p.Currency = obj.Currency
	p.LevelCode = obj.LevelCode
	p.Cooperation = obj.Cooperation
}

/**************************************************/

type GetAgentSecretKeyRequest struct {
	Id int `json:"id" mapstructure:"id"` // 代理id
}

func NewGetAgentSecretKeyRequest() *GetAgentSecretKeyRequest {
	return &GetAgentSecretKeyRequest{}
}

func (p *GetAgentSecretKeyRequest) CheckParams() bool {
	return !(p.Id < 0)
}

type GetAgentSecretKeyResponse struct {
	AesKey string `json:"aeskey"`
	Md5Key string `json:"md5key"`
}

func NewGetAgentSecretKeyResponse() *GetAgentSecretKeyResponse {
	return &GetAgentSecretKeyResponse{}
}

/**************************************************/

type GetAgentCoinSupplyInfoRequest struct {
	Id int `json:"id" mapstructure:"id"` // 代理id
}

func NewGetAgentCoinSupplyInfoRequest() *GetAgentCoinSupplyInfoRequest {
	return &GetAgentCoinSupplyInfoRequest{}
}

func (p *GetAgentCoinSupplyInfoRequest) CheckParams() bool {
	return !(p.Id < 0)
}

type GetAgentCoinSupplyInfoResponse struct {
	Id              int     `json:"id" mapstructure:"id"`
	TopAgentId      int     `json:"top_agent_id" mapstructure:"top_agent_id"`
	Name            string  `json:"name" mapstructure:"name"`               // 代理名稱
	Commission      int     `json:"commission" mapstructure:"commission"`   // 分成(萬分之n)
	Cooperation     int     `json:"cooperation" mapstructure:"cooperation"` // 合作模式(代理結帳類型, 1: 買分, 2: 信用)
	CoinLimit       float64 `json:"coin_limit" mapstructure:"coin_limit"`   // 買分模式分數上限
	CoinUse         float64 `json:"coin_use" mapstructure:"coin_use"`       // 已使用分數
	Info            string  `json:"info" mapstructure:"info"`               // 自動補分設定
	Role            string  `json:"role" mapstructure:"role"`               // admin user role
	RoleName        string  `json:"role_name" mapstructure:"role_name"`     // admin user role 名稱
	WalletType      int     `json:"wallet_type"`                            // 錢包類型(0:轉帳錢包,1:單一錢包)
	WalletConnInfo  string  `json:"wallet_conninfo"`                        // 錢包API連線位置
	LobbySwitchInfo int     `json:"lobby_switch_info"`                      // 大廳開關設定
}

func NewGetAgentCoinSupplyInfoResponse() *GetAgentCoinSupplyInfoResponse {
	return &GetAgentCoinSupplyInfoResponse{}
}

/**************************************************/

type SetAgentCoinSupplyInfoRequest struct {
	Id              int    `json:"id" mapstructure:"id"`                 // 代理id
	Name            string `json:"name" mapstructure:"name"`             // 代理名稱
	Commission      int    `json:"commission" mapstructure:"commission"` // 分成(萬分之n)
	Info            string `json:"info" mapstructure:"info"`             //
	Role            string `json:"role"`                                 // 權限群組id(uuid)
	WalletConnInfo  string `json:"wallet_conninfo"`                      // 錢包API連線位置
	LobbySwitchInfo int    `json:"lobby_switch_info"`                    // 大廳開關設定
}

func NewSetAgentCoinSupplyInfoRequest() *SetAgentCoinSupplyInfoRequest {
	return &SetAgentCoinSupplyInfoRequest{}
}

func (p *SetAgentCoinSupplyInfoRequest) CheckParams() bool {
	if p.Id <= 0 ||
		p.Commission < 0 || p.Commission > 100 ||
		p.Role == "" ||
		!utils.EnglishAndNumber4To16.MatchString(p.Name) ||
		utils.WordLength(p.Info) > 100 ||
		p.LobbySwitchInfo < 0 {
		return false
	}
	return true
}
