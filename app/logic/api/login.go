package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ice/app/model/dao"
	"ice/app/model/data"
	"ice/app/model/entity"
	"ice/app/service"
	"ice/utils/response"
)

var err error

func LoginByPhone(u *entity.User) *response.Resp {
	u, err = data.GetUserByPhone(u.Phone)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return response.InitErrCode(response.DbError)
		}
		if err := dao.CreateUser(u);err != nil {
			return response.InitErrCode(response.DbError)
		}
		// 注册im账号
		// go do something
	} else {
		if u.DeletedAt > 0  || u.Status > 0 {
			return response.InitErrMsg("用户禁止登陆")
		}
	}
	// 创建JWT
	var jwt service.JWT
	token, _ := jwt.CreateToken(service.InitClaims(gin.H{"id": u.ID}, 3*24*3600))
	return response.InitSucc(token)
}

