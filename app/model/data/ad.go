package data

import (
	"ice/app/model/dao"
	"ice/app/model/entity"
)

func GetAdList() (adList []entity.Ad, err error) {
	return dao.GetAdList()
}

func AddAd(ad *entity.Ad) error {
	return dao.AddAd(ad)
}

func EditAd(ad *entity.Ad) error {
	return dao.EditAd(ad)
}

func DelAd(ad *entity.Ad) error {
	return dao.DelAd(ad)
}