package entity

type UserVip struct {
 
  Uid int32 `json:"uid" form:"uid"`
  Level int8 `json:"level" form:"level"`
  EndTime string `json:"end_time" form:"end_time"`
}