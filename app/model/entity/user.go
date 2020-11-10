package entity

type User struct {
  Model
  Username string `json:"username" form:"username"`
  TzoneImg string `json:"tzone_img" form:"tzone_img"`
  Avatar string `json:"avatar" form:"avatar"`
  Sex int8 `json:"sex" form:"sex"`
  Age int8 `json:"age" form:"age"`
  Phone string `json:"phone" form:"phone"`
  WxOpenid string `json:"wx_openid" form:"wx_openid"`
  ZfbOpenid string `json:"zfb_openid" form:"zfb_openid"`
  AreaId int32 `json:"area_id" form:"area_id"`
  Area string `json:"area" form:"area"`
  Signature string `json:"signature" form:"signature"`
  ImAcc string `json:"im_acc" form:"im_acc"`
  ImPwd string `json:"im_pwd" form:"im_pwd"`
  Identity int16 `json:"identity" form:"identity"`
  Vip int8 `json:"vip" form:"vip"`
  RecomId int64 `json:"recom_id" form:"recom_id"`
  Status int8 `json:"status" form:"status"`
  Channel int32 `json:"channel" form:"channel"`
}