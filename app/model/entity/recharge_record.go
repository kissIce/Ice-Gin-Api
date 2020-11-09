package entity

type RechargeRecord struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Openid string `json:"openid" form:"openid"`
  OrderSn string `json:"order_sn" form:"order_sn"`
  OrderSub string `json:"order_sub" form:"order_sub"`
  Amount float64 `json:"amount" form:"amount"`
  RealAmount float64 `json:"real_amount" form:"real_amount"`
  NeedAmount float64 `json:"need_amount" form:"need_amount"`
  Type int64 `json:"type" form:"type"`
  PayType int64 `json:"pay_type" form:"pay_type"`
  Status int64 `json:"status" form:"status"`
  RefundSn string `json:"refund_sn" form:"refund_sn"`
  RefundAt int64 `json:"refund_at" form:"refund_at"`
  Remark string `json:"remark" form:"remark"`
}