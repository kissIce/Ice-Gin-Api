package core

import (
	"github.com/gin-gonic/gin"
	"ice/app/middleware"
	"ice/global"
	"ice/router"
	"ice/utils/response"
)

func initRouter() (Router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.New()
	Router.Use(middleware.Cors()).Use(middleware.RequestId()).Use(middleware.GinRecovery(false))
	global.IceLog.Info("Use Recovery Middleware")
	group := Router.Group("")
	router.InitAdminBase(group) // 初始化后台基础路由
	// 注册没命中路由
	Router.NoRoute(func(c *gin.Context) {
		response.Ret(response.InitErrCode(response.RouteNofound), c)
	})
	// 注册方法不允许
	Router.NoMethod(func(c *gin.Context) {
		response.Ret(response.InitErrCode(response.MethodNoAllow), c)
	})
	return Router
}
