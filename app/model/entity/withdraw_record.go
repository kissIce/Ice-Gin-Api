package entity

type WithdrawRecord struct {
	Id            uint64  `json:"id" form:"id" gorm:"primarykey"`
	Uid           uint64  `json:"uid" form:"uid"`
	OrderSn       string  `json:"order_sn" form:"order_sn"`
	Amount        float64 `json:"amount" form:"amount"`
	SuccessAmount float64 `json:"success_amount" form:"success_amount"`
	Account       string  `json:"account" form:"account"`
	AccountType   int8    `json:"account_type" form:"account_type"`
	Free          float64 `json:"free" form:"free"`
	FreeType      int8    `json:"free_type" form:"free_type"`
	FreeVal       float64 `json:"free_val" form:"free_val"`
	Status        int8    `json:"status" form:"status"`
	Remark        string  `json:"remark" form:"remark"`
	CreatedAt     uint32  `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt     uint32  `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt     uint32  `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
