package service

import (
	"github.com/mojocn/base64Captcha"
	"ice/global"
)

var store = base64Captcha.DefaultMemStore

func Captcha() (id, img string, err error) {
	driver := base64Captcha.NewDriverDigit(global.IceConfig.Captcha.Height, global.IceConfig.Captcha.Width, global.IceConfig.Captcha.Length, 0.5, 60)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, img, err = cp.Generate()
	return
}

func CaptchaVerify(id string, code string) bool {
	return store.Verify(id, code, true)
}
