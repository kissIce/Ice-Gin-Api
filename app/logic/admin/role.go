package admin

import (
	"ice/app/model/data"
	"ice/app/model/entity"
	adminRequest "ice/app/model/request/admin"
	"ice/utils/response"
)

/**
 * 获取角色列表
 */
func GetRoleList() *response.Resp {
	rsp := data.GetRoleTree()
	return response.InitSucc(rsp)
}

/**
 * 获取角色菜单
 */
func GetRoleMenu(roleId uint64) *response.Resp {
	role, _ := data.GetRoleMenu(roleId)
	return response.InitSucc(role)
}

/**
 * 新增角色
 */
func AddRole(r *adminRequest.AddRole) *response.Resp {
	role := &entity.Role{
		Pid:  r.Pid,
		Name: r.Name,
	}
	err := data.AddRole(role)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(role)
}

/**
 * 编辑角色
 */
func EditRole(r *adminRequest.EditRole) *response.Resp {
	role := &entity.Role{
		Id:   r.Id,
		Name: r.Name,
	}
	err := data.EditRole(role)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(role)
}

/**
 * 删除角色
 */
func DelRole(r *adminRequest.IdReq) *response.Resp {
	role := &entity.Role{
		Id:   r.Id,
	}
	err := data.DelRole(role)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(true)
}
