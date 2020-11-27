package dao

import (
	"ice/app/model/entity"
	"ice/global"
	"time"
)

func GetAdmin(where map[string]interface{}) (ret *entity.Admin, err error) {
	var u entity.Admin
	err = global.IceDb.Where(where).First(&u).Error
	return &u, err
}

func GetAdminList() (admins []entity.Admin, err error) {
	err = global.IceDb.Find(&admins).Error
	return
}

func AddAdmin(admin *entity.Admin) error {
	return global.IceDb.Create(&admin).Error
}

func EditAdmin(admin *entity.Admin) error {
	return global.IceDb.Updates(admin).Error
}

func DelAdmin(admin *entity.Admin) error {
	return global.IceDb.Model(admin).Where("id = ?", admin.Id).Update("deleted_at", time.Now().Unix()).Error
}