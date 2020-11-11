package entity

type Menu struct {
 
  Id int64 `json:"id" form:"id"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  DeletedAt int32 `json:"deleted_at" form:"deleted_at"`
  Pid int64 `json:"pid" form:"pid"`
  Title string `json:"title" form:"title"`
  Name string `json:"name" form:"name"`
  Icon string `json:"icon" form:"icon"`
  Path string `json:"path" form:"path"`
  Component string `json:"component" form:"component"`
  KeepAlive int8 `json:"keep_alive" form:"keep_alive"`
  Hidden int8 `json:"hidden" form:"hidden"`
  Sort int16 `json:"sort" form:"sort"`
  Default int8 `json:"default" form:"default"`
}