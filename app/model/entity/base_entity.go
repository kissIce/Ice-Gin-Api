package entity

type Model struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt uint32 `gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `gorm:"autoUpdateTime;default:0;comment:更新时间"`
	DeletedAt uint32 `gorm:"default:0;comment:删除时间"`
}