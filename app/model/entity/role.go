package entity

type Role struct {
  Model
  Pid int64 `json:"pid" form:"pid"`
  Name string `json:"name" form:"name"`
}