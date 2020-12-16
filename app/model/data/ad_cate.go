package data

import (
	"ice/app/model/dao"
	"ice/app/model/entity"
)

func GetAdCateList() (adCateList []entity.AdCate, err error) {
	return dao.GetAdCateList()
}

func AddAdCate(adCate *entity.AdCate) error {
	return dao.AddAdCate(adCate)
}

func EditAdCate(adCate *entity.AdCate) error {
	return dao.EditAdCate(adCate)
}

func DelAdCate(adCate *entity.AdCate) error {
	return dao.DelAdCate(adCate)
}