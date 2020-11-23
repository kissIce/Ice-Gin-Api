package admin

type LoginAdminByPhone struct {
	AdminId uint64 `json:"admin_id"`
	Phone   string `json:"phone"`
	Token   string `json:"token"`
}
