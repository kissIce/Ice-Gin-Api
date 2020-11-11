package entity

type Raffle struct {
 
  Id int32 `json:"id" form:"id"`
  Uid int32 `json:"uid" form:"uid"`
  Uname string `json:"uname" form:"uname"`
  Uavatar string `json:"uavatar" form:"uavatar"`
  ActivityId int32 `json:"activity_id" form:"activity_id"`
  PrizeId int32 `json:"prize_id" form:"prize_id"`
  PrizeName string `json:"prize_name" form:"prize_name"`
  PrizeType int8 `json:"prize_type" form:"prize_type"`
  PrizeField string `json:"prize_field" form:"prize_field"`
  PrizeVal int32 `json:"prize_val" form:"prize_val"`
  Status int8 `json:"status" form:"status"`
  IsThank int8 `json:"is_thank" form:"is_thank"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  UpdatedAt int32 `json:"updated_at" form:"updated_at"`
  DeletedAt int32 `json:"deleted_at" form:"deleted_at"`
}