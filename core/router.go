package core

import (
	"github.com/gin-gonic/gin"
	"ice/app/middleware"
	"ice/router"
	"ice/utils/response"
)

func initRouter() (Router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.New()
	Router.Use(middleware.GinRecovery(true)).Use(middleware.Cors()).Use(middleware.RequestId()).Use(middleware.ParamCheck())
	group := Router.Group("")
	router.InitAdminBase(group) // 初始化后台基础路由
	router.InitMenu(group)      // 初始化后台菜单路由
	router.InitRole(group)      // 初始化后台角色路由
	// 注册没命中路由
	Router.NoRoute(func(c *gin.Context) {
		response.Ret(response.InitErrCode(response.RouteNotfound), c)
		return
	})
	// 注册方法不允许
	Router.NoMethod(func(c *gin.Context) {
		response.Ret(response.InitErrCode(response.MethodNotAllow), c)
		return
	})
	return Router
}
