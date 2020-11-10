package entity

type ArticleCate struct {
  Model
  Type int8 `json:"type" form:"type"`
  Name string `json:"name" form:"name"`
  Alias string `json:"alias" form:"alias"`
  Img string `json:"img" form:"img"`
  Pid int16 `json:"pid" form:"pid"`
  Spid string `json:"spid" form:"spid"`
  Ordid int16 `json:"ordid" form:"ordid"`
  Status int8 `json:"status" form:"status"`
  SeoTitle string `json:"seo_title" form:"seo_title"`
  SeoKeys string `json:"seo_keys" form:"seo_keys"`
  SeoDesc string `json:"seo_desc" form:"seo_desc"`
  IndexTemplet string `json:"index_templet" form:"index_templet"`
  CategoryTemplet string `json:"category_templet" form:"category_templet"`
  DetailTemplet string `json:"detail_templet" form:"detail_templet"`
  IsHome int8 `json:"is_home" form:"is_home"`
  Url string `json:"url" form:"url"`
}