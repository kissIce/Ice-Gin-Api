package entity

type UserAddress struct {
	Id         uint64 `json:"id" form:"id" gorm:"primarykey"`
	Uid        uint64  `json:"uid" form:"uid"`
	Pid        uint32  `json:"pid" form:"pid"`
	Cid        uint32  `json:"cid" form:"cid"`
	Aid        uint32  `json:"aid" form:"aid"`
	Province   string `json:"province" form:"province"`
	City       string `json:"city" form:"city"`
	Area       string `json:"area" form:"area"`
	Addr       string `json:"addr" form:"addr"`
	Zonecode   string `json:"zonecode" form:"zonecode"`
	Realname   string `json:"realname" form:"realname"`
	Phone      string `json:"phone" form:"phone"`
	Postalcode string `json:"postalcode" form:"postalcode"`
	Default    int8   `json:"default" form:"default"`
	CreatedAt  uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt  uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt  uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
