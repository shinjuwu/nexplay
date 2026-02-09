package model

import (
	"time"
)

// same as admin_user of table
type AdminUser struct {
	AgentId      int       `json:"agent_id"`
	Username     string    `json:"username"`
	Nickname     string    `json:"nickname"`
	Password     string    `json:"password"`
	GoogleAuth   bool      `json:"google_auth"`
	GoogleKey    string    `json:"google_key"`
	AllowIp      string    `json:"allow_ip"`
	AccountType  int       `json:"account_type"`
	IsReadonly   int       `json:"is_readonly"`
	IsEnabled    int       `json:"is_enabled"`
	UpdateTime   time.Time `json:"update_time"`
	CreateTime   time.Time `json:"create_time"`
	IsAdded      bool      `json:"is_added"` // 是否為分身帳號
	LoginTime    time.Time `json:"login_time"`
	PermissionId string    `json:"permission_id"` // agent_permission id
	Info         string    `json:"info"`
}

func NewEmptyAdminUser() *AdminUser {
	return &AdminUser{}
}
