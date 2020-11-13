package entity

type RechargeRecord struct {
	Id          uint64  `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Uid         uint64  `json:"uid" form:"uid" gorm:"type:bigint(20) unsigned;default:0;comment:用户id"`
	OrderSn     string  `json:"order_sn" form:"order_sn" gorm:"type:varchar(80);not null;comment:订单编号"`
	OrderSub    string  `json:"order_sub" form:"order_sub" gorm:"type:varchar(80);default:'';comment:订单描述"`
	OrderAmount float64 `json:"order_amount" form:"order_amount" gorm:"type:decimal(15,2) unsigned;default:0;comment:订单金额"`
	NeedAmount  float64 `json:"need_amount" form:"need_amount" gorm:"type:decimal(15,2) unsigned;default:0;comment:需支付金额"`
	PayAmount   float64 `json:"real_amount" form:"real_amount" gorm:"type:decimal(15,2) unsigned;default:0;comment:实际回调支付金额"`
	Type        uint8    `json:"type" form:"type" gorm:"type:tinyint(3) unsigned;default:0;comment:类型 1 充值金豆  2 充值余额"`
	PayType     uint8    `json:"pay_type" form:"pay_type" gorm:"type:bigint(20) unsigned;default:0;comment:支付方式 1 支付宝 2微信 3苹果"`
	Status      int8    `json:"status" form:"status" gorm:"type:tinyint(3);default:0;comment:订单状态 -1 订单异常 0待支付 1支付成功 2已过期 3退款中 4已退款"`
	RefundSn    string  `json:"refund_sn" form:"refund_sn" gorm:"type:varchar(80);default:'';comment:退款单号"`
	RefundAt    uint32   `json:"refund_at" form:"refund_at" gorm:"type:int(10) unsigned;default:0;comment:退款时间"`
	SysRemark   string  `json:"remark" form:"remark" gorm:"type:varchar(80);default:'';comment:系统备注"`
	CreatedAt   uint32  `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt   uint32  `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt   uint32  `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
