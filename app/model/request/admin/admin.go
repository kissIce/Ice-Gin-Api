package admin

type LoginAdmin struct {
	Phone     string `json:"phone" form:"phone"`
	Password  string `json:"password" form:"password"`
	Captcha   string `json:"captcha" form:"captcha"`
	CaptchaId string `json:"captchaId" form:"captchaId"`
}
