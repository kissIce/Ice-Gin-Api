package api

type LoginUserByPhone struct {
	Uid      uint64 `json:"uid"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	ImAcc    string `json:"im_acc"`
	ImPwd    string `json:"im_pwd"`
	Token    string `json:"token"`
}
