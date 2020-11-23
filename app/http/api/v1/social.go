package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	apiRequest "ice/app/model/request/api"
	"ice/utils/response"
)

func HotTzone(c *gin.Context)  {
	var r apiRequest.HotTzone
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	}
	if err := gvalid.CheckStruct(r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}

}
