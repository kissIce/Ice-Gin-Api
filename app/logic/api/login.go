package api

import (
	"errors"
	"gorm.io/gorm"
	"ice/app/model/data"
	"ice/app/model/entity"
	"ice/utils/response"
)

var err error

func LoginByPhone(u *entity.User) *response.Resp {
	var ret map[string]interface{}
	ret, err = data.GetUserByPhone(u.Phone, []string{"phone", "id", "sex", "age"})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return response.InitErrCode(response.DbError)
		}
		if err := data.CreateUser(u);err != nil {
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
	//var jwt service.JWT
	//token, _ := jwt.CreateToken(service.InitClaims(gin.H{"id": u.Id}, 3*24*3600))
	return response.InitSucc(ret)
}

