package dao

import (
	"ice/app/model/entity"
	"ice/global"
	"time"
)

func GetMenuTree() (treeMap map[uint64][]entity.Menu, err error) {
	var allMenus []entity.Menu
	treeMap = make(map[uint64][]entity.Menu)
	err = global.IceDb.Select([]string{"id","pid","title","name","icon","path","component","keep_alive","hidden","sort","default"}).Where("deleted_at = ?", 0).Order("sort").Find(&allMenus).Error
	if err == nil {
		for _, v := range allMenus {
			treeMap[v.Pid] = append(treeMap[v.Pid], v)
		}
	}
	return treeMap, err
}

func AddMenu(menu *entity.Menu) error {
	return global.IceDb.Create(menu).Error
}

func EditMenu(menu *entity.Menu) error {
	return global.IceDb.Updates(menu).Error
}

func DelMenu(menu *entity.Menu) error {
	return global.IceDb.Model(menu).Where("id = ?", menu.Id).Update("deleted_at", time.Now().Unix()).Error
}