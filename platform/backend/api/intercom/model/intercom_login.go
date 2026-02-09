package model

import "backend/pkg/utils"

type LoginGameResponse struct {
	Id               int     `json:"id" mapstructure:"id"`                                     // user id
	AgentId          int     `json:"agent_id" mapstructure:"agent_id"`                         // 代理id
	AgentCode        string  `json:"agent_code" mapstructure:"agent_code"`                     // 代理層級碼
	LevelCode        string  `json:"level_code" mapstructure:"level_code"`                     // 代理層級碼
	Username         string  `json:"username" mapstructure:"username"`                         // 原始帳號
	TransUsername    string  `json:"trans_username" mapstructure:"trans_username"`             // 轉換後帳號
	OrderId          string  `json:"order_id" mapstructure:"order_id"`                         // 加款訂單號
	Coin             float64 `json:"coin" mapstructure:"coin"`                                 // 遊戲幣
	KindId           int     `json:"kind_id" mapstructure:"kind_id"`                           // 遊戲id
	GameLink         string  `json:"game_link" mapstructure:"game_link"`                       // 遊戲 h5 link
	Token            string  `json:"token" mapstructure:"token"`                               // 驗證token
	IsNew            bool    `json:"is_new" mapstructure:"is_new"`                             // 是否創建新帳號
	LoginTime        string  `json:"login_time" mapstructure:"login_time"`                     // 登入時間
	LastLoginTime    string  `json:"last_login_time" mapstructure:"last_login_time"`           // 上次登入時間
	KillDiveState    int     `json:"kill_dive_state" mapstructure:"kill_dive_state"`           // 殺放設定狀態(一般玩家:0、定點玩家:1、黑名單玩家:2)
	UserMetadata     string  `json:"user_metadata" mapstructure:"user_metadata"`               // 遊戲用戶個人資訊
	IsRelogin        bool    `json:"is_relogin" mapstructure:"is_relogin"`                     // 是否斷線重連
	GameIconList     string  `json:"game_icon_list" mapstructure:"game_icon_list"`             // 遊戲icon list
	ServerInfoCode   string  `json:"server_info_code" mapstructure:"server_info_code"`         // 遊戲service id(紀錄玩家在哪個服務)
	IsNotKillDiveCal bool    `json:"is_not_kill_dive_cal" mapstructure:"is_not_kill_dive_cal"` // 是否計算殺放(內部測試代理，不計算殺放)(計算殺放[預設]:false、不計算:true)
	WalletType       int     `json:"wallet_type" mapstructure:"wallet_type"`                   // 錢包模式(轉帳錢包[預設]:0、單一錢包:1)
	LobbySwitchInfo  string  `json:"lobby_switch_info" mapstructure:"lobby_switch_info"`       // 大廳開關資訊
}

func (p *LoginGameResponse) LoginGameResponseOutput() map[string]interface{} {

	temp := make(map[string]interface{}, 0)

	temp["id"] = p.Id
	temp["trans_username"] = p.TransUsername
	temp["coin"] = p.Coin

	return temp

}

type UserMetadata struct {
	IsBeKillDive bool `json:"is_be_killdive"` // 是否曾經被追殺
}

func NewUserMetadataEmpty() *UserMetadata {
	return &UserMetadata{
		IsBeKillDive: false,
	}
}

func (p *UserMetadata) ToJson() string {
	return utils.ToJSON(p)
}
