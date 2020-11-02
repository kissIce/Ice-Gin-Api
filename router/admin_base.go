package router

import (
	"github.com/gin-gonic/gin"
	adminV1 "ice/app/http/admin/v1"
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
		Router.Use(middleware.AdminAuth()).Use(middleware.RbacHandler()).GET("captcha",adminV1.Captcha)

	}
	//Router.GET("jwt", func(c *gin.Context) {
	//	var JWT helper.JWT
	//	token, err := JWT.NewJWT().CreateToken(helper.InitClaims(111))
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	response.Ret(response.InitSucc(token), c)
	//})
	//Router.GET("jwtverify", func(c *gin.Context) {
	//	token := c.Query("token")
	//	var JWT helper.JWT
	//	claims, err := JWT.NewJWT().VerifyToken(token)
	//	if err != nil {
	//
	//	}
	//	response.Ret(response.InitSucc(claims), c)
	//})
}
