package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

// 用戶資訊
type BaseClaims struct {
	ID          string   `json:"id"`
	TopCode     string   `json:"top_code"`
	Username    string   `json:"useranme"`
	Nickname    string   `json:"nickanme"`
	Password    string   `json:"password"`
	LoginTime   string   `json:"login_time"`
	IsAdmin     bool     `json:"is_admin"`
	Permissions []string `json:"permissions"`
}

// 自定義憑證格式
type CustomClaims struct {
	BaseClaims
	BufferTime int64 `json:"bufferTime"` // 更新 token 緩衝時間
	jwt.RegisteredClaims
}
