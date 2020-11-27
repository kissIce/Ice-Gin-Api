package admin

type LoginAdmin struct {
	Phone     string `json:"phone" form:"phone"`
	Password  string `json:"password" form:"password"`
	Captcha   string `json:"captcha" form:"captcha"`
	CaptchaId string `json:"captchaId" form:"captchaId"`
}

type AddAdmin struct {
	Username string `json:"username" form:"username"`
	Realname string `json:"realname" form:"realname"`
	Password string `json:"password" form:"password"`
	Avatar   string `json:"avatar" form:"avatar"`
	Phone    string `json:"phone" form:"phone"`
	Status   int8 `json:"status" form:"status"`
	RoleId   uint64 `json:"role_id" form:"role_id"`
}

type EditAdmin struct {
	IdReq
	AddAdmin
}