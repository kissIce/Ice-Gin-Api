package entity

type AuthMenu struct {
	Menu
	MenuId     string          `gorm:"comment:菜单ID"`
	RoleId     string          `gorm:"comment:角色ID"`
	Children   []AuthMenu      `json:"children" gorm:"-"`
}
