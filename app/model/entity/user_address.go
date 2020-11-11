package entity

type UserAddress struct {
 
  Id int64 `json:"id" form:"id"`
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
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  DeletedAt int32 `json:"deleted_at" form:"deleted_at"`
}