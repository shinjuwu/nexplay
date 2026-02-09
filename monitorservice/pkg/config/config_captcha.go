package config

// CaptchaConfig is configuration generate pic of base64 information.
type CaptchaConfig struct {
	KeyLength     int `yaml:"key_length" json:"key_length"`           // 驗證碼長度
	ImgWidth      int `yaml:"img_width" json:"img_width"`             // 驗證碼寬度
	ImgHeight     int `yaml:"img_height" json:"img_height"`           // 驗證碼高度
	GCLimitNumber int `yaml:"gc_limit_number" json:"gc_limit_number"` // 驗證碼存儲記憶體大小
	ExpiredSec    int `yaml:"expired_sec" json:"expired_sec"`         // 驗證碼存活時間(秒)
}

// NewJwtConfig creates a new CaptchaConfig struct.
func NewCaptchaConfig() *CaptchaConfig {
	return &CaptchaConfig{
		KeyLength:     6,
		ImgWidth:      240,
		ImgHeight:     80,
		GCLimitNumber: 1024,
		ExpiredSec:    60,
	}
}
