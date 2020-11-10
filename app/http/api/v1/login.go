package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	apiLogic "ice/app/logic/api"
	"ice/app/model/entity"
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
	m := &entity.User{
		Phone: r.Phone,
	}
	resp := apiLogic.LoginByPhone(m)
	response.Ret(resp, c)
}
