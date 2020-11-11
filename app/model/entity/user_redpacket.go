package entity

type UserRedpacket struct {
 
  Id int64 `json:"id" form:"id"`
  Uid int64 `json:"uid" form:"uid"`
  FirstAmount float64 `json:"first_amount" form:"first_amount"`
  NowAmount float64 `json:"now_amount" form:"now_amount"`
  Status int8 `json:"status" form:"status"`
  InviteNum int32 `json:"invite_num" form:"invite_num"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  DeletedAt int32 `json:"deleted_at" form:"deleted_at"`
}