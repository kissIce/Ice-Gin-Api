package entity

type SmsRecord struct {
 
  Id int32 `json:"id" form:"id"`
  Phone string `json:"phone" form:"phone"`
  Status int8 `json:"status" form:"status"`
  Content string `json:"content" form:"content"`
  CreatedAt string `json:"created_at" form:"created_at"`
  AddIp int32 `json:"add_ip" form:"add_ip"`
}