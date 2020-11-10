package entity

type UserSign struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  TotalSign int32 `json:"total_sign" form:"total_sign"`
  ContinuousSign int32 `json:"continuous_sign" form:"continuous_sign"`
  SignNum int32 `json:"sign_num" form:"sign_num"`
  LastSign int32 `json:"last_sign" form:"last_sign"`
}