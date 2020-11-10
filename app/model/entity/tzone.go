package entity

type Tzone struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Content string `json:"content" form:"content"`
  Domain string `json:"domain" form:"domain"`
  Img string `json:"img" form:"img"`
  Video string `json:"video" form:"video"`
  VideoLen int8 `json:"video_len" form:"video_len"`
  Support int64 `json:"support" form:"support"`
  Comment int64 `json:"comment" form:"comment"`
  Reward int64 `json:"reward" form:"reward"`
  Transpond int64 `json:"transpond" form:"transpond"`
  Lng float64 `json:"lng" form:"lng"`
  Lat float64 `json:"lat" form:"lat"`
  City string `json:"city" form:"city"`
  District string `json:"district" form:"district"`
  Street string `json:"street" form:"street"`
  LastCuid string `json:"last_cuid" form:"last_cuid"`
  LastCuname string `json:"last_cuname" form:"last_cuname"`
  LastComment string `json:"last_comment" form:"last_comment"`
  LastRuid string `json:"last_ruid" form:"last_ruid"`
  LastRavatar string `json:"last_ravatar" form:"last_ravatar"`
  StickAt int32 `json:"stick_at" form:"stick_at"`
}