package entity

type Raffle struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Uid       uint64 `json:"uid" form:"uid" gorm:"type:bigint(20) unsigned;comment:用户id"`
	Aid       uint64 `json:"aid" form:"aid" gorm:"type:bigint(20) unsigned;comment:所属活动id"`
	PrizeId   uint64 `json:"prize_id" form:"prize_id" gorm:"type:bigint(20) unsigned;comment:奖品id"`
	PrizeName string `json:"prize_name" form:"prize_name" gorm:"type:varchar(50);default:'';comment:所属活动id"`
	SendType  uint8  `json:"send_type" form:"send_type" gorm:"type:tinyint(2) unsigned;default:0;comment:1线上发放 2线下"`
	Type      uint8  `json:"type" form:"type" gorm:"type:tinyint(2) unsigned;default:0;comment:1金豆 2余额"`
	Val       int32  `json:"val" form:"val" gorm:"type:int(10) unsigned;default:0;comment:线上奖品对应的值"`
	Status    int8   `json:"status" form:"status" gorm:"type:tinyint(2) unsigned;comment:状态 0 无效 1有效"`
	IsThank   int8   `json:"is_thank" form:"is_thank" gorm:"type:tinyint(2) unsigned;comment:不计入轮播 0是  1不是 "`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
