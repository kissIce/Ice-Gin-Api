package admin

import (
	"ice/app/model/dao"
	"ice/app/model/entity"
)

func LoginByPhone(admin *entity.Admin) {

}

func RegAdmin(admin *entity.Admin) (err error, id uint64) {
	err, id = dao.AddAdmin(admin)
	return err, admin.Id
}