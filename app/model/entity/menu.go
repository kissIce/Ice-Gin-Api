package entity

type Menu struct {
	Model
	Pid       uint64 `gorm:"default:0;comment:父级id;"`
	Title     string `gorm:"type:varchar(80);default:'';comment:标题;"`
	Name      string `gorm:"type:varchar(80);default:'';comment:名称;"`
	Icon      string `gorm:"type:varchar(80);default:'';comment:Icon;"`
	Path      string `gorm:"type:varchar(200);default:'';comment:路劲;"`
	Component string `gorm:"type:varchar(200);default:'';comment:组件路劲;"`
	KeepAlive uint8  `gorm:"default:0;comment:是否keepalive;"`
	Hidden    uint8  `gorm:"default:0;comment:是否隐藏;"`
	Sort      uint16 `gorm:"default:255;comment:排序越小越靠前;"`
	Default   uint8  `gorm:"default:0;comment:是否是默认菜单;"`
}

