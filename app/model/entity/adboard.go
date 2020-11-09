package entity

type Adboard struct {
  Model
  Name string `json:"name" form:"name"`
  Width int64 `json:"width" form:"width"`
  Height int64 `json:"height" form:"height"`
  Description string `json:"description" form:"description"`
  Status int64 `json:"status" form:"status"`
  AllowType string `json:"allow_type" form:"allow_type"`
}