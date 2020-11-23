package entity

type Role struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key"`
	Pid       uint64 `json:"pid" form:"pid" gorm:"type:bigint(20) unsigned;default:0;comment:父级id"`
	Name      string `json:"name" form:"name" gorm:"type:varchar(50);default:'';comment:角色名"`
	Children  []Role `json:"children" gorm:"-"`
	Menus     []Menu `json:"menus" gorm:"many2many:auth_menu;"`
	CreatedAt uint32 `json:"-" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"-" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"-" gorm:"default:0;comment:删除时间"`
}