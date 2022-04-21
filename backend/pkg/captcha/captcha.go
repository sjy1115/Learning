package captcha

import (
	"github.com/mojocn/base64Captcha"
)

var captcha *base64Captcha.Captcha

func InitCaptcha() {
	captcha = base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, &RedisStore{})
}

func GenerateCaptcha() (string, string, error) {
	return captcha.Generate()
}

func VeryCaptcha(id, code string) bool {
	return captcha.Verify(id, code, true)
}
