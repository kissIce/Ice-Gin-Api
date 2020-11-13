package entity

type SysNotice struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Title     string `json:"title" form:"title" gorm:"type:varchar(20);default:'';comment:标题"`
	Intro     string `json:"intro" form:"intro" gorm:"type:varchar(50);default:'';comment:简介"`
	Info      string `json:"info" form:"info" gorm:"type:varchar(150);default:'';comment:详情"`
	Status    uint8   `json:"status" form:"status" gorm:"type:tinyint(2) unsigned;default:0;comment:状态 0草稿 1发布"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
