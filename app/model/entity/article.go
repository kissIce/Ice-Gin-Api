package entity

type Article struct {
	Id         uint64 `json:"id" form:"id" gorm:"primarykey"`
	Cid        uint64 `json:"cid" form:"cid"`
	Title      string `json:"title" form:"title"`
	Colors     string `json:"colors" form:"colors"`
	ShortTitle string `json:"short_title" form:"short_title"`
	Alias      string `json:"alias" form:"alias"`
	Img        string `json:"img" form:"img"`
	Intro      string `json:"intro" form:"intro"`
	Info       string `json:"info" form:"info"`
	Comments   uint32 `json:"comments" form:"comments"`
	Hits       uint32 `json:"hits" form:"hits"`
	Ordid      uint16 `json:"ordid" form:"ordid"`
	LastTime   uint32 `json:"last_time" form:"last_time"`
	Status     int8   `json:"status" form:"status"`
	SeoTitle   string `json:"seo_title" form:"seo_title"`
	SeoKeys    string `json:"seo_keys" form:"seo_keys"`
	SeoDesc    string `json:"seo_desc" form:"seo_desc"`
	Templet    string `json:"templet" form:"templet"`
	Author     string `json:"author" form:"author"`
	Origin     string `json:"origin" form:"origin"`
	Tags       string `json:"tags" form:"tags"`
	Url        string `json:"url" form:"url"`
	CreatedAt  uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt  uint32 `json:"updated_at" form:"updated_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	DeletedOn  uint32 `json:"deleted_on" form:"deleted_on" gorm:"default:0;comment:删除时间"`
}
