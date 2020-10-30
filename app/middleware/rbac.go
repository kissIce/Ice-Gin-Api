package middleware

import (
	"github.com/gin-gonic/gin"
)

func RbacHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//claims, _ :=c.Get("User")
		//waitUse := claims.(entity.Admin)
		//// 获取请求的URI
		//obj := c.Request.URL.RequestURI()
		//// 获取请求方法
		//act := c.Request.Method
		//// 获取用户的角色
		//sub := waitUse.AuthorityId
		//e := service.Casbin()
		//// 判断策略中是否存在
		//success, _ := e.Enforce(sub, obj, act)
		//if global.IceConfig.System.Env == "develop" || success {
		//	c.Next()
		//} else {
		//	response.Ret(response.InitErr(), c)
		//	c.Abort()
		//	return
		//}
	}
}

