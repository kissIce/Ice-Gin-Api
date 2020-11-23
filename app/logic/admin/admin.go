package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ice/app/model/data"
	adminRequest "ice/app/model/request/admin"
	adminResponse "ice/app/model/response/admin"
	"ice/app/service"
	"ice/utils/response"
)

func LoginByPhone(r *adminRequest.LoginAdmin) *response.Resp {
	admin, err := data.GetAdminByPhone(r.Phone, []string{"id", "password", "avatar", "realname", "username", "status", "deleted_at"})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return response.InitErrCode(response.DbError)
		}
		return response.InitErrMsg("用户不存在")
	}
	if admin.Password != r.Password {
		return response.InitErrMsg("密码错误")
	}
	if admin.Status != 1 || admin.DeletedAt > 0 {
		return response.InitErrMsg("用户禁止登陆")
	}
	// 创建JWT
	var jwt service.JWT
	rsp := &adminResponse.LoginAdminByPhone{
		AdminId: admin.Id,
		Phone:   admin.Phone,
	}
	rsp.Token, _ = jwt.CreateToken(service.InitClaims(gin.H{"id": admin.Id}, 3*24*3600))
	return response.InitSucc(rsp)
}
