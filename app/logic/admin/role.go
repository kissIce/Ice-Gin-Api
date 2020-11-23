package admin

import (
	"ice/app/model/data"
	"ice/app/model/entity"
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
 * 获取角色列表
 */
func GetRoleMenu(roleId uint64) *response.Resp {
	role, _ := data.GetRoleMenu(roleId)
	return response.InitSucc(role)
}

/**
 * 新增角色
 */
func AddRole(role *entity.Role) *response.Resp {
	err := data.AddRole(role)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(role)
}

/**
 * 编辑角色
 */
func EditRole(role *entity.Role) *response.Resp {
	err := data.EditRole(role)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(role)
}

/**
 * 删除角色
 */
func DelRole(role *entity.Role) *response.Resp {
	err := data.DelRole(role)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(true)
}
