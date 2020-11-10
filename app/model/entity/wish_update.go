package entity

type WishUpdate struct {
  Model
  Wid int64 `json:"wid" form:"wid"`
  Uid int64 `json:"uid" form:"uid"`
  Title string `json:"title" form:"title"`
  Img string `json:"img" form:"img"`
  Intro string `json:"intro" form:"intro"`
  Info string `json:"info" form:"info"`
  Num int8 `json:"num" form:"num"`
  Status int8 `json:"status" form:"status"`
}