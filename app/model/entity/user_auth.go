package entity

type UserAuth struct {
	Id        uint64 `json:"id" form:"id" gorm:"primarykey"`
	Uid       uint64 `json:"uid" form:"uid"`
	Realname  string `json:"realname" form:"realname"`
	Type      int8   `json:"type" form:"type"`
	CardNo    string `json:"card_no" form:"card_no"`
	CardFace  string `json:"card_face" form:"card_face"`
	CardBack  string `json:"card_back" form:"card_back"`
	Status    int8   `json:"status" form:"status"`
	Remark    string `json:"remark" form:"remark"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
