package entity

type Prize struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Aid       uint64 `json:"aid" form:"aid" gorm:"type:bigint(20) unsigned;comment:所属活动id"`
	SendType  uint8  `json:"send_type" form:"send_type" gorm:"type:tinyint(2) unsigned;default:0;comment:1线上发放 2线下"`
	Type      uint8  `json:"type" form:"type" gorm:"type:tinyint(2) unsigned;default:0;comment:1金豆 2余额"`
	Val       int32  `json:"val" form:"val" gorm:"type:int(10) unsigned;default:0;comment:线上奖品对应的值"`
	Name      string `json:"name" form:"name" gorm:"type:varchar(50);default:'';comment:奖品名称"`
	Desc      string `json:"desc" form:"desc" gorm:"type:varchar(200);default:'';comment:奖品描述"`
	Img       string `json:"img" form:"img" gorm:"type:varchar(150);default:'';comment:奖品图片"`
	Num       int32  `json:"num" form:"num" gorm:"type:int(10) unsigned;default:0;comment:奖品数量"`
	Level     uint8  `json:"level" form:"level" gorm:"type:tinyint(2) unsigned;default:0;comment:奖品等级 0则不计入中奖轮播"`
	Percent   uint8  `json:"percent" form:"percent" gorm:"type:tinyint(3) unsigned;default:0;comment:中奖概率"`
	Sort      uint32 `json:"sort" form:"sort" gorm:"type:int(10) unsigned;default:0;comment:奖品排序 越小越往前"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
