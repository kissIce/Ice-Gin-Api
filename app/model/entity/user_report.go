package entity

type UserReport struct {
  Model
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
}