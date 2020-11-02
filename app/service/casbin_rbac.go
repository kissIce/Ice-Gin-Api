package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"ice/global"
	"strings"
)

/**
 * casbin实例
 */
func Casbin() *casbin.Enforcer {
	a, _ := gormadapter.NewAdapter("mysql", global.IceConfig.Mysql.Username+":"+global.IceConfig.Mysql.Password+"@tcp("+global.IceConfig.Mysql.Uri+":"+global.IceConfig.Mysql.Port+")/" + global.IceConfig.Mysql.Dbname, global.IceConfig.Mysql.Dbname, global.IceConfig.Casbin.TableName, true)
	e, _ := casbin.NewEnforcer(global.IceConfig.Casbin.ModelPath, a)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = e.LoadPolicy()
	return e
}

// @title    ParamsMatch
// @description   customized rule, 自定义规则函数
// @auth                     （2020/04/05  20:22）
// @param     fullNameKey1    string
// @param     key2            string
// @return                    bool

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// @title    ParamsMatchFunc
// @description   customized function, 自定义规则函数
// @auth                     （2020/04/05  20:22）
// @param     args            ...interface{}
// @return                    interface{}
// @return                    error

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}

