package entity

type UserFocus struct {
	Id        uint64 `json:"id" form:"id" gorm:"primarykey"`
	Uid       uint64 `json:"uid" form:"uid"`
	Tuid      uint64 `json:"tuid" form:"tuid"`
	IsRead    int8   `json:"is_read" form:"is_read"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
