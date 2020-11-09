package entity

type UserZan struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Sid int64 `json:"sid" form:"sid"`
  Suid int64 `json:"suid" form:"suid"`
  Stable string `json:"stable" form:"stable"`
  Scontent string `json:"scontent" form:"scontent"`
  Simg string `json:"simg" form:"simg"`
  Tcid int64 `json:"tcid" form:"tcid"`
  Tcuid int64 `json:"tcuid" form:"tcuid"`
  Tccontent string `json:"tccontent" form:"tccontent"`
  IsRead int64 `json:"is_read" form:"is_read"`
}