package admin

type LoginAdmin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Captcha  string `json:"captcha" form:"captcha"`
	CaptchaId  string `json:"captchaId" form:"captchaId"`
}

type RegisterAdmin struct {

}