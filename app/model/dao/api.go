package dao

import (
	"ice/app/model/entity"
	"ice/app/service"
	"ice/global"
)

func GetApiList() (apis []entity.Api, err error) {
	err = global.IceDb.Find(&apis).Error
	return
}

func AddApi(api *entity.Api) error {
	return global.IceDb.Create(api).Error
}

func EditApi(api *entity.Api) (err error) {
	var old entity.Api
	db := global.IceDb.Model(api).Select([]string{"path", "method"}).Where("id = ?", api.Id).First(&old)
	err = db.Updates(api).Error
	err = service.UpdateCasbin(old.Path, api.Path, old.Method, api.Method)
	return
}

func DelApi(api *entity.Api) (err error) {
	err = global.IceDb.Delete(api).Error
	service.ClearCasbin(1, api.Path, api.Method)
	return err
}
