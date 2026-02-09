package model

import "time"

type Users struct {
	Id            string    `json:"id"`
	TopCode       string    `json:"top_code"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Nickname      string    `json:"nickname"`
	UserMetadata  string    `json:"user_metadata"`
	IsEnabled     bool      `json:"is_enabled"`
	LastLoginTime time.Time `json:"last_login_time"`
	Info          string    `json:"info"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
	DisableTime   time.Time `json:"disable_time"`
	IsAdmin       bool      `json:"is_admin"`
	Permissions   []string  `json:"permissions"`
}
