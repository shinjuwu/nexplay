package captcha

import (
	"backend/pkg/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type ICaptcha interface {
	GenerateCaptcha() (string, string, string, error)
	VerifyCaptcha(string, string) bool
	GetCaptchaValue(string) string
	CheckExpired(string) bool
}

type CaptchaDigit struct {
	DebugMode          bool
	ExpiredSec         time.Duration
	captchaDriverDigit *base64Captcha.Captcha
}

func NewCaptcha(config config.Config) ICaptcha {
	return &CaptchaDigit{
		DebugMode:  config.GetApp().Env == gin.DebugMode,
		ExpiredSec: time.Duration(config.GetCaptcha().ExpiredSec) * time.Second,
		captchaDriverDigit: base64Captcha.NewCaptcha(
			base64Captcha.NewDriverDigit(
				config.GetCaptcha().ImgHeight,
				config.GetCaptcha().ImgWidth,
				config.GetCaptcha().KeyLength,
				0.7, // default as module base64Captcha
				80,  // default as module base64Captcha
			),
			NewMemoryStore(
				config.GetCaptcha().GCLimitNumber,
				time.Duration(config.GetCaptcha().ExpiredSec)*time.Second)),
	}
}

// 依照傳入的類型產生不同種類的驗證碼
func (p *CaptchaDigit) GenerateCaptcha() (string, string, string, error) {
	expiredTimeStr := time.Now().Local().UTC().Add(p.ExpiredSec).String()
	id, b64s, err := p.captchaDriverDigit.Generate()
	return id, b64s, expiredTimeStr, err
}

// 比對傳入的驗證 id 與驗證碼是否正確
func (p *CaptchaDigit) VerifyCaptcha(id string, value string) bool {
	return p.captchaDriverDigit.Verify(id, value, true)
}

// 取得實際的驗證碼, 不清除驗證碼資料
func (p *CaptchaDigit) GetCaptchaValue(id string) string {
	if p.DebugMode {
		return p.captchaDriverDigit.Store.Get(id, false)
	} else {
		return ""
	}
}

// 檢查驗證碼是否超時
func (p *CaptchaDigit) CheckExpired(id string) bool {
	expired := false
	val := p.captchaDriverDigit.Store.Get(id, false)
	if val == ErrCaptchaExpired.Error() {
		expired = true
	}

	return expired
}
