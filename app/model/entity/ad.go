package entity

type Ad struct {
  Model
  CateId int64 `json:"cate_id" form:"cate_id"`
  Name string `json:"name" form:"name"`
  Url string `json:"url" form:"url"`
  Img string `json:"img" form:"img"`
  Type int64 `json:"type" form:"type"`
  Desc string `json:"desc" form:"desc"`
  StartTime int64 `json:"start_time" form:"start_time"`
  EndTime int64 `json:"end_time" form:"end_time"`
  Clicks int64 `json:"clicks" form:"clicks"`
  Ordid int64 `json:"ordid" form:"ordid"`
  Status int64 `json:"status" form:"status"`
}