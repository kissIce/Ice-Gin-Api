package entity

type SiteMsg struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Sid int64 `json:"sid" form:"sid"`
  Title string `json:"title" form:"title"`
  Info string `json:"info" form:"info"`
  Type int64 `json:"type" form:"type"`
  Intro string `json:"intro" form:"intro"`
  IsRead int64 `json:"is_read" form:"is_read"`
}