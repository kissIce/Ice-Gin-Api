package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ice/app/model/data"
	"ice/app/model/entity"
	"ice/app/service"
	"ice/helper"
	"ice/utils/response"
)

var err error

func LoginByPhone(u *entity.User) *response.Resp {
	var ret map[string]interface{}
	ret, err = data.GetUserByPhone(u.Phone, []string{"phone", "id", "username", "avatar", "im_acc", "im_pwd", "sex", "age"})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return response.InitErrCode(response.DbError)
		}
		//u.Username = "用户" + rand.Intn()
		if err = data.CreateUser(u); err != nil {
			return response.InitErrCode(response.DbError)
		}
		ret = helper.Struct2Map(&u)
		// 注册im账号
		// go do something
		//go func() {
		//
		//}()
	} else {
		if u.DeletedAt > 0 || u.Status > 0 {
			return response.InitErrMsg("用户禁止登陆")
		}
	}
	// 创建JWT
	ret["uid"] = ret["id"]
	delete(ret, "id")
	var jwt service.JWT
	ret["token"], _ = jwt.CreateToken(service.InitClaims(gin.H{"id": u.Id}, 3*24*3600))
	return response.InitSucc(ret)
}
