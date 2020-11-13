package entity

type WishGoods struct {
	Id        uint64  `json:"id" form:"id" gorm:"primarykey"`
	Wid       uint64  `json:"wid" form:"wid"`
	Uid       uint64  `json:"uid" form:"uid"`
	Name      string  `json:"name" form:"name"`
	Price     float64 `json:"price" form:"price"`
	Desc      string  `json:"desc" form:"desc"`
	Adv       string  `json:"adv" form:"adv"`
	Domain    string  `json:"domain" form:"domain"`
	Img       string  `json:"img" form:"img"`
	SendType  int8    `json:"send_type" form:"send_type"`
	Num       int32   `json:"num" form:"num"`
	SalerNum  uint32  `json:"saler_num" form:"saler_num"`
	UserLimit int32   `json:"user_limit" form:"user_limit"`
	PlanDate  string  `json:"plan_date" form:"plan_date"`
	CreatedAt uint32  `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32  `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32  `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
