package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"ice/global"
	"io"
	"os"
	"strings"
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
		//initStruct()
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
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "t_",   // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	}
	return
}


type column struct {
	name string
	data_type string
}

func initStruct()  {
	sql_str := "SELECT `COLUMN_NAME`,`DATA_TYPE`,`TABLE_NAME` FROM `information_schema`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? ORDER BY `TABLE_NAME` ASC,`ORDINAL_POSITION` ASC"
	rows, _ := global.IceDb.Raw(sql_str, global.IceConfig.Mysql.Dbname).Rows()
	defer rows.Close()
	var tableColumns  = make(map[string][]column)
	var tableList []string
	for rows.Next() {
		var column_name,data_type,table string
		err := rows.Scan(&column_name,&data_type,&table)
		if err != nil {
			panic("GetAllTable Scan error:"+err.Error())
		}
		item := column{
			column_name,
			data_type,
		}
		if len(tableColumns[table]) < 1 {
			tableList = append(tableList,table)
		}
		tableColumns[table] = append(tableColumns[table],item)
	}
	for _,table := range tableList {
		structContent := "package entity\n\ntype "+ HumpFormat(table) +" struct {\n "
		var param  []interface{}
		for _,column := range tableColumns[table] {
			//if column.name == "id" || column.name == "created_at" || column.name == "updated_at" || column.name == "deleted_at" {
			//	continue
			//}
			type_str := mysqlType(column.data_type)
			param = append(param,HumpFormat(column.name),type_str,column.name,column.name)
			structContent += "\n"+
				"  %s %s `json:\"%s\" form:\"%s\"`"
		}
		structContent += "\n}"
		fileContent := fmt.Sprintf(structContent,param...)
		file,_ := os.Create("./app/model/entity/" + table + ".go")
		defer file.Close()
		_,err = io.WriteString(file, fileContent)
	}
}

func mysqlType(data_type string) string {
	var res string
	switch strings.ToUpper(data_type) {
	case "TINYINT":
		res = "int8"
	case "SMALLINT":
		res = "int16"
	case "MEDIUMINT","INT","INTEGER":
		res = "int32"
	case "BIGINT":
		res = "int64"
	case "FLOAT","DOUBLE","DECIMAL":
		res = "float64"
	default:
		res = "string"
	}
	return  res
}

//驼峰格式
func HumpFormat(str string) string {
	var res string
	ary := strings.Split(str,"_")
	for _,v := range ary {
		res += Capitalize(v)
	}
	return res
}

//首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}