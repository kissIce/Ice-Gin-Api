package entity

type SmsRecord struct {
  Model
  Phone string `json:"phone" form:"phone"`
  Status int8 `json:"status" form:"status"`
  Content string `json:"content" form:"content"`
  AddIp int32 `json:"add_ip" form:"add_ip"`
}