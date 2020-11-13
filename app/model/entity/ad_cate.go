package entity

type AdCate struct {
	Id          uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Name        string `json:"name" form:"name" gorm:"type:varchar(50);default:'';comment:广告位名称"`
	Width       uint16 `json:"width" form:"width" gorm:"type:smallint(5) unsigned;default:0;comment:广告位宽度"`
	Height      uint16 `json:"height" form:"height" gorm:"type:smallint(5) unsigned;default:0;comment:广告位高度"`
	Description string `json:"description" form:"description" gorm:"type:varchar(250);default:'';comment:广告位描述"`
	Status      int8   `json:"status" form:"status" gorm:"type:tinyint(2);default:0;comment:广告位状态 0未启用 1启用 "`
	AllowType   string `json:"allow_type" form:"allow_type" gorm:"type:varchar(50);default:'';comment:广告位允许的类型"`
	CreatedAt   uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt   uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt   uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
