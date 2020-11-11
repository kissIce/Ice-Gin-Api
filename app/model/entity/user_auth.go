package entity

type UserAuth struct {
 
  Id int64 `json:"id" form:"id"`
  Uid int64 `json:"uid" form:"uid"`
  Realname string `json:"realname" form:"realname"`
  Type int8 `json:"type" form:"type"`
  CardNo string `json:"card_no" form:"card_no"`
  CardFace string `json:"card_face" form:"card_face"`
  CardBack string `json:"card_back" form:"card_back"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  DeletedAt int32 `json:"deleted_at" form:"deleted_at"`
  Status int8 `json:"status" form:"status"`
  Remark string `json:"remark" form:"remark"`
}