package entity

type AuthMenu struct {
	MenuId uint64 `gorm:"comment:菜单id;"`
	RoleId uint64 `gorm:"comment:角色id;"`
}