package entity

type Article struct {
	Id         uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Cid        uint64 `json:"cid" form:"cid" gorm:"type:bigint(20) unsigned;comment:所属栏目id"`
	Title      string `json:"title" form:"title" gorm:"type:varchar(50);default:'';comment:文章标题"`
	Color      string `json:"color" form:"color" gorm:"type:varchar(10);default:'';comment:颜色"`
	ShortTitle string `json:"short_title" form:"short_title" gorm:"type:varchar(50);default:'';comment:副标题"`
	Alias      string `json:"alias" form:"alias" gorm:"type:varchar(50);default:'';comment:别名"`
	Img        string `json:"img" form:"img" gorm:"type:varchar(150);default:'';comment:图片"`
	Intro      string `json:"intro" form:"intro" gorm:"type:varchar(100);default:'';comment:简介"`
	Info       string `json:"info" form:"info" gorm:"type:varchar(150);default:'';comment:详情"`
	Comment    uint32 `json:"comment" form:"comment" gorm:"type:int(10)unsigned;default:0;comment:评论数"`
	Hits       uint32 `json:"hits" form:"hits" gorm:"type:int(10) unsigned;default:0;comment:点击数"`
	Sort       uint16 `json:"sort" form:"sort" gorm:"type:int(20) unsigned;comment:排序 越小越靠前"`
	Status     int8   `json:"status" form:"status" gorm:"type:tinyint(2) unsigned;comment:状态 0无效 1有效"`
	SeoTitle   string `json:"seo_title" form:"seo_title" gorm:"type:varchar(50);default:'';comment:seo标题"`
	SeoKeys    string `json:"seo_keys" form:"seo_keys" gorm:"type:varchar(80);default:'';comment:seo关键字"`
	SeoDesc    string `json:"seo_desc" form:"seo_desc" gorm:"type:varchar(150);default:'';comment:seo描述"`
	Author     string `json:"author" form:"author" gorm:"type:varchar(20);default:'';comment:作者"`
	Origin     string `json:"origin" form:"origin" gorm:"type:varchar(20);default:'';comment:来源"`
	Tags       string `json:"tags" form:"tags" gorm:"type:varchar(50);default:'';comment:标签"`
	Url        string `json:"url" form:"url" gorm:"type:varchar(150);default:'';comment:跳转地址"`
	CreatedAt  uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt  uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt  uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
