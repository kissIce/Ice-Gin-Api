package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ice/utils/response"
	"net/url"
	"strconv"
	"time"
)

func ParamAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data url.Values
		switch c.Request.Method {
		case "GET", "DELETE":
			data = c.Request.URL.Query()
		case "POST", "PUT", "PATCH":
			data = c.Request.PostForm
		}
		code := checkParam(data)
		if code != 0 {
			response.Ret(response.InitErrCode(code), c)
			return
		}
		c.Next()
	}
}

func checkParam(data url.Values) int {
	sign := data.Get("sign")
	timestamp := data.Get("timestamp")
	if sign == "" {
		return response.ParamSignMiss
	}
	if len(sign) < 32 {
		return response.ParamSignLenErr
	}
	if timestamp == "" {
		return response.ParamTimeMiss
	}
	reqTime, _ := strconv.ParseInt(timestamp, 10, 64)
	if reqTime > time.Now().Unix() {
		return response.ParamTimeOutUse
	}
	if sign != createSign(data) {
		return response.ParamSignNotEq
	}
	return 0
}

func createSign(data url.Values) string {
	delete(data, "sign")  // sign不参与签名
	sign := data.Encode() // url包里面自带排序
	fmt.Println(sign)
	return ""
}
