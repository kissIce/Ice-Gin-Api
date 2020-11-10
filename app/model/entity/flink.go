package entity

type Flink struct {
  Model
  Name string `json:"name" form:"name"`
  Img string `json:"img" form:"img"`
  Url string `json:"url" form:"url"`
  CateId int16 `json:"cate_id" form:"cate_id"`
  Ordid int8 `json:"ordid" form:"ordid"`
  Status int8 `json:"status" form:"status"`
}