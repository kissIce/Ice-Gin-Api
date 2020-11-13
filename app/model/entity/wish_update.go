package entity

type WishUpdate struct {
	Id        uint64 `json:"id" form:"id" gorm:"primarykey"`
	Wid       uint64 `json:"wid" form:"wid"`
	Uid       uint64 `json:"uid" form:"uid"`
	Title     string `json:"title" form:"title"`
	Img       string `json:"img" form:"img"`
	Intro     string `json:"intro" form:"intro"`
	Info      string `json:"info" form:"info"`
	Num       uint8  `json:"num" form:"num"`
	Status    int8   `json:"status" form:"status"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
