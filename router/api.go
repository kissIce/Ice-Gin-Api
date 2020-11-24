package router

import (
	"github.com/gin-gonic/gin"
	adminV1 "ice/app/http/admin/v1"
)

func InitApi(r *gin.RouterGroup)  {
	Router := r.Group("/admin.v1/api")
	{
		Router.GET("api",adminV1.GetApiList)
		Router.POST("api",adminV1.AddRole)
		Router.PUT("api",adminV1.EditRole)
		Router.DELETE("api",adminV1.DelRole)
	}
}

