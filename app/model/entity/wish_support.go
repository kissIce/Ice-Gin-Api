package entity

type WishSupport struct {
 
  Id int64 `json:"id" form:"id"`
  OrderSn string `json:"order_sn" form:"order_sn"`
  Wid int64 `json:"wid" form:"wid"`
  Wuid int64 `json:"wuid" form:"wuid"`
  Uid int64 `json:"uid" form:"uid"`
  Gid int64 `json:"gid" form:"gid"`
  Gname string `json:"gname" form:"gname"`
  Num int32 `json:"num" form:"num"`
  Price float64 `json:"price" form:"price"`
  Amount float64 `json:"amount" form:"amount"`
  CouponsId int64 `json:"coupons_id" form:"coupons_id"`
  NeedAmount float64 `json:"need_amount" form:"need_amount"`
  PayAmount float64 `json:"pay_amount" form:"pay_amount"`
  PayType int8 `json:"pay_type" form:"pay_type"`
  SendType int8 `json:"send_type" form:"send_type"`
  Virtual string `json:"virtual" form:"virtual"`
  IsUse int8 `json:"is_use" form:"is_use"`
  Realname string `json:"realname" form:"realname"`
  Phone string `json:"phone" form:"phone"`
  Addr string `json:"addr" form:"addr"`
  Remark string `json:"remark" form:"remark"`
  SysRemark string `json:"sys_remark" form:"sys_remark"`
  Status int8 `json:"status" form:"status"`
  ExpressName string `json:"express_name" form:"express_name"`
  ExpressCode string `json:"express_code" form:"express_code"`
  ExpressSn string `json:"express_sn" form:"express_sn"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  DeletedAt int32 `json:"deleted_at" form:"deleted_at"`
}