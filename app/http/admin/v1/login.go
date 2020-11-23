package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	adminLogic "ice/app/logic/admin"
	adminRequest "ice/app/model/request/admin"
	"ice/utils/response"
)

func LoginByPhone(c *gin.Context) {
	var r adminRequest.LoginAdmin
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	resp := adminLogic.LoginByPhone(&r)
	response.Ret(resp, c)
}