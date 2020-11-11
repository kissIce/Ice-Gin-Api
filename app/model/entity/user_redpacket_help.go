package entity

type UserRedpacketHelp struct {
 
  Id int64 `json:"id" form:"id"`
  Rid int64 `json:"rid" form:"rid"`
  Uid int64 `json:"uid" form:"uid"`
  Amount float64 `json:"amount" form:"amount"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  DeletedAt int32 `json:"deleted_at" form:"deleted_at"`
}