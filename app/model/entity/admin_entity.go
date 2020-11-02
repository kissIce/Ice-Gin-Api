package entity

import (
	"gorm.io/gorm"
)

type AdminEntity struct {
	gorm.Model
	Username       string `json:"username" gorm:"comment:用户名"`
	Realname       string `json:"realname" gorm:"comment:真实姓名"`
	Password       string `json:"password" gorm:"comment:密码"`
	Avatar         string `json:"avatar" gorm:"comment:用户头像"`
	Phone          string `json:"phone" gorm:"comment:手机号"`
	Status         int8   `json:"status" gorm:"default:1;comment:状态 0：离职； 1：在职 ；2： 禁用"`
	RoleId		   uint64 `json:"role_id" gorm:"comment:角色id"`
	LastLoginIp    string `gorm:"comment:最后登录ip"`
	LastLoginTime  uint64  `gorm:"comment:最后登录时间"`
	LastLoginCount uint8   `gorm:"comment:登陆次数"`
}

