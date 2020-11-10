package entity

type Ad struct {
  Model
  CateId int16 `json:"cate_id" form:"cate_id"`
  Name string `json:"name" form:"name"`
  Url string `json:"url" form:"url"`
  Img string `json:"img" form:"img"`
  Type int8 `json:"type" form:"type"`
  Desc string `json:"desc" form:"desc"`
  StartTime int32 `json:"start_time" form:"start_time"`
  EndTime int32 `json:"end_time" form:"end_time"`
  Clicks int32 `json:"clicks" form:"clicks"`
  Ordid int8 `json:"ordid" form:"ordid"`
  Status int8 `json:"status" form:"status"`
}