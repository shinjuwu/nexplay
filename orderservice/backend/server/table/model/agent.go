package model

import (
	"backend/pkg/utils"
	"time"
)

// swagger:model
type AgentIPWhitelistObj struct {
	CreateTime int64  `json:"create_time"` // 建立時間(timestamp)
	IPAddress  string `json:"ip_address"`  // ip位址
	Info       string `json:"info"`        // 備註(長度30)
	Creator    string `json:"creator"`     // 創建人
}

func NewEmptyAgentIpWhitelist() []*AgentIPWhitelistObj {
	return make([]*AgentIPWhitelistObj, 0)
}

type Agent struct {
	Id                  int                    `json:"id" mapstructure:"id"` //pkey
	AdminUsername       string                 `json:"admin_username" mapstructure:"admin_username"`
	Name                string                 `json:"name" mapstructure:"name"`
	Code                string                 `json:"code" mapstructure:"code"`
	SecretKey           string                 `json:"secret_key" mapstructure:"secret_key"`
	Md5Key              string                 `json:"md5_key" mapstructure:"md5_key"`
	AesKey              string                 `json:"aes_key" mapstructure:"aes_key"`
	Commission          int                    `json:"commission" mapstructure:"commission"`
	Info                string                 `json:"info" mapstructure:"info"`
	IsEnabled           int                    `json:"is_enabled" mapstructure:"is_enabled"`
	DisableTime         time.Time              `json:"disable_time" mapstructure:"disable_time"`
	UpdateTime          time.Time              `json:"update_time" mapstructure:"update_time"`
	CreateTime          time.Time              `json:"create_time" mapstructure:"create_time"`
	IsTopAgent          bool                   `json:"is_top_agent" mapstructure:"is_top_agent"`
	TopAgentId          int                    `json:"top_agent_id" mapstructure:"top_agent_id"`
	Cooperation         int                    `json:"cooperation" mapstructure:"cooperation"`
	CoinLimit           float64                `json:"coin_limit" mapstructure:"coin_limit"`
	CoinUse             float64                `json:"coin_use" mapstructure:"coin_use"`
	LevelCode           string                 `json:"level_code"`
	MemberCount         int                    `json:"member_count"`
	Creator             string                 `json:"creator"`          // 開線人(那個後台帳號創了這個代理)
	IPWhitelist         []*AgentIPWhitelistObj `json:"ip_whitelist"`     // ip白名單列表(array obj)
	IPWhitelistBytes    []byte                 `json:"-"`                // DB parsing用
	ApiIPWhitelist      []*AgentIPWhitelistObj `json:"api_ip_whitelist"` // api ip白名單列表(array obj)
	ApiIPWhitelistBytes []byte                 `json:"-"`                // DB parsing用
	KillSwitch          bool                   `json:"kill_switch"`      // 殺放開關
	KillRatio           float64                `json:"kill_ratio"`       // 殺放機率
	Currency            string                 `json:"currency"`         // 幣種
}

func NewEmptyAgent() *Agent {
	return &Agent{
		IPWhitelist:    NewEmptyAgentIpWhitelist(),
		ApiIPWhitelist: NewEmptyAgentIpWhitelist(),
	}
}

func (p *Agent) TranslateWhiteIPList() {
	for _, result := range utils.ToArrayMap(p.IPWhitelistBytes) {
		p.IPWhitelist = append(p.IPWhitelist, &AgentIPWhitelistObj{
			CreateTime: int64(result["create_time"].(float64)),
			IPAddress:  result["ip_address"].(string),
			Info:       result["info"].(string),
			Creator:    result["creator"].(string),
		})
	}

	for _, result := range utils.ToArrayMap(p.ApiIPWhitelistBytes) {
		p.ApiIPWhitelist = append(p.ApiIPWhitelist, &AgentIPWhitelistObj{
			CreateTime: int64(result["create_time"].(float64)),
			IPAddress:  result["ip_address"].(string),
			Info:       result["info"].(string),
			Creator:    result["creator"].(string),
		})
	}
}
