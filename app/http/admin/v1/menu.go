package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	adminLogic "ice/app/logic/admin"
	"ice/app/model/entity"
	adminRequest "ice/app/model/request/admin"
	"ice/utils/response"
)


func GetMenuList(c *gin.Context) {
	resp := adminLogic.GetMenuList()
	response.Ret(resp, c)
}

func AddMenu(c *gin.Context) {
	var r adminRequest.AddMenu
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	menu := &entity.Menu{
		Pid:       r.Pid,
		Title:     r.Title,
		Name:      r.Name,
		Icon:      r.Icon,
		Path:      r.Path,
		Component: r.Component,
		KeepAlive: r.KeepAlive,
		Hidden:    r.Hidden,
		Sort:      r.Sort,
		Default:   r.Default,
	}
	resp := adminLogic.AddMenu(menu)
	response.Ret(resp, c)
}

func EditMenu(c *gin.Context) {
	var r adminRequest.EditMenu
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	menu := &entity.Menu{
		Id:        r.Id,
		Pid:       r.Pid,
		Title:     r.Title,
		Name:      r.Name,
		Icon:      r.Icon,
		Path:      r.Path,
		Component: r.Component,
		KeepAlive: r.KeepAlive,
		Hidden:    r.Hidden,
		Sort:      r.Sort,
		Default:   r.Default,
	}
	resp := adminLogic.EditMenu(menu)
	response.Ret(resp, c)
}

func DelMenu(c *gin.Context) {
	var r adminRequest.DelMenu
	if err := c.ShouldBind(&r); err != nil {
		response.Ret(response.InitErrMsg(err.Error()), c)
	} else if err := gvalid.CheckStruct(r, nil); err != nil {
		response.Ret(response.InitErrMsg(err.FirstString()), c)
	}
	menu := &entity.Menu{
		Id: r.Id,
	}
	resp := adminLogic.DelMenu(menu)
	response.Ret(resp, c)
}
