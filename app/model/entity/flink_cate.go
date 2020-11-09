package entity

type FlinkCate struct {
  Model
  Name string `json:"name" form:"name"`
  Ordid int64 `json:"ordid" form:"ordid"`
  Status int64 `json:"status" form:"status"`
}