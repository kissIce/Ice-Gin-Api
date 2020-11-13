package entity

type UserUsually struct {
	Uid             uint64 `json:"uid" form:"uid"`
	UsuallyIp       string `json:"usually_ip" form:"usually_ip"`
	UsuallyPlatform string `json:"usually_platform" form:"usually_platform"`
	UsuallyImei     string `json:"usually_imei" form:"usually_imei"`
	UsuallyProvince string `json:"usually_province" form:"usually_province"`
	UsuallyCity     string `json:"usually_city" form:"usually_city"`
	UsuallyTime     string `json:"usually_time" form:"usually_time"`
}
