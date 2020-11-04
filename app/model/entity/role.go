package entity

type Role struct {
	Model
	Pid  uint64 `gorm:"default:0;comment:父级id"`
	Name string `gorm:"type:varchar(50);default:'';comment:角色名"`
}
