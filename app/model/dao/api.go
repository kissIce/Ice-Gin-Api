package dao

import (
	"ice/app/model/entity"
	"ice/app/service"
	"ice/global"
)

func GetApiList() (ents []entity.Api, err error) {
	err = global.IceDb.Find(&ents).Error
	return
}

func AddApi(ent *entity.Api) error {
	return global.IceDb.Create(ent).Error
}

func EditApi(ent *entity.Api) (err error) {
	var old entity.Api
	db := global.IceDb.Model(ent).Select([]string{"path", "method"}).Where("id = ?", ent.Id).First(&old)
	err = db.Updates(ent).Error
	err = service.UpdateCasbin(old.Path, ent.Path, old.Method, ent.Method)
	return
}

func DelApi(ent *entity.Api) (err error) {
	err = global.IceDb.Delete(ent).Error
	service.ClearCasbin(1, ent.Path, ent.Method)
	return err
}
