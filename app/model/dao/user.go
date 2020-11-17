package dao

import (
	"ice/app/model/entity"
	"ice/global"
)

func GetUserByWhere(where map[string]interface{}) (ret *entity.User, err error) {
	var u entity.User
	err = global.IceDb.Model(u).Where(where).First(&u).Error
	return &u, err
}

func CreateUser(u *entity.User) error {
	return global.IceDb.Create(&u).Error
}

