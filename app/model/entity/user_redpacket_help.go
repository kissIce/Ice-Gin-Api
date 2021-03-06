package entity

type UserRedpacketHelp struct {
	Id        uint64  `json:"id" form:"id" gorm:"primarykey"`
	Rid       uint64  `json:"rid" form:"rid"`
	Uid       uint64  `json:"uid" form:"uid"`
	Amount    float64 `json:"amount" form:"amount"`
	CreatedAt uint32  `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32  `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32  `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
