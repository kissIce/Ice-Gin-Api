package data

import (
	"ice/app/model/dao"
	"ice/app/model/entity"
)

func GetApiList() (apiList []entity.Api, err error) {
	return dao.GetApiList()
}

func AddApi(api *entity.Api) error {
	return dao.AddApi(api)
}

func EditApi(api *entity.Api) error {
	return dao.EditApi(api)
}

func DelApi(api *entity.Api) error {
	return dao.DelApi(api)
}