package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ice/app/model/data"
	"ice/app/model/entity"
	adminRequest "ice/app/model/request/admin"
	adminResponse "ice/app/model/response/admin"
	"ice/app/service"
	"ice/utils/response"
)

func LoginByPhone(r *adminRequest.LoginAdmin) *response.Resp {
	admin, err := data.GetAdminByPhone(r.Phone, []string{"id", "password", "avatar", "realname", "username", "status"})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return response.InitErrCode(response.DbError)
		}
		return response.InitErrMsg("用户不存在")
	}
	if admin.Password != r.Password {
		return response.InitErrMsg("密码错误")
	}
	if admin.Status != 1 {
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

/**
 * 获取管理员列表
 */
func GetAdminList() *response.Resp {
	rsp, _ := data.GetAdminList()
	return response.InitSucc(rsp)
}

/**
 * 新增管理员
 */
func AddAdmin(r *adminRequest.AddAdmin) *response.Resp {
	admin := &entity.Admin{
		Username: r.Username,
		Realname: r.Realname,
		Password: r.Password,
		Avatar:   r.Avatar,
		Phone:    r.Phone,
		Status:   r.Status,
		RoleId:   r.RoleId,
	}
	err := data.AddAdmin(admin)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(admin)
}

/**
 * 编辑管理员
 */
func EditAdmin(r *adminRequest.EditAdmin) *response.Resp {
	admin := &entity.Admin{
		Id:       r.Id,
		Username: r.Username,
		Realname: r.Realname,
		Password: r.Password,
		Avatar:   r.Avatar,
		Phone:    r.Phone,
		Status:   r.Status,
		RoleId:   r.RoleId,
	}
	err := data.EditAdmin(admin)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(admin)
}

/**
 * 删除管理员
 */
func DelAdmin(r *adminRequest.IdReq) *response.Resp {
	admin := &entity.Admin{
		Id: r.Id,
	}
	err := data.DelAdmin(admin)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(true)
}
