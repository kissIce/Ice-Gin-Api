package entity

type Wish struct {
	Id        uint64  `json:"id" form:"id" gorm:"primarykey"`
	Uid       uint64  `json:"uid" form:"uid"`
	Title     string  `json:"title" form:"title"`
	Intro     string  `json:"intro" form:"intro"`
	Info      string  `json:"info" form:"info"`
	Domain    string  `json:"domain" form:"domain"`
	Img       string  `json:"img" form:"img"`
	Coverplan string  `json:"coverplan" form:"coverplan"`
	Video     string  `json:"video" form:"video"`
	Money     float64 `json:"money" form:"money"`
	MoneyYes  float64 `json:"money_yes" form:"money_yes"`
	TagId     uint64  `json:"tag_id" form:"tag_id"`
	TagName   string  `json:"tag_name" form:"tag_name"`
	RateType  int8    `json:"rate_type" form:"rate_type"`
	Rate      float64 `json:"rate" form:"rate"`
	NeedFull  int8    `json:"need_full" form:"need_full"`
	Support   uint32  `json:"support" form:"support"`
	Comment   uint32  `json:"comment" form:"comment"`
	Transpond uint32  `json:"transpond" form:"transpond"`
	PreTime   uint32  `json:"pre_time" form:"pre_time"`
	StartTime uint32  `json:"start_time" form:"start_time"`
	EndTime   uint32  `json:"end_time" form:"end_time"`
	Status    int8    `json:"status" form:"status"`
	Reason    string  `json:"reason" form:"reason"`
	AuditTime uint32  `json:"audit_time" form:"audit_time"`
	StickAt   uint32  `json:"stick_at" form:"stick_at"`
	CreatedAt uint32  `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32  `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32  `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
