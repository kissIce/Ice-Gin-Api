package entity

type Coupons struct {
 
  Id int64 `json:"id" form:"id"`
  Name string `json:"name" form:"name"`
  Color string `json:"color" form:"color"`
  Num int32 `json:"num" form:"num"`
  SendNum int32 `json:"send_num" form:"send_num"`
  UserLimit int32 `json:"user_limit" form:"user_limit"`
  Type int8 `json:"type" form:"type"`
  MinMoney float64 `json:"min_money" form:"min_money"`
  Discount float64 `json:"discount" form:"discount"`
  Derate float64 `json:"derate" form:"derate"`
  StartTime int32 `json:"start_time" form:"start_time"`
  EndTime int32 `json:"end_time" form:"end_time"`
  Remark string `json:"remark" form:"remark"`
  Status int8 `json:"status" form:"status"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  DeletedAt int32 `json:"deleted_at" form:"deleted_at"`
}