package entity

type Article struct {
  Model
  CateId int16 `json:"cate_id" form:"cate_id"`
  Title string `json:"title" form:"title"`
  Colors string `json:"colors" form:"colors"`
  ShortTitle string `json:"short_title" form:"short_title"`
  Alias string `json:"alias" form:"alias"`
  Img string `json:"img" form:"img"`
  Intro string `json:"intro" form:"intro"`
  Info string `json:"info" form:"info"`
  Comments int32 `json:"comments" form:"comments"`
  Hits int32 `json:"hits" form:"hits"`
  Ordid int8 `json:"ordid" form:"ordid"`
  LastTime int32 `json:"last_time" form:"last_time"`
  Status int8 `json:"status" form:"status"`
  SeoTitle string `json:"seo_title" form:"seo_title"`
  SeoKeys string `json:"seo_keys" form:"seo_keys"`
  SeoDesc string `json:"seo_desc" form:"seo_desc"`
  Templet string `json:"templet" form:"templet"`
  Author string `json:"author" form:"author"`
  Origin string `json:"origin" form:"origin"`
  Tags string `json:"tags" form:"tags"`
  Url string `json:"url" form:"url"`
}