package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	adminLogic "ice/app/logic/admin"
	adminRequest "ice/app/model/request/admin"
	"ice/utils/response"
)


func GetMenuList(c *gin.Context) {
	resp := adminLogic.GetMenuList()
	response.Ret(resp, c)
}

func AddMenu(c *gin.Context) {
	var r adminRequest.AddMenu
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(&r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	resp := adminLogic.AddMenu(&r)
	response.Ret(resp, c)
}

func EditMenu(c *gin.Context) {
	var r adminRequest.EditMenu
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(&r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	resp := adminLogic.EditMenu(&r)
	response.Ret(resp, c)
}

func DelMenu(c *gin.Context) {
	var r adminRequest.IdReq
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(&r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	resp := adminLogic.DelMenu(&r)
	response.Ret(resp, c)
}
