package data

import (
	"ice/app/model/dao"
	"ice/app/model/entity"
)

func GetRoleTree() (roleList []entity.Role) {
	treeRole, _ := dao.GetRoleTree()
	roleList = treeRole[0]
	for i := 0; i < len(roleList); i++ {
		getRoleChildList(&roleList[i], treeRole)
	}
	return roleList
}

func GetRoleMenu(roleId uint64) ([]entity.Menu, error) {
	return dao.GetRoleMenu(roleId)
}

func getRoleChildList(role *entity.Role, treeMap map[uint64][]entity.Role) {
	role.Children = treeMap[role.Id]
	for i := 0; i < len(role.Children); i++ {
		getRoleChildList(&role.Children[i], treeMap)
	}
}

func AddRole(role *entity.Role) error {
	return dao.AddRole(role)
}

func EditRole(role *entity.Role) error {
	return dao.EditRole(role)
}

func DelRole(role *entity.Role) error {
	return dao.DelRole(role)
}
