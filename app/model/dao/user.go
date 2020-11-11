package dao

import (
	"ice/app/model/entity"
	"ice/global"
)

func GetUserByPhone(u *entity.User) error {
	return global.IceDb.Where("phone = ?", u.Phone).First(&u).Error
}

func GetUserById(u *entity.User) error {
	return global.IceDb.Where("id = ?", u.Id).First(&u).Error
}

func CreateUser(u *entity.User) error {
	return global.IceDb.Create(&u).Error
}