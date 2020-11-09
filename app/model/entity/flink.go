package entity

type Flink struct {
  Model
  Name string `json:"name" form:"name"`
  Img string `json:"img" form:"img"`
  Url string `json:"url" form:"url"`
  CateId int64 `json:"cate_id" form:"cate_id"`
  Ordid int64 `json:"ordid" form:"ordid"`
  Status int64 `json:"status" form:"status"`
}