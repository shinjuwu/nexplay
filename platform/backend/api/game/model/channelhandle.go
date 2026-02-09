package model

import (
	"backend/pkg/encrypt/aescbc"
	"backend/pkg/encrypt/md5hash"
	"backend/pkg/utils"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

type ChannelHandleRequest struct {
	Agent     string              `json:"agent"`      // 代理編號（平台提供）
	AgentCode string              `json:"agent_code"` // 代理代碼
	LevelCode string              `json:"level_code"` // 代理層級代碼
	TimeStamp string              `json:"timestamp"`  // 時間戳(Unix 時間戳帶上毫秒),獲取當前時間（1488781836949）
	Param     string              `json:"param"`      // 參數加密字符串(aes 加密)
	ParamMap  map[string][]string `json:"param_map"`  // 參數解密後
	Currency  string              `json:"currency"`   // 幣種
	ToCoin    float64             `json:"to_coin"`    // 轉換比值
	Key       string              `json:"key"`        // Md5 校驗字符串 Encrypt.MD5(agent+timestamp+MD5Key)
	Lang      string              `json:"lang"`       // 遊戲前端語系(串接時傳入)
}

func (p *ChannelHandleRequest) Set(agent, agentCode, levelCode, timestamp, param, key, currency, lang string) {
	p.Agent = agent
	p.AgentCode = agentCode
	p.LevelCode = levelCode
	p.TimeStamp = timestamp
	p.Param = param
	p.Key = key
	p.Currency = currency
	p.Lang = lang
}

// assign & check
func (p *ChannelHandleRequest) Assign(agent, agentCode, levelCode, timestamp, param, key, aesKey, md5Key, currency, lang string) (isSuccess bool, returnMsg string) {

	isSuccess = false
	if len(agent) == 0 {
		returnMsg = "invalid agent length"
		return
	}
	agentId := utils.ToInt(agent, 0)
	if agentId <= 0 {
		returnMsg = "invalid agent id"
		return
	}

	if len(agentCode) > 8 {
		returnMsg = "invalid agentCode string length"
		return
	}

	if len(timestamp) != 13 { // 毫秒
		returnMsg = "invalid timestamp string length"
		return
	}

	if len(key) == 0 {
		returnMsg = "invalid key string length"
		return
	}

	if len(param) == 0 {
		returnMsg = "invalid string length"
		return
	}

	/*
		AES加密遇到的问题 aes加密后变成空格
		项目中使用aes加密传递数据的时候，发现数据根据aes密钥无法解密，于是开始寻找解决方案。在分析加密数据的时候发现数据有些地方会产生空格，于是怀疑可能是空格的问题。最后发现后台只要遇到“+”的字符串就会变为空格。
		解决方案：
		将空格替换成"+"号
		原因分析：get请求会过滤特殊字符导致数据差异，post请求就不会有这种情况。

		www.csdn.net/tags/MtTaEgwsNTM4NTc3LWJsb2cO0O0O
	*/
	param = strings.Replace(param, " ", "+", -1)

	p.Set(agent, agentCode, levelCode, timestamp, param, key, currency, lang)
	// 先用 base64 解密
	b64Decoding, _ := base64.StdEncoding.DecodeString(p.Param)
	// 再用 ase 解密
	decryptString, err := aescbc.AesDecrypt(b64Decoding, []byte(aesKey))
	if err != nil {
		returnMsg = err.Error()
		return
	}

	p.ParamMap, err = url.ParseQuery(string(decryptString))
	if err != nil {
		returnMsg = err.Error()
		return
	}

	if !p.CheckMD5Param(md5Key) {
		returnMsg = "invalid verify"
		return
	}

	isSuccess = true
	returnMsg = ""
	return
}

func (p *ChannelHandleRequest) CheckMD5Param(md5key string) bool {
	/*
	   Md5 校驗字符串 Encrypt.MD5(agent+timestamp+MD5Key)
	*/

	data := fmt.Sprintf("%s%s%s", p.Agent, p.TimeStamp, md5key)
	md5Data := md5hash.Hash32bit(data)
	return p.Key == md5Data
}

// convert url param string to url.Values
// type Values map[string][]string
func (p *ChannelHandleRequest) ParamToUrlMap(aesKey string) bool {

	dataBytes, err := aescbc.AesDecrypt([]byte(p.Param), []byte(aesKey))
	if err != nil {
		return false
	}
	queryData, err := url.ParseQuery(string(dataBytes))
	if err != nil {
		return false
	}

	p.ParamMap = queryData
	return true
}

type ChannelHandleResponse struct {
	S int         `json:"s"` // 子操作類型
	M string      `json:"m"` // 主操作類型
	D interface{} `json:"d"` // 數據結果
}

func NewEmptyChannelHandleResponse() *ChannelHandleResponse {
	return &ChannelHandleResponse{}
}

func CreateChannelHandleResponse(subOper int, data interface{}) *ChannelHandleResponse {
	var tmp ChannelHandleResponse
	tmp.Response(subOper, data)
	return &tmp
}

func (p *ChannelHandleResponse) Response(subOper int, data interface{}) {

	if subOper == GetRecordHandle_CheckGameRecord {
		p.M = GetRecordHandle_Path
	} else {
		p.M = ChannelHandle_Path
	}
	// 原命令+100
	p.S = subOper + ChannelHandle_ResponseBase
	// 帶入資料
	p.D = data
}

//-------------------------------------------------------------------

type LoginGameResponse struct {
	Id               int     `json:"id" mapstructure:"id"`                                     // user id
	AgentId          int     `json:"agent_id" mapstructure:"agent_id"`                         // 代理id
	AgentCode        string  `json:"agent_code" mapstructure:"agent_code"`                     // 代理層級碼
	LevelCode        string  `json:"level_code" mapstructure:"level_code"`                     // 代理層級碼
	Username         string  `json:"username" mapstructure:"username"`                         // 原始帳號
	TransUsername    string  `json:"trans_username" mapstructure:"trans_username"`             // 轉換後帳號
	OrderId          string  `json:"order_id" mapstructure:"order_id"`                         // 加款訂單號
	Coin             float64 `json:"coin" mapstructure:"coin"`                                 // 遊戲幣
	KindId           int64   `json:"kind_id" mapstructure:"kind_id"`                           // 遊戲id
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
