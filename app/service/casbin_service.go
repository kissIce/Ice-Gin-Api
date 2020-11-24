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
	// 自定义匹配规则，对应config里面的 matchers
	e.AddFunction("ParamsMatch", ParamsMatch)
	_ = e.LoadPolicy()
	return e
}

/**
 * 自定义匹配
 */
func ParamsMatch(args ...interface{}) (interface{}, error) {
	arg1 := args[0].(string)
	arg2 := args[1].(string)
	key1 := strings.Split(arg1, "?")[0]
	return util.KeyMatch2(key1, arg2), nil
}

func UpdateCasbin(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.IceDb.Table(global.IceConfig.Casbin.TableName).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}
