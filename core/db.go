package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"ice/global"
	"os"
)

var err error

func initDB() {
	m := global.IceConfig.Mysql
	sqlUri := m.Username + ":" + m.Password + "@tcp(" + m.Uri + ":" + m.Port + ")/" + m.Dbname + "?charset=" + m.Charset
	mysqlCfg := mysql.Config{
		DSN:                       sqlUri,
		DefaultStringSize:         120,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	fmt.Println(sqlUri)
	gormCfg := config(m.LogMode)
	if global.IceDb, err = gorm.Open(mysql.New(mysqlCfg), gormCfg); err != nil {
		fmt.Printf("mysql 链接失败，错误原因：%s\n", err)
		os.Exit(0)
	} else {
		db, _ := global.IceDb.DB()
		db.SetMaxIdleConns(m.MaxIdleConns)
		db.SetMaxOpenConns(m.MaxOpenConns)
	}
}

func config(b bool) (cfg *gorm.Config) {
	if b {
		cfg = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		cfg = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
	return
}
