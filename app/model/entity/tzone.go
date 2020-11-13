package entity

type Tzone struct {
	Id          uint64  `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Uid         uint64  `json:"uid" form:"uid" gorm:"type:bigint(20) unsigned;default:0;comment:用户id"`
	Content     string  `json:"content" form:"content" gorm:"type:varchar(300);default:'';comment:内容"`
	Domain      string  `json:"domain" form:"domain" gorm:"type:varchar(300);default:'';comment:内容"`
	Img         string  `json:"img" form:"img" gorm:"type:varchar(300);default:'';comment:内容"`
	Video       string  `json:"video" form:"video" gorm:"type:varchar(300);default:'';comment:内容"`
	VideoLen    uint8   `json:"video_len" form:"video_len" gorm:"type:varchar(300);default:'';comment:内容"`
	Support     uint32  `json:"support" form:"support" gorm:"type:varchar(300);default:'';comment:内容"`
	Comment     uint32  `json:"comment" form:"comment" gorm:"type:varchar(300);default:'';comment:内容"`
	Reward      uint32  `json:"reward" form:"reward" gorm:"type:varchar(300);default:'';comment:内容"`
	Transpond   uint32  `json:"transpond" form:"transpond" gorm:"type:varchar(300);default:'';comment:内容"`
	Lng         float64 `json:"lng" form:"lng" gorm:"type:varchar(300);default:'';comment:内容"`
	Lat         float64 `json:"lat" form:"lat" gorm:"type:varchar(300);default:'';comment:内容"`
	City        string  `json:"city" form:"city" gorm:"type:varchar(300);default:'';comment:内容"`
	District    string  `json:"district" form:"district" gorm:"type:varchar(300);default:'';comment:内容"`
	Street      string  `json:"street" form:"street" gorm:"type:varchar(300);default:'';comment:内容"`
	LastCuid    string  `json:"last_cuid" form:"last_cuid" gorm:"type:varchar(300);default:'';comment:内容"`
	LastCuname  string  `json:"last_cuname" form:"last_cuname" gorm:"type:varchar(300);default:'';comment:内容"`
	LastComment string  `json:"last_comment" form:"last_comment" gorm:"type:varchar(300);default:'';comment:内容"`
	LastRuid    string  `json:"last_ruid" form:"last_ruid" gorm:"type:varchar(300);default:'';comment:内容"`
	LastRavatar string  `json:"last_ravatar" form:"last_ravatar" gorm:"type:varchar(300);default:'';comment:内容"`
	StickAt     uint32  `json:"stick_at" form:"stick_at" gorm:"type:varchar(300);default:'';comment:内容"`
	CreatedAt   uint32  `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt   uint32  `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt   uint32  `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
