package entity

type Raffle struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Uname string `json:"uname" form:"uname"`
  Uavatar string `json:"uavatar" form:"uavatar"`
  ActivityId int64 `json:"activity_id" form:"activity_id"`
  PrizeId int64 `json:"prize_id" form:"prize_id"`
  PrizeName string `json:"prize_name" form:"prize_name"`
  PrizeType int64 `json:"prize_type" form:"prize_type"`
  PrizeField string `json:"prize_field" form:"prize_field"`
  PrizeVal int64 `json:"prize_val" form:"prize_val"`
  Status int64 `json:"status" form:"status"`
  IsThank int64 `json:"is_thank" form:"is_thank"`
}