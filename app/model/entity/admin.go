package entity

type Admin struct {
	Model
	Username      string `json:"username" gorm:"default:'';comment:用户名"`
	Realname      string `json:"realname" gorm:"default:'';comment:真实姓名"`
	Password      string `json:"password" gorm:"default:'';comment:密码"`
	Avatar        string `json:"avatar" gorm:"default:'';comment:用户头像"`
	Phone         string `json:"phone" gorm:"unique;comment:手机号"`
	Status        int8   `json:"status" gorm:"default:0;comment:状态 0：离职； 1：在职 ；2： 禁用"`
	RoleId        uint64 `json:"role_id" gorm:"comment:角色id"`
	LastLoginIp   string `gorm:"comment:最后登录ip"`
	LastLoginTime uint32 `gorm:"comment:最后登录时间"`
	LoginCount    uint32  `gorm:"comment:登陆次数"`
}

