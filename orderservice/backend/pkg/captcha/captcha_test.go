package captcha_test

import (
	"backend/pkg/captcha"
	"backend/pkg/config"
	"testing"
)

func TestCaptchaGenerate(t *testing.T) {
	config := config.NewConfig()
	c := captcha.NewCaptcha(config)

	id, b64s, _, err := c.GenerateCaptcha()
	if err != nil {
		t.Errorf("GenerateCaptcha err is %v", err)
	}

	t.Logf("GenerateCaptcha id is: %v", id)
	t.Logf("GenerateCaptcha b64s is: %v", b64s)

	// code := ""
	// verify := c.VerifyCaptcha(id, code)
	// t.Logf("VerifyCaptcha verify is: %v", verify)

}

func TestCaptchaVerify(t *testing.T) {
	config := config.NewConfig()
	c := captcha.NewCaptcha(config)

	id := "mXBRxQrj4yTKAGbzhhC8"
	code := "264051"
	verify := c.VerifyCaptcha(id, code)
	t.Logf("VerifyCaptcha verify is: %v", verify)

}
