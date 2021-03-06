package dao

import (
	"ice/app/model/entity"
	"ice/global"
	"time"
)

func GetAdmin(where map[string]interface{}) (ent *entity.Admin, err error) {
	err = global.IceDb.Where(where).First(&ent).Error
	return ent, err
}

func GetAdminList() (ents []entity.Admin, err error) {
	err = global.IceDb.Find(&ents).Error
	return
}

func AddAdmin(ent *entity.Admin) error {
	return global.IceDb.Create(&ent).Error
}

func EditAdmin(ent *entity.Admin) error {
	return global.IceDb.Updates(ent).Error
}

func DelAdmin(ent *entity.Admin) error {
	return global.IceDb.Model(ent).Where("id = ?", ent.Id).Update("deleted_at", time.Now().Unix()).Error
}