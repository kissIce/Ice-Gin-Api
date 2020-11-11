package entity

type RewardGif struct {
 
  Id int32 `json:"id" form:"id"`
  Name string `json:"name" form:"name"`
  Bit int32 `json:"bit" form:"bit"`
  Img string `json:"img" form:"img"`
  Flash string `json:"flash" form:"flash"`
  FlashTime float64 `json:"flash_time" form:"flash_time"`
  IsVip int8 `json:"is_vip" form:"is_vip"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
  AdminId int32 `json:"admin_id" form:"admin_id"`
}