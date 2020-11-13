package entity

type UserLoginLog struct {
	Id            uint64 `json:"id" form:"id" gorm:"primarykey"`
	Uid           uint64 `json:"uid" form:"uid"`
	LoginIp       string `json:"login_ip" form:"login_ip"`
	LoginPort     uint32 `json:"login_port" form:"login_port"`
	LoginPlatform int8   `json:"login_platform" form:"login_platform"`
	LoginImei     string `json:"login_imei" form:"login_imei"`
	LoginProvince string `json:"login_province" form:"login_province"`
	LoginCity     string `json:"login_city" form:"login_city"`
	CreatedAt     uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt     uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt     uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
