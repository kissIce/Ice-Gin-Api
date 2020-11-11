package dao

import (
	"errors"
	"gorm.io/gorm"
	"ice/app/model/entity"
	"ice/global"
	"ice/helper"
)

func CountAdmin(where map[string]interface{}) (count int64) {
	global.IceDb.Where(where).Count(&count)
	return
}

func GetAdmin(where map[string]interface{}, field string) (err error, admin *entity.Admin) {
	var adminEntity entity.Admin
	err = global.IceDb.Where(where).Select(field).First(&adminEntity).Error
	return err, &adminEntity
}

func AddAdmin(admin *entity.Admin) (error, uint64) {
	err, _ := GetAdmin(map[string]interface{}{"phone": admin.Phone}, "id")
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册"), 0
	}
	admin.Password = helper.PwdHash(admin.Password)
	err = global.IceDb.Create(&admin).Error
	return err, admin.Id
}

func PhoneWhere(phone string) func (db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Where("phone = ?", phone)
	}
}