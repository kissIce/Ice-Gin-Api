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
	var loglevel logger.LogLevel
	if b {
		loglevel = logger.Info
	} else {
		loglevel = logger.Silent
	}
	cfg = &gorm.Config{
		Logger:                                   logger.Default.LogMode(loglevel),
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	return
}
