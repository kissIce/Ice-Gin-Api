package entity

type Prize struct {
  Model
  ActivityId int32 `json:"activity_id" form:"activity_id"`
  Type int8 `json:"type" form:"type"`
  TypeField string `json:"type_field" form:"type_field"`
  Val int32 `json:"val" form:"val"`
  Name string `json:"name" form:"name"`
  Desc string `json:"desc" form:"desc"`
  Img string `json:"img" form:"img"`
  Num int32 `json:"num" form:"num"`
  Level int8 `json:"level" form:"level"`
  Percent int8 `json:"percent" form:"percent"`
  Ordid int32 `json:"ordid" form:"ordid"`
}