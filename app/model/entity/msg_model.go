package entity

type MsgModel struct {
	Id      uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Type    int8   `json:"type" form:"type" gorm:"type:tinyint(2);default:1;comment:1 系统私信模板 2短信模板"`
	Tpl     string `json:"tpl" form:"tpl" gorm:"type:varchar(20);default:'';comment:模板代号"`
	Content string `json:"content" form:"content" gorm:"type:varchar(255);default:'';comment:模板内容"`
	Status  int8   `json:"status" form:"status" gorm:"type:tinyint(2);default:0;comment:1启用 0 未启用"`
}
