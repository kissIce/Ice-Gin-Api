package entity

type SmsRecord struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Phone     string `json:"phone" form:"phone" gorm:"type:varchar(15);not null;comment:手机号"`
	Status    int8   `json:"status" form:"status" gorm:"type:tinyint(2);default:0;comment:状态 0失败 1成功"`
	Content   string `json:"content" form:"content" gorm:"type:varchar(300);default:'';comment:短信内容"`
	AddIp     string  `json:"add_ip" form:"add_ip" gorm:"type:varchar(50);default:'';comment:发送ip"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
