package entity

type UserReport struct {
 
  Id int64 `json:"id" form:"id"`
  Uid int64 `json:"uid" form:"uid"`
  Sid int64 `json:"sid" form:"sid"`
  Suid int64 `json:"suid" form:"suid"`
  Stable string `json:"stable" form:"stable"`
  Scontent string `json:"scontent" form:"scontent"`
  Simg string `json:"simg" form:"simg"`
  Svideo string `json:"svideo" form:"svideo"`
  Sdomain string `json:"sdomain" form:"sdomain"`
  Domain string `json:"domain" form:"domain"`
  Img string `json:"img" form:"img"`
  Content string `json:"content" form:"content"`
  Status int8 `json:"status" form:"status"`
  Do int8 `json:"do" form:"do"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  DeletedAt int32 `json:"deleted_at" form:"deleted_at"`
}