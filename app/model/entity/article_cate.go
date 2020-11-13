package entity

type ArticleCate struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Pid       uint64 `json:"pid" form:"pid" gorm:"type:bigint(20) unsigned;default:0;comment:父级id"`
	Spid      string `json:"spid" form:"spid" gorm:"type:varchar(100);default:'';comment:所有父级id"`
	Type      uint8  `json:"type" form:"type" gorm:"type:tinyint(3) unsigned;default:0;comment:栏目类型"`
	Name      string `json:"name" form:"name" gorm:"type:varchar(20);default:'';comment:栏目名称"`
	Alias     string `json:"alias" form:"alias" gorm:"type:varchar(20);default:'';comment:栏目别名"`
	Img       string `json:"img" form:"img" gorm:"type:varchar(150);default:'';comment:图片"`
	Sort      uint16 `json:"sort" form:"sort" gorm:"type:int(10) unsigned;comment:排序越小越靠前"`
	Status    int8   `json:"status" form:"status" gorm:"type:tinyint(2) unsigned;default:0;comment:状态0无效 1有效"`
	IsHome    uint8  `json:"is_home" form:"is_home" gorm:"type:tinyint(2) unsigned;default:0;comment:0首页 1不在首页"`
	Url       string `json:"url" form:"url" gorm:"type:varchar(150);default:''comment:所属栏目id"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
