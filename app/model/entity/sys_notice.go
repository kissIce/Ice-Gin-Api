package entity

type SysNotice struct {
  Model
  Title string `json:"title" form:"title"`
  Intro string `json:"intro" form:"intro"`
  Info string `json:"info" form:"info"`
  Status int8 `json:"status" form:"status"`
}