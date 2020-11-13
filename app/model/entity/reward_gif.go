package entity

type RewardGif struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Name      string  `json:"name" form:"name" gorm:"type:varchar(50) unsigned;default:0;comment:礼物名称"`
	Bit       float64 `json:"bit" form:"bit" gorm:"type:decimal(15,2) unsigned;default:0;comment:礼物价格"`
	Img       string  `json:"img" form:"img" gorm:"type:varchar(150) unsigned;default:'';comment:礼物图片"`
	Flash     string  `json:"flash" form:"flash" gorm:"type:varchar(150)  unsigned;default:'';comment:礼物动画"`
	FlashTime float64 `json:"flash_time" form:"flash_time" gorm:"type:decimal(4,2) unsigned;default:0;comment:礼物动画时长"`
	IsVip     int8    `json:"is_vip" form:"is_vip" gorm:"type:tinyint(2) unsigned;default:0;comment:1vip礼物 0普通"`
	AdminId   uint64  `json:"admin_id" form:"admin_id" gorm:"type:bigint(20) unsigned;default:0;comment:操作人员"`
	CreatedAt uint32  `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32  `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32  `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}