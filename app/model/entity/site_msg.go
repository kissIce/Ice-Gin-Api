package entity

type SiteMsg struct {
 
  Id int64 `json:"id" form:"id"`
  Uid int64 `json:"uid" form:"uid"`
  Sid int64 `json:"sid" form:"sid"`
  Title string `json:"title" form:"title"`
  Info string `json:"info" form:"info"`
  Type int8 `json:"type" form:"type"`
  Intro string `json:"intro" form:"intro"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  IsRead int8 `json:"is_read" form:"is_read"`
}