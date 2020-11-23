package dao

import (
	"gorm.io/gorm"
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
	return global.IceDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(menu).Update("deleted_at", time.Now().Unix()).Error; err != nil {
			return err
		}
		if err := tx.Model(menu).Association("Roles").Clear(); err != nil {
			return err
		}
		return nil
	})
}