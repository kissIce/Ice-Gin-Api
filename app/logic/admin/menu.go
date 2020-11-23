package admin

import (
	"ice/app/model/data"
	"ice/app/model/entity"
	"ice/utils/response"
)

/**
 * 获取菜单列表
 */
func GetMenuList() *response.Resp {
	rsp := data.GetMenuTree()
	return response.InitSucc(rsp)
}

/**
 * 新增菜单
 */
func AddMenu(menu *entity.Menu) *response.Resp {
	err := data.AddMenu(menu)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(menu)
}

/**
 * 编辑菜单
 */
func EditMenu(menu *entity.Menu) *response.Resp {
	err := data.EditMenu(menu)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(menu)
}

/**
 * 删除菜单
 */
func DelMenu(menu *entity.Menu) *response.Resp {
	err := data.DelMenu(menu)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(true)
}
