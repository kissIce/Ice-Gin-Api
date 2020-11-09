package entity

type Wish struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Title string `json:"title" form:"title"`
  Intro string `json:"intro" form:"intro"`
  Info string `json:"info" form:"info"`
  Domain string `json:"domain" form:"domain"`
  Img string `json:"img" form:"img"`
  Coverplan string `json:"coverplan" form:"coverplan"`
  Video string `json:"video" form:"video"`
  Money float64 `json:"money" form:"money"`
  MoneyYes float64 `json:"money_yes" form:"money_yes"`
  TagId int64 `json:"tag_id" form:"tag_id"`
  TagName string `json:"tag_name" form:"tag_name"`
  RateType int64 `json:"rate_type" form:"rate_type"`
  Rate float64 `json:"rate" form:"rate"`
  NeedFull int64 `json:"need_full" form:"need_full"`
  Support int64 `json:"support" form:"support"`
  Comment int64 `json:"comment" form:"comment"`
  Transpond int64 `json:"transpond" form:"transpond"`
  PreTime int64 `json:"pre_time" form:"pre_time"`
  StartTime int64 `json:"start_time" form:"start_time"`
  EndTime int64 `json:"end_time" form:"end_time"`
  Status int64 `json:"status" form:"status"`
  Reason string `json:"reason" form:"reason"`
  AuditTime int64 `json:"audit_time" form:"audit_time"`
  StickAt int64 `json:"stick_at" form:"stick_at"`
}