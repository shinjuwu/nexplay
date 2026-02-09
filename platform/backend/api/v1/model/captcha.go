package model

type CaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`     // 驗證碼id
	PicPath       string `json:"picPath"`       // base64 圖片字串
	CaptchaLength int    `json:"captchaLength"` // 驗證碼長度
	CaptchaValue  string `json:"captchaValue"`  // 驗證碼(測試模式才顯示)
	ExpiredTime   string `json:"expiredTime"`   // 驗證碼過期時間
}
