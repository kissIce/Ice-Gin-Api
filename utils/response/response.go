package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func InitSucc(data interface{}) *Resp {
	return &Resp{
		Code: ApiSuccess,
		Msg:  "请求成功",
		Data: data,
	}
}

func InitErr() *Resp {
	return &Resp{
		Code: ApiError,
		Msg:  "请求失败",
	}
}

func Ret(resp *Resp, c *gin.Context)  {
	c.JSON(http.StatusOK, &resp)
}