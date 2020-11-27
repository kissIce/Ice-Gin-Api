package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	adminLogic "ice/app/logic/admin"
	adminRequest "ice/app/model/request/admin"
	"ice/utils/response"
)

func GetAdminList(c *gin.Context) {
	resp := adminLogic.GetAdminList()
	response.Ret(resp, c)
}

func AddAdmin(c *gin.Context) {
	var r adminRequest.AddAdmin
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(&r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	resp := adminLogic.AddAdmin(&r)
	response.Ret(resp, c)
}

func EditAdmin(c *gin.Context) {
	var r adminRequest.EditAdmin
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(&r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	resp := adminLogic.EditAdmin(&r)
	response.Ret(resp, c)
}

func DelAdmin(c *gin.Context) {
	var r adminRequest.IdReq
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(&r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	resp := adminLogic.DelAdmin(&r)
	response.Ret(resp, c)
}
