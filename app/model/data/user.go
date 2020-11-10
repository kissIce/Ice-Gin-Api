package data

import (
	"ice/app/model/dao"
	"ice/app/model/entity"
)

func GetUserByPhone(u *entity.User) error {
	err = dao.GetUserByPhone(u)
}
