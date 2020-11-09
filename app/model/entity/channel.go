package entity

type Channel struct {
  Model
  Name string `json:"name" form:"name"`
  Icon string `json:"icon" form:"icon"`
  RegUrl string `json:"reg_url" form:"reg_url"`
}