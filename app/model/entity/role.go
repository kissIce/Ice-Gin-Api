package entity

type Role struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Pid       uint64 `json:"pid" form:"pid" gorm:"type:bigint(20) unsigned;default:0;comment:父级id"`
	Name      string `json:"name" form:"name" gorm:"type:varchar(50);default:'';comment:角色名"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
