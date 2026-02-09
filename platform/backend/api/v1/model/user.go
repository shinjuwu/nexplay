package model

import (
	"backend/pkg/encrypt/base64url"
	"backend/pkg/jwt"
	"backend/pkg/utils"
	"backend/server/global"
	"definition"
	"errors"
	"time"

	"github.com/goccy/go-json"
)

const (
	READONLY_OFF     = 0
	READONLY_ON      = 1
	DEFAULT_NICKNAME = "New user"
)

type AdminUserResponse struct {
	AgentId   int    `json:"agent_id"`
	LevelCode string `json:"level_code"`
	Username  string `json:"username"`
	// Password string `json:"password"`
	Nickname string `json:"nickname"`
	// TODO: add 個人化設定
	// ex: 頭像,佈景主題,功能權限
	Permission  []int  `json:"permission"`
	AccountType int    `json:"account_type"` // 帳號類別(開發者:1, 總代理:2, 子代理:3)
	IsAdded     bool   `json:"is_added"`     // 是否為分身帳號(分身: true, 主帳號: false)
	Cooperation int    `json:"cooperation"`  // 合作模式(代理結帳類型, 1: 買分, 2: 信用)
	Currency    string `json:"currency"`
	// Jackpot設定
	JackpotStartTime time.Time `json:"jackpot_start_time"`
	JackpotEndTime   time.Time `json:"jackpot_end_time"`
	WalletType       int       `json:"wallet_type"` // 錢包類型(0:轉帳錢包,1:單一錢包)
}

type AliveToken struct {
	Username  string `json:"username"` // 帳戶
	Nickname  string `json:"nickname"` // 暱稱
	Token     string `json:"token"`    // 當下token
	ExpiresAt int    `json:"exp"`      // 過期時間
	IssuedAt  int    `json:"iat"`      // 發行時間
}

func (p *AliveToken) Set(un, nn, tk string, exp, iat int) {
	p.Username = un
	p.Nickname = nn
	p.Token = tk
	p.ExpiresAt = exp
	p.IssuedAt = iat
}

type AliveTokenListResponse struct {
	TokenList []AliveToken `json:"tokenList"`
}

// for api response return
func (p *AliveTokenListResponse) Response(_data *map[string]interface{}) error {

	for k, v := range *_data {
		cTmp, isConvert := v.(*jwt.CustomClaims)
		if !isConvert {
			return errors.New("AliveTokenListResponse: convert list error")
		}

		var tmp AliveToken
		u, _ := base64url.Decode(cTmp.Username)
		n, _ := base64url.Decode(cTmp.Nickname)
		tmp.Set(string(u), string(n), k, int(cTmp.ExpiresAt.Unix()), int(cTmp.IssuedAt.Unix()))

		p.TokenList = append(p.TokenList, tmp)
	}
	return nil
}

/*--------------------------------------------------------------------*/

// type GetAdminUsersRequest struct {
// }

// type GetAdminUsersResponse struct {
// 	GameUserList []*GetAdminUsersData `json:"game_adminlist"`
// }

type GetAdminUsersData struct {
	Username    string    `json:"username"`     // 帳號
	Nickname    string    `json:"nickname"`     // 暱稱
	TopUsername string    `json:"top_username"` // 父帳號
	RoleName    string    `json:"role_name"`    // 角色權限群組名稱
	IsEnabled   bool      `json:"id_enabled"`   // 狀態 (true: 開啟,false: 關閉)
	CreateTime  time.Time `json:"create_time"`  // 創建時間
	LoginTime   time.Time `json:"login_time"`   // 最後登入時間
}

/*--------------------------------------------------------------------*/

// 創建帳號必要資料
type AdminUserCreateRequest struct {
	Username string `json:"username"` // 用戶名
	Password string `json:"password"` // 密碼(加密後的字串,不可用明碼)
	Nickname string `json:"nickname"` // 用戶名
	Role     string `json:"role"`     // 角色權限群組id
	Info     string `json:"info"`     // 備註
}

// correct: true, failed: false
func (p *AdminUserCreateRequest) CheckParams() bool {
	if !utils.LowercaseEnglishAndNumber4To16.MatchString(p.Username) ||
		!utils.EnglishAndNumber8To16.MatchString(p.Password) ||
		p.Nickname == "" || utils.WordLength(p.Nickname) > 16 ||
		p.Role == "" ||
		utils.WordLength(p.Info) > 100 {
		return false
	}
	return true
}

/*--------------------------------------------------------------------*/

type GetAdminUserInfoRequest struct {
	Username string `json:"username"` // 用戶名
}

func (p *GetAdminUserInfoRequest) CheckParams() bool {
	return p.Username != ""
}

type GetAdminUserInfoResponse struct {
	IsEnabled int    `json:"is_enabled"` // 帳號狀態(1: 開啟, 0: 關閉)
	Role      string `json:"role"`       // 角色
	Info      string `json:"info"`       // 備註
}

/*--------------------------------------------------------------------*/

type UpdateAdminUserInfoRequest struct {
	Username     string `json:"username"`   // 用戶名
	PermissionId string `json:"role"`       // 角色權限id
	IsEnabled    int    `json:"is_enabled"` // 帳號狀態(1: 開啟, 0: 關閉)
	Info         string `json:"info"`       // 備註
}

func (p *UpdateAdminUserInfoRequest) CheckParams() bool {
	if !utils.LowercaseEnglishAndNumber4To16.MatchString(p.Username) ||
		p.PermissionId == "" ||
		utils.WordLength(p.Info) > 100 {
		return false
	}
	return true
}

/*--------------------------------------------------------------------*/

// type GetGameUsersRequest struct {
// }

// type GetGameUsersResponse struct {
// 	GameUserList []*GetGameUsersData `json:"game_userlist"`
// }

type GetGameUsersData struct {
	Id                 int                           `json:"id"`                    // 用戶id
	Username           string                        `json:"username"`              // 用戶名(第三方平台原始帳號)
	VipLevel           int                           `json:"vip_level"`             // vip 等級
	IsEnabled          bool                          `json:"id_enabled"`            // 狀態 (true: 開啟,false: 關閉)
	AgentId            int                           `json:"agent_id"`              // 代理商id
	AgentName          string                        `json:"agent_name"`            // 代理商名稱
	CoinIn             float64                       `json:"coin_in"`               // 轉入遊戲幣總和
	CoinOut            float64                       `json:"coin_out"`              // 轉出遊戲幣總和
	CreateTime         string                        `json:"create_time"`           // 帳號創建時間
	LastLoginTime      string                        `json:"last_login_time"`       // 最後登入時間
	IsOnline           bool                          `json:"is_online"`             // 是否在線 (true: 在線,false: 離線)
	HighRisk           bool                          `json:"high_risk"`             // 是否為高風險
	KillDiveState      int                           `json:"kill_dive_state"`       // 殺放設定狀態(一般玩家:0、定點玩家:1、黑名單玩家:2)
	KillDiveValue      float64                       `json:"kill_dive_value"`       // 定點額度(只有在殺放狀態設定為定點時有效)
	TagList            string                        `json:"tag_list"`              // 玩家自訂義標示(8bit)
	RiskControlTagList string                        `json:"risk_control_tag_list"` // 風控處置標示(4bit)
	CustomTagInfo      map[int]*global.CustomTagInfo `json:"custom_tag_info"`       // 自定義標籤資訊(array in json)
	WalletType         int                           `json:"wallet_type"`           // 錢包模式
}

/*--------------------------------------------------------------------*/

type GetGameUserInfoRequest struct {
	Id       int `json:"id"`        // 用戶id
	TimeZone int `json:"time_zone"` // 時間區域(UTC+X, X時間, 單位分鐘, ex: UTC+8 要傳入 -60*8=-480)
}

func (p *GetGameUserInfoRequest) CheckParams() bool {
	return p.Id > 0
}

type GetGameUserInfoResponse struct {
	ValidBetSum float64 `json:"valid_bet_sum"` // 總有效投注
	ValidBet    float64 `json:"valid_bet"`     // 有效投注
	ProfitSum   float64 `json:"profit_sum"`    // 總盈利
	Profit      float64 `json:"profit"`        // 盈利
	TaxSum      float64 `json:"tax_sum"`       // 總抽水
	BonusSum    float64 `json:"bonus_sum"`     // 總紅利
	Tax         float64 `json:"tax"`           // 抽水
	Bonus       float64 `json:"bonus"`         // 紅利
	Info        string  `json:"info"`          // 備註
	IsEnabled   bool    `json:"is_enabled"`    // 帳號狀態(true: 開啟)
}

func NewEmptyGetGameUserInfoResponse() *GetGameUserInfoResponse {
	return &GetGameUserInfoResponse{}
}

/*--------------------------------------------------------------------*/

type UpdateGameUserInfoRequest struct {
	Id        int    `json:"id"`         // 用戶id
	IsEnabled bool   `json:"is_enabled"` // 帳號狀態(true: 開啟)
	Info      string `json:"info"`       // 備註
}

func (p *UpdateGameUserInfoRequest) CheckParams() bool {
	return p.Id > 0 && utils.WordLength(p.Info) <= 100
}

/*--------------------------------------------------------------------*/

type SetPersonalInfoRequest struct {
	Nickname string `json:"nickname"` // 暱稱
}

func (p *SetPersonalInfoRequest) CheckParams() int {
	if p.Nickname == "" || utils.WordLength(p.Nickname) > 16 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}

/*--------------------------------------------------------------------*/

type SetPersonalPasswordRequest struct {
	OldPassword     string `json:"old_password"`     // 舊密碼(8~16)
	NewPassword     string `json:"new_password"`     // 新密碼(8~16)
	ConfirmPassword string `json:"confirm_password"` // 確認新密碼(8~16)
}

func (p *SetPersonalPasswordRequest) CheckParams() int {
	if !utils.EnglishAndNumber8To16.MatchString(p.OldPassword) ||
		!utils.EnglishAndNumber8To16.MatchString(p.NewPassword) ||
		!utils.EnglishAndNumber8To16.MatchString(p.ConfirmPassword) {
		return definition.ERROR_CODE_ERROR_PASSWORD_FORMAT
	} else if p.NewPassword != p.ConfirmPassword {
		return definition.ERROR_CODE_ERROR_CONFIRM_PASSWORD
	} else if p.OldPassword == p.NewPassword {
		return definition.ERROR_CODE_ERROR_PASSWORD_SAME
	}
	return definition.ERROR_CODE_SUCCESS
}

/*--------------------------------------------------------------------*/

type ResetPasswordRequest struct {
	Username string `json:"username"` // 用戶名
}

func (p *ResetPasswordRequest) CheckParams() int {
	if p.Username == "" {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	return definition.ERROR_CODE_SUCCESS
}

type ResetPasswordResponse struct {
	Username    string `json:"username"`     // 用戶名
	NewPassword string `json:"new_password"` // 新密碼
}

/*--------------------------------------------------------------------*/

type GetGameUserBalanceRequest struct {
	UserId int `json:"user_id"` // 用戶id
}

type GetGameUserBalanceResponse struct {
	UserId        int     `json:"user_id"`        // 用戶id
	WalletBalance float64 `json:"wallet_balance"` // 用戶錢包餘額
}

/*--------------------------------------------------------------------*/

type GetGameUserPlayCountDataRequest struct {
	UserId int `json:"user_id"` // 用戶id
}

type GameUserPlayCountData struct {
	GameCode    string `json:"game_code"`    // 遊戲代碼 (ex :百家樂=baccarat)
	RoomId      int    `json:"room_id"`      // 房間ID (ex : 1001X)
	PlayCount   int    `json:"play_count"`   // 單一遊戲局數加總
	NewbieLimit int    `json:"newbie_limit"` // 新手限制(單一遊戲限制)
}

type GetGameUserPlayCountDataResponse struct {
	UserId           int                     `json:"user_id"`            // 用戶id
	TotalNewbieLimit int                     `json:"total_newbie_limit"` // 新手限制(所有遊戲加總限制)
	Data             []GameUserPlayCountData `json:"data"`               // 用戶遊戲資料歷史紀錄
}

func (p *GetGameUserPlayCountDataResponse) DataConvert(jsonStr string) bool {
	if err := json.Unmarshal([]byte(jsonStr), &p.Data); err != nil {
		return false
	}

	return true
}
