package dao

import (
	"ice/app/model/entity"
	"ice/global"
	"time"
)

func GetAd(where map[string]interface{}) (ent *entity.Ad, err error) {
	err = global.IceDb.Where(where).First(&ent).Error
	return ent, err
}

func GetAdList() (ents []entity.Ad, err error) {
	err = global.IceDb.Find(&ents).Error
	return
}

func AddAd(ent *entity.Ad) error {
	return global.IceDb.Create(&ent).Error
}

func EditAd(ent *entity.Ad) error {
	return global.IceDb.Updates(ent).Error
}

func DelAd(ent *entity.Ad) error {
	return global.IceDb.Model(ent).Where("id = ?", ent.Id).Update("deleted_at", time.Now().Unix()).Error
}
