package entity

type AppUpdate struct {
	Id          uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Type        string `json:"type" form:"type" gorm:"type:varchar(10);default:'';comment:ios  android"`
	Version     string `json:"version" form:"version" gorm:"type:varchar(20);default:'';comment:版本号"`
	Title       string `json:"title" form:"title" gorm:"type:varchar(50);default:'';comment:更新标题"`
	Info        string `json:"info" form:"info" gorm:"type:varchar(250);default:'';comment:更新说明"`
	Force       uint8  `json:"force" form:"force" gorm:"type:tinyint(2);default:0;comment:0不强制更新 1强制更新"`
	DownloadUrl string `json:"download_url" form:"download_url" gorm:"type:varchar(50);default:'';comment:下载地址"`
	Status      int8   `json:"status" form:"status" gorm:"type:tinyint(2);default:0;comment:状态 0无效 1有效"`
	CreatedAt   uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt   uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt   uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
