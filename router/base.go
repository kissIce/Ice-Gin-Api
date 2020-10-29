package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ice/helper"
	"ice/utils/response"
)

func InitBase(r *gin.RouterGroup) (Router gin.IRoutes) {
	Router = r.Group("base")
	Router.GET("jwt", func(c *gin.Context) {
		var JWT helper.JWT
		token, err := JWT.NewJWT().CreateToken(helper.InitClaims(111))
		if err != nil {
			fmt.Println(err)
		}
		response.Ret(response.InitSucc(token), c)
	})
	Router.GET("jwtverify", func(c *gin.Context) {
		token := c.Query("token")
		var JWT helper.JWT
		claims, err := JWT.NewJWT().VerifyToken(token)
		if err != nil {

		}
		response.Ret(response.InitSucc(claims), c)
	})
	return
}
