package entity

type UserLoginLog struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  LoginIp string `json:"login_ip" form:"login_ip"`
  LoginPort int64 `json:"login_port" form:"login_port"`
  LoginPlatform int64 `json:"login_platform" form:"login_platform"`
  LoginImei string `json:"login_imei" form:"login_imei"`
  LoginProvince string `json:"login_province" form:"login_province"`
  LoginCity string `json:"login_city" form:"login_city"`
}