package entity

type Setting struct {
  Model
  Name string `json:"name" form:"name"`
  Data string `json:"data" form:"data"`
  Remark string `json:"remark" form:"remark"`
}