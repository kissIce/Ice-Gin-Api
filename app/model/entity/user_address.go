package entity

type UserAddress struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Pid int32 `json:"pid" form:"pid"`
  Cid int32 `json:"cid" form:"cid"`
  Aid int32 `json:"aid" form:"aid"`
  Province string `json:"province" form:"province"`
  City string `json:"city" form:"city"`
  Area string `json:"area" form:"area"`
  Addr string `json:"addr" form:"addr"`
  Zonecode string `json:"zonecode" form:"zonecode"`
  Realname string `json:"realname" form:"realname"`
  Phone string `json:"phone" form:"phone"`
  Postalcode string `json:"postalcode" form:"postalcode"`
  Default int8 `json:"default" form:"default"`
}