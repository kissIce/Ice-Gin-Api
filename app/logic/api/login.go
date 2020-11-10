package api

import (
	"errors"
	"gorm.io/gorm"
	"ice/app/model/dao"
	"ice/app/model/entity"
	"ice/utils/response"
)

func LoginByPhone(m *entity.User) *response.Resp {
	err := dao.GetUserByPhone(m)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

		} else {
			return response.InitErrCode(response.DbError)
		}
	}
	if m.DeletedAt > 0 {
		return response.InitErrMsg("用户已被删除")
	}
	if m.Status > 0 {
		return response.InitErrMsg("用户已被删除")
	}
	return response.InitSucc(m)
}
