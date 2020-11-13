package entity

type Coupons struct {
	Id        uint64  `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Name      string  `json:"name" form:"name" gorm:"type:varchar(50);default:'';comment:优惠券名称"`
	Color     string  `json:"color" form:"color" gorm:"type:varchar(10);default:'';comment:优惠券颜色"`
	Num       int32  `json:"num" form:"num" gorm:"type:int(10);default:-1;comment:发放数量，-1不限量"`
	SendNum   uint32  `json:"send_num" form:"send_num" gorm:"type:int(10) unsigned;default:0;comment:已领取数量"`
	UserLimit int32   `json:"user_limit" form:"user_limit" gorm:"type:int(10);default:1;comment:用户限领张数 -1代表不限量"`
	Type      int8    `json:"type" form:"type" gorm:"type:tinyint(2);default:0;comment:0无效类型 1满减券 2折扣券"`
	MinMoney  float64 `json:"min_money" form:"min_money" gorm:"type:decimal(15,2) unsigned;not null;comment:最低使用金额"`
	Discount  float64 `json:"discount" form:"discount" gorm:"type:decimal(5,2) unsigned;default:0;comment:折扣率"`
	Derate    float64 `json:"derate" form:"derate" gorm:"type:decimal(15,2) unsigned;default:0;comment:减免金额"`
	StartTime uint32  `json:"start_time" form:"start_time" gorm:"type:int(10) unsigned;default:0;comment:开始时间"`
	EndTime   uint32  `json:"end_time" form:"end_time" gorm:"type:int(10) unsigned;default:0;comment:结束时间"`
	Remark    string  `json:"remark" form:"remark" gorm:"type:varchar(150);default:'';comment:优惠券说明"`
	Status    int8    `json:"status" form:"status" gorm:"type:tinyint(2);default:0;comment:0无效 1有效"`
	CreatedAt uint32  `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32  `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32  `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
