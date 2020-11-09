package entity

type Prize struct {
  Model
  ActivityId int64 `json:"activity_id" form:"activity_id"`
  Type int64 `json:"type" form:"type"`
  TypeField string `json:"type_field" form:"type_field"`
  Val int64 `json:"val" form:"val"`
  Name string `json:"name" form:"name"`
  Desc string `json:"desc" form:"desc"`
  Img string `json:"img" form:"img"`
  Num int64 `json:"num" form:"num"`
  Level int64 `json:"level" form:"level"`
  Percent int64 `json:"percent" form:"percent"`
  Ordid int64 `json:"ordid" form:"ordid"`
}