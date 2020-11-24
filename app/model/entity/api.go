package entity

type Api struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Group     string `json:"group" form:"group" gorm:"type:varchar(50);default:'';comment:api组"`
	Path      string `json:"path" form:"path" gorm:"type:varchar(150);default:'';comment:请求路径"`
	Intro     string `json:"intro" form:"intro" gorm:"type:varchar(150);default:'';comment:api描述"`
	Method    string `json:"method" form:"method" gorm:"type:varchar(10);default:'';comment:请求方法"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
