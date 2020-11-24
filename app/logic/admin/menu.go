package admin

import (
	"ice/app/model/data"
	"ice/app/model/entity"
	adminRequest "ice/app/model/request/admin"
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
func AddMenu(r *adminRequest.AddMenu) *response.Resp {
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
	err := data.AddMenu(menu)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(menu)
}

/**
 * 编辑菜单
 */
func EditMenu(r *adminRequest.EditMenu) *response.Resp {
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
	err := data.EditMenu(menu)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(menu)
}

/**
 * 删除菜单
 */
func DelMenu(r *adminRequest.IdReq) *response.Resp {
	menu := &entity.Menu{
		Id: r.Id,
	}
	err := data.DelMenu(menu)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(true)
}
