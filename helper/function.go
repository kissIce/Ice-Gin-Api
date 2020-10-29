package helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"ice/global"
	"math"
	"os"
	"strings"
	"time"
)

/**
 * 加盐MD5
 */
func SafeMd5(str string) string {
	h := md5.New()
	h.Write([]byte(global.IceConfig.Md5Salt + str))
	return hex.EncodeToString(h.Sum(nil))
}

/**
 * 密码加密
 */
func PwdHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

/**
 * 密码验证
 */
func PwdVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/**
 * 快速缓存
 */
func Cache(name string, val interface{}, timeout time.Duration) interface{} {
	if val == "" {
		return global.IceRedis.Get(name)
	} else if val == nil {
		return global.IceRedis.Del(name)
	}
	return global.IceRedis.Set(name, val, timeout)
}

/**
 * 根据id 生成唯一邀请码
 */
func InviteCode(uid uint64) string {
	num := uid
	var code string
	for {
		if num <= 0 {
			break
		}
		mod := num % 35
		fmt.Println(mod)
		num = (num - mod) / 35
		code = string(global.IceConfig.InviteStr[mod]) + code
	}
	if len(code) < 3 {
		code = fmt.Sprintf("%04s", code)
	}
	return code
}

/**
 * 根据邀请码解码用户id
 */
func InviteDecode(str string) uint64 {
	if idx := strings.LastIndex(str, "0"); idx != -1 {
		str = str[idx+1:]
	}
	str = ReverseString(str)
	var id uint64 = 0
	for i := 0; i < len(str); i++ {
		id += uint64(strings.Index(global.IceConfig.InviteStr, string(str[i]))) * uint64(math.Pow(35, float64(i)))
	}
	return id
}

/**
 * 翻转字符串
 */
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

/**
 * 文件目录是否存在
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
 * 批量创建文件夹
 */
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.IceLog.Debug("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.IceLog.Error("create directory"+ v, zap.Any(" error:", err))
			}
		}
	}
	return err
}
