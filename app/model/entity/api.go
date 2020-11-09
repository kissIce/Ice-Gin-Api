package entity

type Api struct {
  Model
  Group string `json:"group" form:"group"`
  Path string `json:"path" form:"path"`
  Method string `json:"method" form:"method"`
}