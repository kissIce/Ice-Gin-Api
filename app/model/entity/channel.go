package entity

type Channel struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Name      string `json:"name" form:"name" gorm:"type:varchar(50);default:'';comment:渠道名称"`
	Icon      string `json:"icon" form:"icon" gorm:"type:varchar(150);default:'';comment:渠道icon"`
	RegUrl    string `json:"reg_url" form:"reg_url" gorm:"type:varchar(200);default:'';comment:渠道地址"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
