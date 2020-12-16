package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	adminLogic "ice/app/logic/admin"
	adminRequest "ice/app/model/request/admin"
	"ice/utils/response"
)

func GetAdCateList(c *gin.Context) {
	resp := adminLogic.GetAdCateList()
	response.Ret(resp, c)
}

func AddAdCate(c *gin.Context) {
	var r adminRequest.AddAdCate
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(&r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	resp := adminLogic.AddAdCate(&r)
	response.Ret(resp, c)
}

func EditAdCate(c *gin.Context) {
	var r adminRequest.EditAdCate
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(&r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	resp := adminLogic.EditAdCate(&r)
	response.Ret(resp, c)
}

func DelAdCate(c *gin.Context) {
	var r adminRequest.IdReq
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(&r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	resp := adminLogic.DelAdCate(&r)
	response.Ret(resp, c)
}
