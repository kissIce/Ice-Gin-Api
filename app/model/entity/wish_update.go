package entity

type WishUpdate struct {
 
  Id int64 `json:"id" form:"id"`
  Wid int64 `json:"wid" form:"wid"`
  Uid int64 `json:"uid" form:"uid"`
  Title string `json:"title" form:"title"`
  Img string `json:"img" form:"img"`
  Intro string `json:"intro" form:"intro"`
  Info string `json:"info" form:"info"`
  Num int8 `json:"num" form:"num"`
  Status int8 `json:"status" form:"status"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  DeletedAt int32 `json:"deleted_at" form:"deleted_at"`
}