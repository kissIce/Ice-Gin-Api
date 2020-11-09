package entity

type Coupons struct {
  Model
  Name string `json:"name" form:"name"`
  Color string `json:"color" form:"color"`
  Num int64 `json:"num" form:"num"`
  SendNum int64 `json:"send_num" form:"send_num"`
  UserLimit int64 `json:"user_limit" form:"user_limit"`
  Type int64 `json:"type" form:"type"`
  MinMoney float64 `json:"min_money" form:"min_money"`
  Discount float64 `json:"discount" form:"discount"`
  Derate float64 `json:"derate" form:"derate"`
  StartTime int64 `json:"start_time" form:"start_time"`
  EndTime int64 `json:"end_time" form:"end_time"`
  Remark string `json:"remark" form:"remark"`
  Status int64 `json:"status" form:"status"`
}