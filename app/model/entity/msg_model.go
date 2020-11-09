package entity

type MsgModel struct {
  Model
  Type int64 `json:"type" form:"type"`
  Tpl string `json:"tpl" form:"tpl"`
  Content string `json:"content" form:"content"`
  Status int64 `json:"status" form:"status"`
}