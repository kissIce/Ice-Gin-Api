package entity

type UserAuth struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Realname string `json:"realname" form:"realname"`
  Type int64 `json:"type" form:"type"`
  CardNo string `json:"card_no" form:"card_no"`
  CardFace string `json:"card_face" form:"card_face"`
  CardBack string `json:"card_back" form:"card_back"`
  Status int64 `json:"status" form:"status"`
  Remark string `json:"remark" form:"remark"`
}