package entity

type SmsRecord struct {
  Model
  Phone string `json:"phone" form:"phone"`
  Status int64 `json:"status" form:"status"`
  Content string `json:"content" form:"content"`
  AddIp int64 `json:"add_ip" form:"add_ip"`
}