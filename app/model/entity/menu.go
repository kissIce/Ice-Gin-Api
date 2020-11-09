package entity

type Menu struct {
  Model
  Pid int64 `json:"pid" form:"pid"`
  Title string `json:"title" form:"title"`
  Name string `json:"name" form:"name"`
  Icon string `json:"icon" form:"icon"`
  Path string `json:"path" form:"path"`
  Component string `json:"component" form:"component"`
  KeepAlive int64 `json:"keep_alive" form:"keep_alive"`
  Hidden int64 `json:"hidden" form:"hidden"`
  Sort int64 `json:"sort" form:"sort"`
  Default int64 `json:"default" form:"default"`
}