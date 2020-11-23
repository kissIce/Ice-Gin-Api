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
/**
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `auth_menu` AS select `menu`.`pid` AS `pid`,`menu`.`path` AS `path`,`menu`.`name` AS `name`,`menu`.`hidden` AS `hidden`,`menu`.`component` AS `component`, `menu`.`title`  AS `title`,`menu`.`icon` AS `icon`,`menu`.`sort` AS `sort`,`auth_menu`.`role_id` AS `role_id`,`auth_menu`.`menu_id` AS `menu_id`,`menu`.`keep_alive` AS `keep_alive`,`menu`.`default` AS `default` from (`auth_menu` join `menu` on ((`auth_menu`.`menu_id` = `menu`.`id` and `menu`.`deleted_at` = 0)))

select r.id,m.* from role r join auth_menu am on r.id = am.role_id join menu m on m.id = am.menu_id where r.id = 1
 */
func GetRoleMenu(roleId uint64) (menu []entity.Menu, err error) {
	//err := global.IceDb.Where("id = ?", roleId).Preload("Menus").First(&role).Error
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