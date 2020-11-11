package entity

type Admin struct {
	Id            uint64 `json:"id" form:"id" gorm:"primarykey"`
	Username      string `json:"username" form:"username"`
	Realname      string `json:"realname" form:"realname"`
	Password      string `json:"password" form:"password"`
	Avatar        string `json:"avatar" form:"avatar"`
	Phone         string `json:"phone" form:"phone"`
	Status        int8   `json:"status" form:"status"`
	RoleId        uint64  `json:"role_id" form:"role_id"`
	LastLoginIp   string `json:"last_login_ip" form:"last_login_ip"`
	LastLoginTime uint32  `json:"last_login_time" form:"last_login_time"`
	LoginCount    uint32  `json:"login_count" form:"login_count"`
	CreatedAt     uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt     uint32 `json:"updated_at" form:"updated_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	DeletedOn     uint32 `json:"deleted_on" form:"deleted_on" gorm:"default:0;comment:删除时间"`
}
