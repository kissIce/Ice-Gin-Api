package entity

type UserReward struct {
 
  Id int64 `json:"id" form:"id"`
  Uid int64 `json:"uid" form:"uid"`
  Tid int64 `json:"tid" form:"tid"`
  Tuid int64 `json:"tuid" form:"tuid"`
  Gid int32 `json:"gid" form:"gid"`
  Gamount float64 `json:"gamount" form:"gamount"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  DeletedAt int32 `json:"deleted_at" form:"deleted_at"`
}