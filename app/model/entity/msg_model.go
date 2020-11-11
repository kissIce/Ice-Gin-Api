package entity

type MsgModel struct {
 
  Id int32 `json:"id" form:"id"`
  Type int8 `json:"type" form:"type"`
  Tpl string `json:"tpl" form:"tpl"`
  Content string `json:"content" form:"content"`
  Status int8 `json:"status" form:"status"`
}