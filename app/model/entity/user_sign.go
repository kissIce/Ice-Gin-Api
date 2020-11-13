package entity

type UserSign struct {
	Id             uint64 `json:"id" form:"id" gorm:"primarykey"`
	Uid            uint64 `json:"uid" form:"uid"`
	TotalSign      uint32  `json:"total_sign" form:"total_sign"`
	ContinuousSign uint32  `json:"continuous_sign" form:"continuous_sign"`
	SignNum        uint32  `json:"sign_num" form:"sign_num"`
	LastSign       uint32  `json:"last_sign" form:"last_sign"`
	CreatedAt      uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt      uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt      uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
