package data

import (
	"ice/app/model/dao"
	"ice/app/model/entity"
)

func GetMenuTree() (menuList []entity.Menu) {
	treeMenu, _ := dao.GetMenuTree()
	menuList = treeMenu[0]
	for i := 0; i < len(menuList); i++ {
		getMenuChildList(&menuList[i], treeMenu)
	}
	return menuList
}

func getMenuChildList(menu *entity.Menu, treeMap map[uint64][]entity.Menu) {
	menu.Children = treeMap[menu.Id]
	for i := 0; i < len(menu.Children); i++ {
		getMenuChildList(&menu.Children[i], treeMap)
	}
}

func AddMenu(menu *entity.Menu) error {
	return dao.AddMenu(menu)
}

func EditMenu(menu *entity.Menu) error {
	return dao.EditMenu(menu)
}

func DelMenu(menu *entity.Menu) error {
	return dao.DelMenu(menu)
}