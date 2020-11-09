package entity

type UserRedpacketHelp struct {
  Model
  Rid int64 `json:"rid" form:"rid"`
  Uid int64 `json:"uid" form:"uid"`
  Amount float64 `json:"amount" form:"amount"`
}