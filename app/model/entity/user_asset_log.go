package entity

type UserAssetLog struct {
	Id          uint64  `json:"id" form:"id"`
	Uid         uint64  `json:"uid" form:"uid"`
	Tuid        uint64  `json:"tuid" form:"tuid"`
	SourceId    uint64  `json:"source_id" form:"source_id"`
	SourceTable string  `json:"source_table" form:"source_table"`
	Op          int8    `json:"op" form:"op"`
	Type        uint16  `json:"type" form:"type"`
	AssetField  string  `json:"asset_field" form:"asset_field"`
	BeforeVal   float64 `json:"before_val" form:"before_val"`
	Money       float64 `json:"money" form:"money"`
	Remark      string  `json:"remark" form:"remark"`
	CreatedAt   uint32  `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt   uint32  `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt   uint32  `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
