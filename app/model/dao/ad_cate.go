package dao

import (
	"ice/app/model/entity"
	"ice/global"
	"time"
)

func GetAdCate(where map[string]interface{}) (ent *entity.AdCate, err error) {
	err = global.IceDb.Where(where).First(&ent).Error
	return ent, err
}

func GetAdCateList() (ents []entity.AdCate, err error) {
	err = global.IceDb.Find(&ents).Error
	return
}

func AddAdCate(ent *entity.AdCate) error {
	return global.IceDb.Create(&ent).Error
}

func EditAdCate(ent *entity.AdCate) error {
	return global.IceDb.Updates(ent).Error
}

func DelAdCate(ent *entity.AdCate) error {
	return global.IceDb.Model(ent).Where("id = ?", ent.Id).Update("deleted_at", time.Now().Unix()).Error
}