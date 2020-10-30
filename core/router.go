package core

import (
	"github.com/gin-gonic/gin"
	"ice/app/middleware"
	"ice/global"
	"ice/router"
	"ice/utils/response"
)

func initRouter() (Router *gin.Engine) {
	Router = gin.New()
	Router.Use(middleware.Cors()).Use(middleware.GinRecovery(true))
	global.IceLog.Info("use middleware logger")
	group := Router.Group("")
	// 注册没命中路由
	Router.NoRoute(func(c *gin.Context) {
		response.Ret(response.InitErrCode(response.RouteNofound), c)
	})
	// 注册方法不允许
	Router.NoMethod(func(c *gin.Context) {
		response.Ret(response.InitErrCode(response.MethodNoAllow), c)
	})
	router.InitBase(group)
	return Router
}
