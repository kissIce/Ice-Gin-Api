package router

import (
	"github.com/gin-gonic/gin"
	adminV1 "ice/app/http/admin/v1"
)

func InitRole(r *gin.RouterGroup)  {
	//Router := r.Group("/admin.v1/menu").Use(middleware.AdminAuth()).Use(middleware.RbacHandler())
	Router := r.Group("/admin.v1/role")
	{
		Router.GET("role",adminV1.GetRoleList)
		Router.GET("role_menu",adminV1.GetRoleMenu)
		Router.POST("role",adminV1.AddRole)
		Router.PUT("role",adminV1.EditRole)
		Router.DELETE("role",adminV1.DelRole)
	}
}
