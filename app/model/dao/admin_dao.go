package dao

import (
	"ice/app/model/entity"
	"ice/global"
)

func GetAdmin(where map[string]interface{}, field string) (err error, admin *entity.AdminEntity) {
	var adminEntity entity.AdminEntity
	err = global.IceDb.Where(where).Select(field).First(&adminEntity).Error
	return err, &adminEntity
}

