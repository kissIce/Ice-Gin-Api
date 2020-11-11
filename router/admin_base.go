package router

import (
	"github.com/gin-gonic/gin"
	adminV1 "ice/app/http/admin/v1"
	apiV1 "ice/app/http/api/v1"
	"ice/app/middleware"
	adminRequest "ice/app/model/request/admin"
	"ice/app/service"
	"ice/utils/response"
)

func InitAdminBase(r *gin.RouterGroup) {
	Router := r.Group("base")
	{
		Router.GET("jwt", func(c *gin.Context) {
			j := service.NewJWT()
			token, _ := j.CreateToken(service.InitClaims(adminRequest.AdminJwt{
				ID:       1,
				UserName: "Ice",
				RealName: "Guo",
				RoleId:   "11",
			}, 100))
			response.Ret(response.InitSucc(token), c)
		})
		Router.POST("login", apiV1.LoginByPhone)
		Router.Use(middleware.AdminAuth()).Use(middleware.RbacHandler()).GET("captcha",adminV1.Captcha)
		//Router.Use(middleware.AdminAuth()).Use(middleware.RbacHandler()).GET("captcha/:id",adminV1.Captcha)
		Router.POST("reg_admin", adminV1.RegAdmin)
	}
}
