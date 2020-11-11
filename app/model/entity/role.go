package entity

type Role struct {
 
  Id int64 `json:"id" form:"id"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  DeletedAt int32 `json:"deleted_at" form:"deleted_at"`
  Pid int64 `json:"pid" form:"pid"`
  Name string `json:"name" form:"name"`
}