package entity

type UserRedpacket struct {
	Id          uint64  `json:"id" form:"id" gorm:"primarykey"`
	Uid         uint64  `json:"uid" form:"uid"`
	FirstAmount float64 `json:"first_amount" form:"first_amount"`
	NowAmount   float64 `json:"now_amount" form:"now_amount"`
	Status      int8    `json:"status" form:"status"`
	InviteNum   uint32   `json:"invite_num" form:"invite_num"`
	CreatedAt   uint32  `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt   uint32  `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt   uint32  `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
