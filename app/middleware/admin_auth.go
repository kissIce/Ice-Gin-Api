package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	adminRequest "ice/app/model/request/admin"
	"ice/app/service"
	"ice/global"
	"ice/utils/response"
	"strconv"
	"time"
)



func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("auth-token")
		if token == "" {
			response.Ret(response.InitErrCode(response.TokenMiss), c)
			return
		}
		if ret, _ := global.IceRedis.HExists("BlackList", token).Result(); ret {
			response.Ret(response.InitErrCode(response.TokenBlack), c)
			return
		}
		j := service.NewJWT()
		claims, errCode := j.VerifyToken(token)
		if errCode != 0 {
			response.Ret(response.InitErrCode(errCode), c)
			return
		}
		// 如果过期了，则判断是否在缓冲时间内
		if claims.ExpiresAt - time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + 60*60*24*7
			newToken, _ := j.CreateToken(claims)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
			if global.IceConfig.System.Singlelogin {

			}
		}
		var admin adminRequest.AdminJwt
		_ = mapstructure.Decode(claims.Data, &admin)
		fmt.Println(claims.Data)
		c.Set("admin", admin)
		c.Next()
	}
}