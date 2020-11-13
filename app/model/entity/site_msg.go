package entity

type SiteMsg struct {
	Id        uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Uid       uint64 `json:"uid" form:"uid" gorm:"type:bigint(20) unsigned;comment:用户id"`
	Sid       uint64 `json:"sid" form:"sid" gorm:"type:bigint(20) unsigned;comment:原id"`
	Stable    string `json:"stable" form:"stable" gorm:"type:varchar(50);default:'';comment:原表"`
	Title     string `json:"title" form:"title" gorm:"type:varchar(50);default:'';comment:标题"`
	Intro     string `json:"intro" form:"intro" gorm:"type:varchar(120);default:'';comment:简介"`
	Info      string `json:"info" form:"info" gorm:"type:varchar(250);default:'';comment:内容"`
	Type      int8   `json:"type" form:"type" gorm:"type:tinyint(2) unsigned;default:0;comment:类型 0平台公告  1心愿审核通过 2心愿审核未通过 4心愿更新通过 5心愿更新未通过 6收到动态打赏 7 收到聊天打赏 8送出打赏 9置顶动态 10充值金豆 11众筹失败退回款 12 提现成功 13实名认证通过 14实名未通过 15心愿发货 16 举报已处理  17被举报通知"`
	IsRead    int8   `json:"is_read" form:"is_read" gorm:"type:tinyint(2) unsigned;default:0;comment:0未读 1已读"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
