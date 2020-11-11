package entity

type UserLoginLog struct {
 
  Id int64 `json:"id" form:"id"`
  Uid int64 `json:"uid" form:"uid"`
  LoginIp string `json:"login_ip" form:"login_ip"`
  LoginPort int32 `json:"login_port" form:"login_port"`
  LoginPlatform int8 `json:"login_platform" form:"login_platform"`
  LoginImei string `json:"login_imei" form:"login_imei"`
  LoginProvince string `json:"login_province" form:"login_province"`
  LoginCity string `json:"login_city" form:"login_city"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
}