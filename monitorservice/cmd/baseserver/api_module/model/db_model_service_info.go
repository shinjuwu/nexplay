package model

import "time"

type ApiUrl struct {
	URL     string `json:"url"`      // api url
	Method  string `json:"method"`   // api 方法
	AuthKey string `json:"auth_key"` // Basic Auth key
}

type ServiceInfo struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	SubName      string    `json:"sub_name"`
	APIURLs      ApiUrl    `json:"api_urls"`
	APIURLsBytes []byte    `json:"-"`
	Info         string    `json:"info"`
	IsEnabled    bool      `json:"is_enabled"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
	DisableTime  time.Time `json:"disable_time"`
}
