package api

type LoginUserPhone struct {
	Phone string `json:"phone" form:"phone"`
	Code  string `json:"code" form:"code"`
}
