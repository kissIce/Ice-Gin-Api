package dao

import (
	"ice/app/model/entity"
	"ice/global"
)

func GetTzoneByWhere(where map[string]interface{}) ([]entity.Tzone, error) {
	var t []entity.Tzone
	err := global.IceDb.Where(where).Find(&t).Error
	return t, err
}
