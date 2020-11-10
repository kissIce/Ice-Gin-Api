package entity

type Raffle struct {
  Model
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
}