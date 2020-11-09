package entity

type UserVip struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Level int64 `json:"level" form:"level"`
  EndTime string `json:"end_time" form:"end_time"`
}