package helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"ice/global"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
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
func PwdHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
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
func Cache(name string, val interface{}, timeout time.Duration) (interface{}, error) {
	if val == "" {
		return global.IceRedis.Get(name).Result()
	} else if val == nil {
		return global.IceRedis.Del(name).Result()
	}
	return global.IceRedis.Set(name, val, timeout).Result()
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
				global.IceLog.Error("create directory"+v, zap.Any(" error:", err))
			}
		}
	}
	return err
}

func RandInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	var data = make(map[string]interface{}, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fieldInfo := t.Field(i)
		tag := fieldInfo.Tag.Get("json")
		if tag == "" {
			tag = strings.ToLower(fieldInfo.Name)
		}
		if tag == "-" {
			continue
		}
		data[tag] = v.Field(i).Interface()
	}
	return data
}

func Map2Struct(data map[string]string, s interface{}) error {
	v := reflect.ValueOf(s).Elem()
	if !v.CanAddr() {
		return fmt.Errorf("must be a pointer")
	}
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag.Get("json")
		if tag == "" {
			tag = strings.ToLower(fieldInfo.Name)
		}
		if tag == "-" {
			continue
		}
		if value, ok := data[tag]; ok {
			kind := v.FieldByName(fieldInfo.Name).Kind()
			switch {
			case kind == reflect.Int64, kind == reflect.Int, kind == reflect.Int8,
				kind == reflect.Int16, kind == reflect.Int32:
				{
					val, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						panic(err)
					}
					fi := v.FieldByName(fieldInfo.Name)
					if !fi.CanSet() {
						return fmt.Errorf("can not set value of:%s", fieldInfo.Name)
					}
					fi.SetInt(val)
				}
			case kind == reflect.Uint64, kind == reflect.Uint, kind == reflect.Uint8,
				kind == reflect.Uint16, kind == reflect.Uint32:
				{
					val, err := strconv.ParseUint(value, 10, 64)
					if err != nil {
						panic(err)
					}
					fi := v.FieldByName(fieldInfo.Name)
					if !fi.CanSet() {
						return fmt.Errorf("can not set value of:%s", fieldInfo.Name)
					}
					fi.SetUint(val)
				}
			case kind == reflect.String:
				{
					fi := v.FieldByName(fieldInfo.Name)
					if !fi.CanSet() {
						return fmt.Errorf("can not set value of:%s", fieldInfo.Name)
					}
					fi.SetString(value)
				}
			case kind == reflect.Bool:
				{
					val, err := strconv.ParseBool(value)
					if err != nil {
						panic(err)
					}
					fi := v.FieldByName(fieldInfo.Name)
					if !fi.CanSet() {
						return fmt.Errorf("can not set value of:%s", fieldInfo.Name)
					}
					fi.SetBool(val)

				}
			case kind == reflect.Float32, kind == reflect.Float64:
				{
					val, err := strconv.ParseFloat(value, 64)
					if err != nil {
						panic(err)
					}
					fi := v.FieldByName(fieldInfo.Name)
					if !fi.CanSet() {
						return fmt.Errorf("can not set value of:%s", fieldInfo.Name)
					}
					fi.SetFloat(val)
				}
			default:
				continue
				// return fmt.Errorf("暂不支持%s类型", kind)
			}
		}
	}
	return nil
}

func Slice2Struct(key []string, val []interface{}, s interface{}) error {
	sv := reflect.ValueOf(s).Elem()
	if !sv.CanAddr() {
		return fmt.Errorf("must be a pointer")
	}
	for k, v := range val {
		fmt.Println(Ucfirst(Case2Camel(key[k])))
		fieldInfo, _ := sv.Type().FieldByName(Ucfirst(Case2Camel(key[k])))
		fi := sv.FieldByName(fieldInfo.Name)
		if !fi.CanSet() {
			return fmt.Errorf("can not set value of:%s", fieldInfo.Name)
		}
		kind := fi.Kind()
		switch {
		case kind == reflect.Int64, kind == reflect.Int, kind == reflect.Int8,
			kind == reflect.Int16, kind == reflect.Int32:
			{
				val, err := strconv.ParseInt(v.(string), 10, 64)
				if err != nil {
					panic(err)
				}
				fi := sv.FieldByName(fieldInfo.Name)
				if !fi.CanSet() {
					return fmt.Errorf("can not set value of:%s", fieldInfo.Name)
				}
				fi.SetInt(val)
			}
		case kind == reflect.Uint64, kind == reflect.Uint, kind == reflect.Uint8,
			kind == reflect.Uint16, kind == reflect.Uint32:
			{
				val, err := strconv.ParseUint(v.(string), 10, 64)
				if err != nil {
					panic(err)
				}
				fi := sv.FieldByName(fieldInfo.Name)
				if !fi.CanSet() {
					return fmt.Errorf("can not set value of:%s", fieldInfo.Name)
				}
				fi.SetUint(val)
			}
		case kind == reflect.String:
			{
				fi := sv.FieldByName(fieldInfo.Name)
				if !fi.CanSet() {
					return fmt.Errorf("can not set value of:%s", fieldInfo.Name)
				}
				fi.SetString(v.(string))
			}
		case kind == reflect.Bool:
			{
				val, err := strconv.ParseBool(v.(string))
				if err != nil {
					panic(err)
				}
				fi := sv.FieldByName(fieldInfo.Name)
				if !fi.CanSet() {
					return fmt.Errorf("can not set value of:%s", fieldInfo.Name)
				}
				fi.SetBool(val)

			}
		case kind == reflect.Float32, kind == reflect.Float64:
			{
				val, err := strconv.ParseFloat(v.(string), 64)
				if err != nil {
					panic(err)
				}
				fi := sv.FieldByName(fieldInfo.Name)
				if !fi.CanSet() {
					return fmt.Errorf("can not set value of:%s", fieldInfo.Name)
				}
				fi.SetFloat(val)
			}
		default:
			continue
			// return fmt.Errorf("暂不支持%s类型", kind)
		}
		//fi.se
	}
	return nil
}

func Slice2Map(key []string, val []interface{}) map[string]interface{} {
	var s = make(map[string]interface{}, len(val))
	for k, v := range val {
		s[key[k]] = v
	}
	return s
}

func MakeVal2Str(val interface{}) (ret string) {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int:
		return strconv.Itoa(val.(int))
	case reflect.Int8:
		return strconv.Itoa(int(val.(int8)))
	case reflect.Int16:
		return strconv.Itoa(int(val.(int16)))
	case reflect.Int32:
		return strconv.Itoa(int(val.(int32)))
	case reflect.Int64:
		return strconv.Itoa(int(val.(int64)))
	case reflect.Uint:
		return strconv.Itoa(int(val.(uint)))
	case reflect.Uint8:
		return strconv.Itoa(int(val.(uint8)))
	case reflect.Uint16:
		return strconv.Itoa(int(val.(uint16)))
	case reflect.Uint32:
		return strconv.Itoa(int(val.(uint32)))
	case reflect.Uint64:
		return strconv.FormatUint(val.(uint64), 10)
	case reflect.Bool:
		return strconv.FormatBool(val.(bool))
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.(float64), 'E', -1, 64)
	default:
		return val.(string)
	}
}

func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 首字母大写
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}
