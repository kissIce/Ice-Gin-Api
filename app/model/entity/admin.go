package entity

type Admin struct {
  Model
  Username string `json:"username" form:"username"`
  Realname string `json:"realname" form:"realname"`
  Password string `json:"password" form:"password"`
  Avatar string `json:"avatar" form:"avatar"`
  Phone string `json:"phone" form:"phone"`
  Status int8 `json:"status" form:"status"`
  RoleId int64 `json:"role_id" form:"role_id"`
  LastLoginIp string `json:"last_login_ip" form:"last_login_ip"`
  LastLoginTime int32 `json:"last_login_time" form:"last_login_time"`
  LoginCount int32 `json:"login_count" form:"login_count"`
}