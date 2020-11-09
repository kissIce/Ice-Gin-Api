package entity

type Comment struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Content string `json:"content" form:"content"`
  Sid int64 `json:"sid" form:"sid"`
  Suid int64 `json:"suid" form:"suid"`
  Scontent string `json:"scontent" form:"scontent"`
  Simg string `json:"simg" form:"simg"`
  Stable string `json:"stable" form:"stable"`
  Tcid int64 `json:"tcid" form:"tcid"`
  Tcuid int64 `json:"tcuid" form:"tcuid"`
  Tccontent string `json:"tccontent" form:"tccontent"`
  Cid int64 `json:"cid" form:"cid"`
  Cuid int64 `json:"cuid" form:"cuid"`
  Ccontent string `json:"ccontent" form:"ccontent"`
  IsRead int64 `json:"is_read" form:"is_read"`
  Support int64 `json:"support" form:"support"`
  Comment int64 `json:"comment" form:"comment"`
}