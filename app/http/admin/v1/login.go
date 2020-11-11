package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	adminLogic "ice/app/logic/admin"
	"ice/app/model/entity"
	adminRequest "ice/app/model/request/admin"
	"ice/app/service"
	"ice/global"
	"ice/utils/response"
)

func LoginByPhone(c *gin.Context) {
	//var r adminRequest.LoginAdmin
	//if err := c.ShouldBind(&r); err != nil {
	//	response.Ret(response.InitErrMsg(err.Error()), c)
	//} else if err := gvalid.CheckStruct(r, nil); err != nil {
	//	response.Ret(response.InitErrMsg(err.FirstString()), c)
	//}
	//adminLogic.LoginByPhone()
}

func RegAdmin(c *gin.Context)  {
	var r adminRequest.RegisterAdmin
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	}
	if err := gvalid.CheckStruct(r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	m := &entity.Admin{
		Username:      r.Username,
		Realname:      r.Realname,
		Password:      r.Password,
		Avatar:        r.Avatar,
		Phone:         r.Phone,
		Status:        0,
	}
	err, id := adminLogic.RegAdmin(m)
	if err != nil {
		response.Ret(response.InitErr(), c)
	} else {
		response.Ret(response.InitSucc(gin.H{"id": id}), c)
	}
}

func DelAdmin(c *gin.Context)  {
	var admin entity.Admin
	global.IceDb.Where("id = ?", 1).Delete(&admin)
}

func SelAdmin(c *gin.Context)  {
	var admin []entity.Admin
	global.IceDb.Find(&admin)
	response.Ret(response.InitSucc(gin.H{"id": admin}), c)
}

func Captcha(c *gin.Context)  {
	id, img, err := service.Captcha()
	if err != nil {
		response.Ret(response.InitErrCode(response.CaptchaError), c)
	}
	response.Ret(response.InitSucc(gin.H{"id": id, "img": img}), c)
}