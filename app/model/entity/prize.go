package entity

type Prize struct {
 
  Id int32 `json:"id" form:"id"`
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
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  DeletedAt int32 `json:"deleted_at" form:"deleted_at"`
}