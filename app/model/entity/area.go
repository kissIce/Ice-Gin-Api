package entity

type Area struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Pid       uint64 `json:"pid" form:"pid" gorm:"type:bigint(20);default:0;comment:父级id"`
	Area      string `json:"area" form:"area" gorm:"type:varchar(100);default:'';comment:地区名称"`
	Pinyin    string `json:"pinyin" form:"pinyin" gorm:"type:varchar(200);default:'';comment:地区拼音"`
	FirstCode string `json:"first_code" form:"first_code" gorm:"type:char(1);default:'';comment:首字母"`
	Level     uint8  `json:"level" form:"level" gorm:"type:tinyint(3);not null;comment:城市等级"`
	Status    int8   `json:"status" form:"status" gorm:"type:tinyint(2);default:0;comment:状态 0无效 1有效"`
}
