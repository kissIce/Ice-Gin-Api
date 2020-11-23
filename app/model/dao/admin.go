package dao

import (
	"ice/app/model/entity"
	"ice/global"
)

func GetAdminByWhere(where map[string]interface{}) (ret *entity.Admin, err error) {
	var u entity.Admin
	err = global.IceDb.Where(where).First(&u).Error
	return &u, err
}

func CreateAdmin(admin *entity.Admin) error {
	return global.IceDb.Create(&admin).Error
}

func UpdateAdminById(uid uint64, data map[string]interface{}) error {
	var admin entity.Admin
	return global.IceDb.Model(&admin).Where("id = ?", uid).Updates(data).Error
}