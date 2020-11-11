package entity

type UserSign struct {
 
  Id int64 `json:"id" form:"id"`
  Uid int64 `json:"uid" form:"uid"`
  TotalSign int32 `json:"total_sign" form:"total_sign"`
  ContinuousSign int32 `json:"continuous_sign" form:"continuous_sign"`
  SignNum int32 `json:"sign_num" form:"sign_num"`
  LastSign int32 `json:"last_sign" form:"last_sign"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
}