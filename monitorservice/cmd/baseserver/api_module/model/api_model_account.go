package model

import "time"

type UserInfoResponse struct {
	TopCode     string   `json:"top_code"`    // 層級碼
	Username    string   `json:"username"`    // 帳號
	Nickname    string   `json:"nickname"`    // 暱稱
	IsAdmin     bool     `json:"is_admin"`    // 是否是管理者帳號
	LoginTime   string   `json:"login_time"`  // 本次登入時間
	ExpiresAt   int64    `json:"expiresAt"`   // token 過期時間
	Permissions []string `json:"permissions"` // 自訂權限，格式["dev", "qa"]
}

//*****************************************************************************

type ModifyUsersInfoRequest struct {
	Username    string   `json:"username"`    // 用戶帳號
	Nickname    string   `json:"nickname"`    // 暱稱
	IsEnabled   bool     `json:"is_enabled"`  // 帳號開啟狀態
	Permissions []string `json:"permissions"` // 自訂權限，格式["dev", "qa"]
	Info        string   `json:"info"`        // 備註
}

//*****************************************************************************

type ModifyUsersPasswordRequest struct {
	Username string `json:"username"` // 用戶帳號
	Password string `json:"password"` // 密碼(加密後的字串,不可用明碼)
}

//*****************************************************************************

type BlockUsersRequest struct {
	ID        string `json:"id"`         // 用戶id
	Username  string `json:"username"`   // 帳號
	IsEnabled bool   `json:"is_enabled"` // 是否關閉帳號
}

//*****************************************************************************

type GetUserInfoListResponse struct {
	Data []*GetUserInfoResponse `json:"data"`
}

//*****************************************************************************

type GetUserInfoRequest struct {
	Username string `json:"username"`
}

type GetUserInfoResponse struct {
	// Id            string    `json:"id"`
	Username      string    `json:"username"`
	Nickname      string    `json:"nickname"`
	UserMetadata  string    `json:"user_metadata"`
	IsEnabled     bool      `json:"is_enabled"`
	LastLoginTime time.Time `json:"last_login_time"`
	Info          string    `json:"info"`
	CreateTime    time.Time `json:"create_time"`
	Permissions   []string  `json:"permissions"`
}

//*****************************************************************************

type ModifyInfoRequest struct {
	Nickname string `json:"nickname"` // 暱稱
}

//*****************************************************************************

type ModifyPasswordRequest struct {
	Password string `json:"password"` // 密碼(加密後的字串,不可用明碼)
}
