package data

import (
	"errors"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"ice/app/model/dao"
	"ice/app/model/entity"
	"ice/global"
	"ice/helper"
	"math/rand"
	"strconv"
	"time"
)

const (
	cacheExpire = 7 * 3600 * 24
	userCacheKey = "userInfo:"
	phone2UidKey = "userPhone:"
)

func GetUserByPhone(phone string, field []string) (ret map[string]interface{}, err error) {
	var uid string
	uid, err = global.IceRedis.Get(phone2UidKey + phone).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			ret, err = cacheUser(map[string]interface{}{"phone": phone}, append(field, "id"))
			if err != nil {
				global.IceRedis.Set(phone2UidKey+phone, 0, time.Duration(rand.Intn(20)+20)*time.Second)
			} else {
				global.IceRedis.Set(phone2UidKey+phone, ret["id"], time.Duration(cacheExpire)*time.Second)
			}
		}
	} else {
		uid, _ := strconv.ParseUint(uid, 10, 64)
		if uid == 0 {
			err = gorm.ErrRecordNotFound
		} else {
			ret, err = GetUserById(uid, field)
		}
	}
	return
}

/**
 * 根据用户id从缓存查找用户数据
 */
func GetUserById(uid uint64, field []string) (ret map[string]interface{}, err error) {
	redisCache := global.IceRedis.HMGet(userCacheKey+strconv.FormatUint(uid, 10), field...).Val()
	if redisCache[0] == nil {
		ret, err = cacheUser(map[string]interface{}{"id": uid}, field)
		if err != nil {
			global.IceRedis.HSet(userCacheKey+strconv.FormatUint(uid, 10), "ID", 0)
			global.IceRedis.Expire(userCacheKey+strconv.FormatUint(uid, 10), time.Duration(rand.Intn(20)+20)*time.Second)
		}
	} else {
		ret = helper.Slice2Map(field, redisCache)
	}
	return
}

/**
 * 创建用户并删除请求缓存
 */
func CreateUser(u *entity.User) (err error) {
	err = dao.CreateUser(u)
	if err == nil {
		global.IceRedis.Del(phone2UidKey+u.Phone)
	}
	return
}

/**
 * 获取并缓存用户信息返回指定字段
 */
func cacheUser(where map[string]interface{}, field []string) (map[string]interface{}, error) {
	res, err := dao.GetUserByWhere(where)
	if err == nil {
		global.IceRedis.HMSet(userCacheKey+strconv.FormatUint(res["id"].(uint64), 10), res)
	}
	ret := make(map[string]interface{}, len(field))
	for _, v := range field {
		ret[v] = helper.MakeVal2Str(res[v])
	}
	return ret, err
}
