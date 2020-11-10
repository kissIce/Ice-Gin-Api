package entity

type Area struct {
  Model
  Pid int32 `json:"pid" form:"pid"`
  Area string `json:"area" form:"area"`
  Pinyin string `json:"pinyin" form:"pinyin"`
  FirstCode string `json:"first_code" form:"first_code"`
  Level int8 `json:"level" form:"level"`
  Status int8 `json:"status" form:"status"`
}