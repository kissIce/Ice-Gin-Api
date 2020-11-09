package entity

type RewardGif struct {
  Model
  Name string `json:"name" form:"name"`
  Bit int64 `json:"bit" form:"bit"`
  Img string `json:"img" form:"img"`
  Flash string `json:"flash" form:"flash"`
  FlashTime float64 `json:"flash_time" form:"flash_time"`
  IsVip int64 `json:"is_vip" form:"is_vip"`
  AdminId int64 `json:"admin_id" form:"admin_id"`
}