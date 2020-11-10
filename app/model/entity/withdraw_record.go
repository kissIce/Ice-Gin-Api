package entity

type WithdrawRecord struct {
  Model
  Uid int32 `json:"uid" form:"uid"`
  OrderSn string `json:"order_sn" form:"order_sn"`
  Amount int32 `json:"amount" form:"amount"`
  SuccessAmount int32 `json:"success_amount" form:"success_amount"`
  Account string `json:"account" form:"account"`
  AccountType int8 `json:"account_type" form:"account_type"`
  Free float64 `json:"free" form:"free"`
  FreeType int8 `json:"free_type" form:"free_type"`
  FreeVal float64 `json:"free_val" form:"free_val"`
  Status int8 `json:"status" form:"status"`
  Remark string `json:"remark" form:"remark"`
}