package entity

type Comment struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Uid       uint64 `json:"uid" form:"uid" gorm:"type:bigint(20) unsigned;default:0;comment:用户id"`
	Content   string `json:"content" form:"content" gorm:"type:varchar(300);default:'';comment:评论内容"`
	Sid       uint64 `json:"sid" form:"sid" gorm:"type:bigint(20) unsigned;default:0;comment:原id"`
	Suid      uint64 `json:"suid" form:"suid" gorm:"type:bigint(20) unsigned;default:0;comment:原用户id"`
	Scontent  string `json:"scontent" form:"scontent" gorm:"type:varchar(300);default:'';comment:原内容"`
	Simg      string `json:"simg" form:"simg" gorm:"type:varchar(150);default:'';comment:原图片"`
	Stable    string `json:"stable" form:"stable" gorm:"type:varchar(50);default:'';comment:原表名"`
	Tcid      uint64 `json:"tcid" form:"tcid" gorm:"type:bigint(20) unsigned;default:0;comment:顶级评论id"`
	Tcuid     uint64 `json:"tcuid" form:"tcuid" gorm:"type:bigint(20) unsigned;default:0;comment:顶级评论uid"`
	Tccontent string `json:"tccontent" form:"tccontent" gorm:"type:varchar(300);default:'';comment:顶级评论内容"`
	Cid       uint64 `json:"cid" form:"cid" gorm:"type:bigint(20) unsigned;default:0;comment:被评论id"`
	Cuid      uint64 `json:"cuid" form:"cuid" gorm:"type:bigint(20) unsigned;default:0;comment:被评论uid"`
	Ccontent  string `json:"ccontent" form:"ccontent" gorm:"type:varchar(300);default:'';comment:被评论内容"`
	IsRead    uint8  `json:"is_read" form:"is_read" gorm:"type:tinyint(2);default:0;comment:被评论者是否已读 0未读 1已读"`
	Support   uint32 `json:"support" form:"support" gorm:"type:int(10) unsigned;default:0;comment:点赞数"`
	Comment   uint32 `json:"comment" form:"comment" gorm:"type:int(10) unsigned;default:0;comment:评论数"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
