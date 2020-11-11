package entity

type RechargeRecord struct {
 
  Id int32 `json:"id" form:"id"`
  Uid int32 `json:"uid" form:"uid"`
  Openid string `json:"openid" form:"openid"`
  OrderSn string `json:"order_sn" form:"order_sn"`
  OrderSub string `json:"order_sub" form:"order_sub"`
  Amount float64 `json:"amount" form:"amount"`
  RealAmount float64 `json:"real_amount" form:"real_amount"`
  NeedAmount float64 `json:"need_amount" form:"need_amount"`
  Type int8 `json:"type" form:"type"`
  PayType int8 `json:"pay_type" form:"pay_type"`
  Status int8 `json:"status" form:"status"`
  RefundSn string `json:"refund_sn" form:"refund_sn"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  RefundAt int32 `json:"refund_at" form:"refund_at"`
  Remark string `json:"remark" form:"remark"`
}