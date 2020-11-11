package entity

type AuthMenu struct {
	MenuId int64 `json:"menu_id" form:"menu_id"`
	RoleId int64 `json:"role_id" form:"role_id"`
}
