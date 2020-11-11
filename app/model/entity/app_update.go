package entity

type AppUpdate struct {
	Id          uint64 `json:"id" form:"id" gorm:"primarykey"`
	Type        string `json:"type" form:"type"`
	Version     string `json:"version" form:"version"`
	Title       string `json:"title" form:"title"`
	Info        string `json:"info" form:"info"`
	Force       uint8   `json:"force" form:"force"`
	DownloadUrl string `json:"download_url" form:"download_url"`
	Status      int8   `json:"status" form:"status"`
	CreatedAt   uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt   uint32 `json:"updated_at" form:"updated_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	DeletedOn   uint32 `json:"deleted_on" form:"deleted_on" gorm:"default:0;comment:删除时间"`
}
