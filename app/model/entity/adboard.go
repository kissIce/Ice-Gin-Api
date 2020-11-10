package entity

type Adboard struct {
  Model
  Name string `json:"name" form:"name"`
  Width int16 `json:"width" form:"width"`
  Height int16 `json:"height" form:"height"`
  Description string `json:"description" form:"description"`
  Status int8 `json:"status" form:"status"`
  AllowType string `json:"allow_type" form:"allow_type"`
}