package router

import (
	"github.com/gin-gonic/gin"
	adminV1 "ice/app/http/admin/v1"
	apiV1 "ice/app/http/api/v1"
	"ice/app/middleware"
)

func InitAdminBase(r *gin.RouterGroup) {
	Router := r.Group("base")
	{
		Router.POST("login", apiV1.LoginByPhone)
		Router.POST("reg_admin", adminV1.RegAdmin)
		Router.POST("del_admin", adminV1.DelAdmin)
		Router.POST("get_admin", adminV1.SelAdmin)
		Router.Use(middleware.AdminAuth()).Use(middleware.RbacHandler()).GET("captcha",adminV1.Captcha)
		//Router.Use(middleware.AdminAuth()).Use(middleware.RbacHandler()).GET("captcha/:id",adminV1.Captcha)
	}
}
