package entity

type ArticleCate struct {
	Id              uint64 `json:"id" form:"id" gorm:"primarykey"`
	Pid             uint64 `json:"pid" form:"pid"`
	Type            uint8  `json:"type" form:"type"`
	Name            string `json:"name" form:"name"`
	Alias           string `json:"alias" form:"alias"`
	Img             string `json:"img" form:"img"`
	Spid            string `json:"spid" form:"spid"`
	Ordid           uint16 `json:"ordid" form:"ordid"`
	Status          int8   `json:"status" form:"status"`
	SeoTitle        string `json:"seo_title" form:"seo_title"`
	SeoKeys         string `json:"seo_keys" form:"seo_keys"`
	SeoDesc         string `json:"seo_desc" form:"seo_desc"`
	IndexTemplet    string `json:"index_templet" form:"index_templet"`
	CategoryTemplet string `json:"category_templet" form:"category_templet"`
	DetailTemplet   string `json:"detail_templet" form:"detail_templet"`
	IsHome          uint8  `json:"is_home" form:"is_home"`
	Url             string `json:"url" form:"url"`
	CreatedAt       uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt       uint32 `json:"updated_at" form:"updated_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	DeletedOn       uint32 `json:"deleted_on" form:"deleted_on" gorm:"default:0;comment:删除时间"`
}
