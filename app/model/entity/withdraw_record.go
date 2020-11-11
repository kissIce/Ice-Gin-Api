package entity

type WithdrawRecord struct {
 
  Id int64 `json:"id" form:"id"`
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
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  Remark string `json:"remark" form:"remark"`
}