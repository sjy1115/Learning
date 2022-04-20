package captcha

import "github.com/mojocn/base64Captcha"

var captcha *base64Captcha.Captcha

func init() {
	captcha = base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, newRedisStore(nil))
}

func GenerateCaptcha() (string, string, error) {
	return captcha.Generate()
}

func VeryCaptcha(id, code string) bool {
	return captcha.Verify(id, code, true)
}
