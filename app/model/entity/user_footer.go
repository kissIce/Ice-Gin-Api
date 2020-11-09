package entity

type UserFooter struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Tuid int64 `json:"tuid" form:"tuid"`
  IsRead int64 `json:"is_read" form:"is_read"`
}