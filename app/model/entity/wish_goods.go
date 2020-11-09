package entity

type WishGoods struct {
  Model
  Wid int64 `json:"wid" form:"wid"`
  Uid int64 `json:"uid" form:"uid"`
  Name string `json:"name" form:"name"`
  Price float64 `json:"price" form:"price"`
  Desc string `json:"desc" form:"desc"`
  Adv string `json:"adv" form:"adv"`
  Domain string `json:"domain" form:"domain"`
  Img string `json:"img" form:"img"`
  SendType int64 `json:"send_type" form:"send_type"`
  Num int64 `json:"num" form:"num"`
  SalerNum int64 `json:"saler_num" form:"saler_num"`
  UserLimit int64 `json:"user_limit" form:"user_limit"`
  PlanDate string `json:"plan_date" form:"plan_date"`
}