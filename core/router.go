package core

import (
	"github.com/gin-gonic/gin"
	"ice/app/middleware/common"
	"ice/global"
	"ice/router"
)

func initRouter() (Router *gin.Engine) {
	Router = gin.New()
	Router.Use(common.Cors())
	global.IceLog.Info("use middleware logger")
	group := Router.Group("")
	router.InitBase(group)
	return Router
}
