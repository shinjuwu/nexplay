package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

// 用戶資訊
type BaseClaims struct {
	ID           uint   `json:"id"`
	LevelCode    string `json:"level_code"`
	Username     string `json:"useranme"`
	Nickname     string `json:"nickanme"`
	Password     string `json:"password"`
	AccountType  int    `json:"account_type"`
	Cooperation  int    `json:"cooperation"` // 合作模式(代理結帳類型, 1: 買分, 2: 信用)
	PermissionId string `json:"permission_id"`
	IsAdded      bool   `json:"is_added"`
	LoginTime    string `json:"login_time"`
	Ip           string `json:"ip"`
}

// 自定義憑證格式
type CustomClaims struct {
	BaseClaims
	BufferTime int64 `json:"bufferTime"` // 更新 token 緩衝時間
	jwt.RegisteredClaims
	// IsBlackList bool
}
