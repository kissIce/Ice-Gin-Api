package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	adminLogic "ice/app/logic/admin"
	"ice/app/model/entity"
	adminRequest "ice/app/model/request/admin"
	"ice/app/service"
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
	} else if err := gvalid.CheckStruct(r, nil); err != nil {
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
		response.Ret(response.InitSucc(map[string]uint64{"id": id}), c)
	}
}

func Captcha(c *gin.Context)  {
	id, img, err := service.Captcha()
	if err != nil {
		response.Ret(response.InitErrCode(response.CaptchaError), c)
	}
	response.Ret(response.InitSucc(&map[string]string{"id": id, "img": img}), c)
}