package model

import (
	"net/url"
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

type WalletConnInfo struct {
	// 	{
	// 		"path": "/dev-tools/channelHandle",
	// 		"domain": "172.30.0.152",
	// 		"scheme": "http",
	// 		"api_key": ""
	// 	  }
	Path   string `json:"path"`
	Domain string `json:"domain"`
	Scheme string `json:"scheme"`
	ApiKey string `json:"api_key"`
}

func (p *WalletConnInfo) ParseUrl(urlString string) (bool, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return false, err
	}
	p.Scheme = u.Scheme
	p.Domain = u.Host
	p.Path = u.Path

	if p.Domain == "" || p.Path == "" || p.Scheme == "" {
		return false, nil
	}

	return true, nil
}

func (p *WalletConnInfo) GetUrlPath() string {
	u := url.URL{
		Scheme: p.Scheme,
		Host:   p.Domain,
		Path:   p.Path,
	}
	return u.String()
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
	Creator             string                 `json:"creator"`                                                  // 開線人(那個後台帳號創了這個代理)
	IPWhitelist         []*AgentIPWhitelistObj `json:"ip_whitelist"`                                             // ip白名單列表(array obj)
	IPWhitelistBytes    []byte                 `json:"-"`                                                        // DB parsing用
	ApiIPWhitelist      []*AgentIPWhitelistObj `json:"api_ip_whitelist"`                                         // api ip白名單列表(array obj)
	ApiIPWhitelistBytes []byte                 `json:"-"`                                                        // DB parsing用
	KillSwitch          bool                   `json:"kill_switch"`                                              // 殺放開關
	KillRatio           float64                `json:"kill_ratio"`                                               // 殺放機率
	Currency            string                 `json:"currency"`                                                 // 幣種
	IsNotKillDiveCal    bool                   `json:"is_not_kill_dive_cal" mapstructure:"is_not_kill_dive_cal"` // 是否計算殺放(內部測試代理，不計算殺放)(計算殺放[預設]:false、不計算:true)
	WalletType          int                    `json:"wallet_type"`                                              // 錢包模式(轉帳錢包[預設]:0、單一錢包:1)
	WalletConnInfoBytes []byte                 `json:"-"`                                                        // 錢包連線方式(array obj)
	WalletConnInfo      *WalletConnInfo        `json:"wallet_conninfo"`                                          // 錢包連線方式(json)
	ChildAgentCount     int                    `json:"child_agent_count"`                                        // 子代理數量
	JackpotStatus       int                    `json:"jackpot_status"`                                           // jackpot 狀態
	JackpotStartTime    time.Time              `json:"jackpot_start_time"`                                       // jackpot 開始時間
	JackpotEndTime      time.Time              `json:"jackpot_end_time"`                                         // jackpot 結束時間
	LobbySwitchInfo     int                    `json:"lobby_switch_info"`                                        // 大廳開關(bitwise儲存)
}

func NewEmptyAgent() *Agent {
	return &Agent{
		IPWhitelist:    NewEmptyAgentIpWhitelist(),
		ApiIPWhitelist: NewEmptyAgentIpWhitelist(),
		WalletConnInfo: new(WalletConnInfo),
	}
}
