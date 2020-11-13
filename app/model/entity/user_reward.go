package entity

type UserReward struct {
	Id        uint64  `json:"id" form:"id" gorm:"primarykey"`
	Uid       uint64  `json:"uid" form:"uid"`
	Tid       uint64  `json:"tid" form:"tid"`
	Tuid      uint64  `json:"tuid" form:"tuid"`
	Gid       uint64  `json:"gid" form:"gid"`
	Amount    float64 `json:"amount" form:"amount"`
	CreatedAt uint32  `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32  `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32  `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
