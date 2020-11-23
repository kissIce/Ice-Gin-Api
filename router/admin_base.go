package router

import (
	"github.com/gin-gonic/gin"
	adminV1 "ice/app/http/admin/v1"
)

func InitAdminBase(r *gin.RouterGroup) {
	Router := r.Group("base")
	{
		Router.POST("login",adminV1.LoginByPhone)
	}
}
