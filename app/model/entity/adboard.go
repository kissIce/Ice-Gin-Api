package entity

type Adboard struct {
	Id          uint64 `json:"id" form:"id" gorm:"primarykey"`
	Name        string `json:"name" form:"name"`
	Width       uint16  `json:"width" form:"width"`
	Height      uint16  `json:"height" form:"height"`
	Description string `json:"description" form:"description"`
	Status      int8   `json:"status" form:"status"`
	AllowType   string `json:"allow_type" form:"allow_type"`
	CreatedAt   uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt   uint32 `json:"updated_at" form:"updated_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	DeletedOn   uint32 `json:"deleted_on" form:"deleted_on" gorm:"default:0;comment:删除时间"`
}
