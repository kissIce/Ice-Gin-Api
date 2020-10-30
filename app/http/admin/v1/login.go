package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	adminLogic "ice/app/logic/admin"
	adminRequest "ice/app/model/request/admin"
	"ice/utils/response"
)

func LoginByPhone(c *gin.Context) {
	var L adminRequest.LoginAdmin
	if err := c.ShouldBind(&L); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(L, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	adminLogic.LoginByPhone(&L)
}
