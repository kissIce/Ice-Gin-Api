package entity

type Activity struct {
	Id            uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Title         string `json:"title" form:"title" gorm:"type:varchar(50);comment:活动名称;default:''"`
	PagePath      string `json:"page_path" form:"page_path" gorm:"type:varchar(255);comment:页面路径;default:''"`
	ShareTitle    string `json:"share_title" form:"share_title" gorm:"type:varchar(50);comment:分享标题;default:''"`
	ShareSubtitle string `json:"share_subtitle" form:"share_subtitle" gorm:"type:varchar(80);comment:活动分享副标题;default:''"`
	ShareImg      string `json:"share_img" form:"share_img" gorm:"type:varchar(150);comment:分享小图片;default:''"`
	ShareUrl      string `json:"share_url" form:"share_url" gorm:"type:varchar(250);comment:分享链接;default:''"`
	StartTime     uint32 `json:"start_time" form:"start_time" gorm:"type:int(10) unsigned;default:0;comment:开始时间"`
	EndTime       uint32 `json:"end_time" form:"end_time" gorm:"type:int(10) unsigned;default:0;comment:结束时间"`
	Status        int8   `json:"status" form:"status" gorm:"type:tinyint(3) unsigned;comment:活动状态;default:0"`
	AdminId       uint64 `json:"admin_id" form:"admin_id" gorm:"type:bigint(20) unsigned;comment:添加者id;default:0"`
	CreatedAt     uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt     uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt     uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
