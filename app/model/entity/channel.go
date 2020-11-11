package entity

type Channel struct {
	Id        uint64 `json:"id" form:"id" gorm:"primarykey"`
	Name      string `json:"name" form:"name"`
	Icon      string `json:"icon" form:"icon"`
	RegUrl    string `json:"reg_url" form:"reg_url"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	DeletedOn uint32 `json:"deleted_on" form:"deleted_on" gorm:"default:0;comment:删除时间"`
}
