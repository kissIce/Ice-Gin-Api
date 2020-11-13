package entity

type Admin struct {
	Id            uint64 `json:"id" form:"id" gorm:"primary_key;type:bigint(20) unsigned auto_increment;"`
	Username      string `json:"username" form:"username" gorm:"type:varchar(20);default:'';comment:用户名"`
	Realname      string `json:"realname" form:"realname" gorm:"type:varchar(20);default:'';comment:真实姓名"`
	Password      string `json:"password" form:"password" gorm:"type:varchar(150);default:'';comment:密码"`
	Avatar        string `json:"avatar" form:"avatar" gorm:"type:varchar(150);default:'';comment:头像"`
	Phone         string `json:"phone" form:"phone" gorm:"type:varchar(15);default:'';comment:手机号"`
	Status        int8   `json:"status" form:"status" gorm:"type:tinyint(3);default:0;comment:状态 0无效 1正常 2异常"`
	RoleId        uint64 `json:"role_id" form:"role_id" gorm:"type:bigint(20);default:0;comment:广告位状态 0未启用 1启用 "`
	LastLoginIp   string `json:"last_login_ip" form:"last_login_ip" gorm:"type:varchar(50);default:'';comment:登录ip"`
	LastLoginTime uint32 `json:"last_login_time" form:"last_login_time" gorm:"type:int(10) unsigned;default:0;comment:最后登录时间"`
	LoginCount    uint32 `json:"login_count" form:"login_count" gorm:"type:int(10) unsigned;default:0;comment:登录次数"`
	CreatedAt     uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt     uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt     uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
