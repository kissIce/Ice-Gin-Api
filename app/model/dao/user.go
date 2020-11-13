package dao

import (
	"errors"
	"gorm.io/gorm"
	"ice/app/model/entity"
	"ice/global"
	"ice/helper"
)

func GetUserByWhere(where map[string]interface{}) (ret map[string]interface{}, err error) {
	var u entity.User
	err = global.IceDb.Model(u).Where(where).First(&u).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		ret = helper.Struct2Map(u)
	}
	return ret, err
}

func CreateUser(u *entity.User) error {
	return global.IceDb.Create(&u).Error
}

