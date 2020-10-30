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
	Status         int8   `json:"status" gorm:"comment:状态 0：离职； 1：在职 ；2： 禁用"`
	LastLoginIp    string `gorm:"comment:最后登录ip"`
	LastLoginTime  int64  `gorm:"comment:最后登录时间"`
	LastLoginCount int8   `gorm:"comment:登陆次数"`
}

