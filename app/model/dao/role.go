package dao

import (
	"ice/app/model/entity"
	"ice/global"
	"time"
)

func GetRoleTree() (treeMap map[uint64][]entity.Role, err error) {
	var allRoles []entity.Role
	treeMap = make(map[uint64][]entity.Role)
	err = global.IceDb.Select([]string{"id","pid","name"}).Where("deleted_at = ?", 0).Find(&allRoles).Error
	if err == nil {
		for _, v := range allRoles {
			treeMap[v.Pid] = append(treeMap[v.Pid], v)
		}
	}
	return treeMap, err
}

func GetRoleMenu(roleId uint64) (*entity.Role, error) {
	var role entity.Role
	err := global.IceDb.Where("id = ?", roleId).Preload("menus").First(&role).Error
	return &role, err
}

func AddRole(role *entity.Role) error {
	return global.IceDb.Create(role).Error
}

func EditRole(role *entity.Role) error {
	return global.IceDb.Updates(role).Error
}

func DelRole(role *entity.Role) error {
	return global.IceDb.Model(role).Where("id = ?", role.Id).Update("deleted_at", time.Now().Unix()).Error
}