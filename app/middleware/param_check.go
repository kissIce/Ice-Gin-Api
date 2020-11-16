package middleware

import (
	"github.com/gin-gonic/gin"
	"ice/helper"
	"ice/utils/response"
	"net/url"
	"strconv"
	"time"
)

func ParamCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data url.Values
		switch c.Request.Method {
		case "GET", "DELETE":
			data = c.Request.URL.Query()
		case "POST", "PUT", "PATCH":
			data = c.Request.PostForm
		}
		// 有请求参数就进行请求校验
		if len(data) > 0 {
			code := checkParam(data)
			if code != 0 {
				response.Ret(response.InitErrCode(code), c)
				return
			}
		}
		c.Next()
	}
}

func checkParam(data url.Values) int {
	sign := data.Get("sign")
	timestamp := data.Get("timestamp")
	nonceStr := data.Get("nonce_str")
	if sign == "" {
		return response.ParamSignMiss
	}
	if len(sign) != 32 {
		return response.ParamSignLenErr
	}
	if timestamp == "" {
		return response.ParamTimeMiss
	}
	if nonceStr == "" {
		return response.ParamNonceStrMiss
	}
	if len(nonceStr) != 32 {
		return response.ParamNonceStrLenErr
	}
	use, _ := helper.Cache("Request:nonceStr"+nonceStr, "", 0) // 判断是否请求过
	if use != "" {
		return response.ParamNonceStrRepeat
	}
	_, _ = helper.Cache("Request:nonceStr"+nonceStr, 1, 600) // 缓存请求
	reqTime, _ := strconv.ParseInt(timestamp, 10, 64)
	if reqTime > time.Now().Unix() || reqTime+600 < time.Now().Unix() { // 请求误差10分钟
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
	return helper.SafeMd5(sign)
}
