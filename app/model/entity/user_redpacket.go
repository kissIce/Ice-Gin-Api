package entity

type UserRedpacket struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  FirstAmount float64 `json:"first_amount" form:"first_amount"`
  NowAmount float64 `json:"now_amount" form:"now_amount"`
  Status int8 `json:"status" form:"status"`
  InviteNum int32 `json:"invite_num" form:"invite_num"`
}