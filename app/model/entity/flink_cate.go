package entity

type FlinkCate struct {
 
  Id int16 `json:"id" form:"id"`
  Name string `json:"name" form:"name"`
  Ordid int8 `json:"ordid" form:"ordid"`
  Status int8 `json:"status" form:"status"`
}