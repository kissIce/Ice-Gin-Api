package entity

type UserReward struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Tid int64 `json:"tid" form:"tid"`
  Tuid int64 `json:"tuid" form:"tuid"`
  Gid int64 `json:"gid" form:"gid"`
  Gamount float64 `json:"gamount" form:"gamount"`
}