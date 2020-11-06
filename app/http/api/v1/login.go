package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"ice/app/model/request/api"
	"ice/helper"
	"ice/utils/response"
)

func LoginByPhone(c *gin.Context)  {
	var r api.LoginUserPhone
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	}
	if err := gvalid.CheckStruct(r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	cacheCode, _ := helper.Cache("loginPhoneCode" + r.Phone, "", 0)
	if r.Code != cacheCode {
		response.Ret(response.InitErrCode(response.SmsVerifyFail), c)
	}
	//m := &entity.User{
	//	Phone: r.Phone,
	//}

}
