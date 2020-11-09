package entity

type WithdrawRecord struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  OrderSn string `json:"order_sn" form:"order_sn"`
  Amount int64 `json:"amount" form:"amount"`
  SuccessAmount int64 `json:"success_amount" form:"success_amount"`
  Account string `json:"account" form:"account"`
  AccountType int64 `json:"account_type" form:"account_type"`
  Free float64 `json:"free" form:"free"`
  FreeType int64 `json:"free_type" form:"free_type"`
  FreeVal float64 `json:"free_val" form:"free_val"`
  Status int64 `json:"status" form:"status"`
  Remark string `json:"remark" form:"remark"`
}