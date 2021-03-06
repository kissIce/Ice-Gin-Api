package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	apiLogic "ice/app/logic/api"
	apiRequest "ice/app/model/request/api"
	"ice/utils/response"
)

func LoginByPhone(c *gin.Context) {
	var r apiRequest.LoginUserPhone
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	}
	if err := gvalid.CheckStruct(r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}

	resp := apiLogic.LoginByPhone(r.Phone)
	response.Ret(resp, c)
}
