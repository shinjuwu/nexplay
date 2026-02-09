package model

import (
	"backend/pkg/utils"
	"backend/server/table/model"
)

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用戶名
	Password  string `json:"password"`  // 密碼(加密後的字串,不可用明碼)
	Captcha   string `json:"captcha"`   // 驗證碼
	CaptchaId string `json:"captchaId"` // 驗證碼ID
}

// correct: true, failed: false
func (p *Login) CheckParams() bool {
	if p.Password == "" || p.Username == "" || p.CaptchaId == "" || p.Captcha == "" || utils.IsChinese(p.Password) || utils.IsChinese(p.Username) {
		return false
	}
	return true
}

type LoginResponse struct {
	UserData  AdminUserResponse `json:"userData"` // 用戶個人資料
	Token     string            `json:"token"`
	ExpiresAt int64             `json:"expiresAt"`
}

type LoginAdminUser struct {
	model.AdminUser
	Permission      model.PermissionSlice `json:"permission"` // 權限內容
	PermissionBytes []byte                `json:"-"`          // 權限內容(DB parsing用)
}
