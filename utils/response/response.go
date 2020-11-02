package response

import (
	"github.com/gin-gonic/gin"
	"ice/global"
	"net/http"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func InitSucc(data interface{}) *Resp {
	return &Resp{
		Code: ApiSuccess,
		Msg:  resultMsg(ApiSuccess),
		Data: data,
	}
}

func InitErr() *Resp {
	return &Resp{
		Code: ApiError,
		Msg:  resultMsg(ApiError),
	}
}

func InitErrMsg(str string) *Resp {
	return &Resp{
		Code: ApiError,
		Msg:  str,
	}
}

func InitErrCode(code int) *Resp {
	return &Resp{
		Code: code,
		Msg:  resultMsg(code),
	}
}

func Ret(resp *Resp, c *gin.Context) {
	c.JSON(http.StatusOK, &resp)
	c.Abort()
}

func resultMsg(code int) string {
	if _, ok := RetMsg[code]; ok {
		return RetMsg[code][global.IceConfig.System.Language]
	}
	if code == ApiSuccess {
		return RetMsg[ApiSuccess][global.IceConfig.System.Language]
	} else {
		return RetMsg[ApiError][global.IceConfig.System.Language]
	}
}
