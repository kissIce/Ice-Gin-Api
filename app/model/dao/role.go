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

func GetRoleMenu(roleId uint64) (menu []entity.Menu, err error) {
	sql := "select m.id,m.pid,m.title,m.`name`,m.icon,m.path,m.component,m.keep_alive,m.hidden,m.sort,m.`default` from role r join auth_menu am on r.id = am.role_id join menu m on m.id = am.menu_id where r.id = ?"
	err = global.IceDb.Raw(sql, roleId).Scan(&menu).Error
	return menu, err
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