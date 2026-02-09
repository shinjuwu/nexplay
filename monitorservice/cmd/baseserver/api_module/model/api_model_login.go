package model

import "time"

type CaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`     // 驗證碼id
	PicPath       string `json:"picPath"`       // base64 圖片字串
	CaptchaLength int    `json:"captchaLength"` // 驗證碼長度
	CaptchaValue  string `json:"captchaValue"`  // 驗證碼(測試模式才顯示)
	ExpiredTime   string `json:"expiredTime"`   // 驗證碼過期時間
}

//*****************************************************************************

type UsersRegisterRequest struct {
	Username    string   `json:"username"`    // 登入帳號
	Password    string   `json:"password"`    // 密碼(加密後的字串,不可用明碼)
	Nickname    string   `json:"nickname"`    // 自訂名稱
	Permissions []string `json:"permissions"` // 自訂權限，格式["dev", "qa"]
}

type UsersRegisterResponse struct {
	Username     string `json:"username"`
	Nickname     string `json:"nickname"`
	UserMetadata string `json:"user_metadata"`
}

//*****************************************************************************

type UsersLoginRequest struct {
	Username  string `json:"username"`  // 登入帳號
	Password  string `json:"password"`  // 密碼(加密後的字串,不可用明碼)
	Captcha   string `json:"captcha"`   // 驗證碼
	CaptchaId string `json:"captchaId"` // 驗證碼ID
}

type UserLoginData struct {
	TopCode       string    `json:"top_code"`
	Username      string    `json:"username"`
	Nickname      string    `json:"nickname"`
	UserMetadata  string    `json:"user_metadata"`
	CreateTime    time.Time `json:"create_time"`
	LastLoginTime time.Time `json:"last_login_time"`
	IsAdmin       bool      `json:"is_admin"`
	Permissions   []string  `json:"permissions"`
}

type UsersLoginResponse struct {
	UserData  UserLoginData
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}
