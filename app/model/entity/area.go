package entity

type Area struct {
	Id        uint64 `json:"id" form:"id" gorm:"primarykey"`
	Pid       uint64 `json:"pid" form:"pid"`
	Area      string `json:"area" form:"area"`
	Pinyin    string `json:"pinyin" form:"pinyin"`
	FirstCode string `json:"first_code" form:"first_code"`
	Level     uint8   `json:"level" form:"level"`
	Status    int8   `json:"status" form:"status"`
}
