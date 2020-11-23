package v1

import (
	"github.com/gin-gonic/gin"
	adminLogic "ice/app/logic/admin"
	"ice/utils/response"
)

func GetRoleList(c *gin.Context) {
	resp := adminLogic.GetRoleList()
	response.Ret(resp, c)
}

func GetRoleMenu(c *gin.Context) {
	resp := adminLogic.GetRoleMenu(1)
	response.Ret(resp, c)
}

func AddRole(c *gin.Context)  {
	
}

func EditRole(c *gin.Context)  {
	
}

func DelRole(c *gin.Context)  {
	
}