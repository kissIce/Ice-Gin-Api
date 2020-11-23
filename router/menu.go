package router

import (
	"github.com/gin-gonic/gin"
	adminV1 "ice/app/http/admin/v1"
)

func InitMenu(r *gin.RouterGroup)  {
	//Router := r.Group("/admin.v1/menu").Use(middleware.AdminAuth()).Use(middleware.RbacHandler())
	Router := r.Group("/admin.v1/menu")
	{
		Router.GET("menu",adminV1.GetMenuList)
		Router.POST("menu",adminV1.AddMenu)
		Router.PUT("menu",adminV1.EditMenu)
		Router.DELETE("menu",adminV1.DelMenu)
	}
}
