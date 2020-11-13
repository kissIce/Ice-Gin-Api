package entity

type UserCoupons struct {
	Id              uint64  `json:"id" form:"id" gorm:"primarykey"`
	Uid             uint64  `json:"uid" form:"uid"`
	Cid             uint64  `json:"cid" form:"cid"`
	CouponSn        string  `json:"coupon_sn" form:"coupon_sn"`
	CouponName      string  `json:"coupon_name" form:"coupon_name"`
	CouponColor     string  `json:"coupon_color" form:"coupon_color"`
	CouponType      int8    `json:"coupon_type" form:"coupon_type"`
	CouponMinMoney  float64 `json:"coupon_min_money" form:"coupon_min_money"`
	CouponDiscount  float64 `json:"coupon_discount" form:"coupon_discount"`
	CouponDerate    float64 `json:"coupon_derate" form:"coupon_derate"`
	CouponStartTime uint32  `json:"coupon_start_time" form:"coupon_start_time"`
	CouponEndTime   uint32  `json:"coupon_end_time" form:"coupon_end_time"`
	CouponRemark    string  `json:"coupon_remark" form:"coupon_remark"`
	Status          int8    `json:"status" form:"status"`
	CreatedAt       uint32  `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt       uint32  `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt       uint32  `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
