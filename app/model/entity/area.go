package entity

type Area struct {
  Model
  Pid int64 `json:"pid" form:"pid"`
  Area string `json:"area" form:"area"`
  Pinyin string `json:"pinyin" form:"pinyin"`
  FirstCode string `json:"first_code" form:"first_code"`
  Level int64 `json:"level" form:"level"`
  Status int64 `json:"status" form:"status"`
}