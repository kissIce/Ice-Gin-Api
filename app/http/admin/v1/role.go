package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	adminLogic "ice/app/logic/admin"
	adminRequest "ice/app/model/request/admin"
	"ice/utils/response"
)

func GetRoleList(c *gin.Context) {
	resp := adminLogic.GetRoleList()
	response.Ret(resp, c)
}

func GetRoleMenu(c *gin.Context) {
	resp := adminLogic.GetRoleMenu(1)
	response.Ret(resp, c)
}

func AddRole(c *gin.Context) {
	var r adminRequest.AddRole
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(&r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	resp := adminLogic.AddRole(&r)
	response.Ret(resp, c)
}

func EditRole(c *gin.Context) {
	var r adminRequest.EditRole
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(&r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	resp := adminLogic.EditRole(&r)
	response.Ret(resp, c)
}

func DelRole(c *gin.Context) {
	var r adminRequest.IdReq
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(&r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	resp := adminLogic.DelRole(&r)
	response.Ret(resp, c)
}

