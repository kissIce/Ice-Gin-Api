package entity

type UserSign struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  TotalSign int64 `json:"total_sign" form:"total_sign"`
  ContinuousSign int64 `json:"continuous_sign" form:"continuous_sign"`
  SignNum int64 `json:"sign_num" form:"sign_num"`
  LastSign int64 `json:"last_sign" form:"last_sign"`
}