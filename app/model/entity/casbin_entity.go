package entity

type CasbinEntity struct {
	Id     uint64  `json:"id" gorm:"primarykey"`
	PType  string `json:"p_type" gorm:"column:p_type"`
	RoleId uint64  `json:"role_id" gorm:"column:role_id"`
	Path   string `json:"path" gorm:"column:path"`
	Method string `json:"method" gorm:"column:method"`
}
