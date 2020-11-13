package entity

type Menu struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Pid       uint64  `json:"pid" form:"pid" gorm:"type:bigint(20) unsigned;default:0;comment:父级id"`
	Title     string `json:"title" form:"title" gorm:"type:varchar(50);default:'';comment:标题"`
	Name      string `json:"name" form:"name" gorm:"type:varchar(50);default:'';comment:名称"`
	Icon      string `json:"icon" form:"icon" gorm:"type:varchar(50);default:'';comment:栏目Icon"`
	Path      string `json:"path" form:"path" gorm:"type:varchar(200);default:'';comment:路径"`
	Component string `json:"component" form:"component" gorm:"type:varchar(200);default:'';comment:组件路径"`
	KeepAlive int8   `json:"keep_alive" form:"keep_alive" gorm:"type:tinyint(2);default:0;comment:0不keep 1keep"`
	Hidden    int8   `json:"hidden" form:"hidden" gorm:"type:tinyint(2);default:0;comment:0不隐藏 1隐藏"`
	Sort      int16  `json:"sort" form:"sort" gorm:"type:int(10) unsigned;default:0;comment:排序越小越靠前"`
	Default   int8   `json:"default" form:"default" gorm:"type:tinyint(2);default:0;comment:0不默认 1默认"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
