package entity

type Activity struct {
	Id            uint64 `json:"id" form:"id" gorm:"primarykey"`
	Title         string `json:"title" form:"title"`
	StartTime     uint32 `json:"start_time" form:"start_time"`
	EndTime       uint32 `json:"end_time" form:"end_time"`
	PagePath      string `json:"page_path" form:"page_path"`
	ShareTitle    string `json:"share_title" form:"share_title"`
	ShareSubtitle string `json:"share_subtitle" form:"share_subtitle"`
	ShareImg      string `json:"share_img" form:"share_img"`
	ShareUrl      string `json:"share_url" form:"share_url"`
	Status        int8   `json:"status" form:"status"`
	AdminId       uint64 `json:"admin_id" form:"admin_id"`
	CreatedAt     uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt     uint32 `json:"updated_at" form:"updated_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	DeletedOn     uint32 `json:"deleted_on" form:"deleted_on" gorm:"default:0;comment:删除时间"`
}
