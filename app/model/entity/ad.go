package entity

type Ad struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Cid       uint64 `json:"cid" form:"cid" gorm:"type:bigint(20) unsigned;comment:所属栏目id"`
	Name      string `json:"name" form:"name" gorm:"type:varchar(50);comment:名称;default:''"`
	Url       string `json:"url" form:"url" gorm:"type:varchar(250);comment:跳转链接;default:''"`
	Img       string `json:"img" form:"img" gorm:"type:varchar(250);comment:图片;default:''"`
	Type      int8   `json:"type" form:"type" gorm:"type:tinyint(2);comment:广告类型;default:0"`
	Desc      string `json:"desc" form:"desc" gorm:"type:varchar(250);comment:广告描述;default:''"`
	StartTime int32  `json:"start_time" form:"start_time" gorm:"type:int(10) unsigned;default:0;comment:开始时间"`
	EndTime   int32  `json:"end_time" form:"end_time" gorm:"type:int(10) unsigned;default:0;comment:结束时间"`
	Clicks    int32  `json:"clicks" form:"clicks" gorm:"type:int(10) unsigned;default:0;comment:点击次数"`
	Sort      uint16 `json:"sort" form:"sort" gorm:"type:int(10) unsigned;default:0;comment:排序越小越靠前"`
	Status    int8   `json:"status" form:"status" gorm:"type:tinyint(3);default:0;comment:状态 0未启用 1启用"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
