package entity

type UserZan struct {
 
  Id int64 `json:"id" form:"id"`
  Uid int64 `json:"uid" form:"uid"`
  Sid int64 `json:"sid" form:"sid"`
  Suid int64 `json:"suid" form:"suid"`
  Stable string `json:"stable" form:"stable"`
  Scontent string `json:"scontent" form:"scontent"`
  Simg string `json:"simg" form:"simg"`
  Tcid int64 `json:"tcid" form:"tcid"`
  Tcuid int64 `json:"tcuid" form:"tcuid"`
  Tccontent string `json:"tccontent" form:"tccontent"`
  IsRead int8 `json:"is_read" form:"is_read"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  DeletedAt int32 `json:"deleted_at" form:"deleted_at"`
}