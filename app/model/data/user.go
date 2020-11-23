package data

import (
	"errors"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"ice/app/model/dao"
	"ice/app/model/entity"
	"ice/global"
	"ice/helper"
	"strconv"
	"time"
)

const (
	userCacheKey = "userInfo:"
	userPhoneKey = "userPhone:"
)

func GetUserByPhone(phone string, field []string) (ret *entity.User, err error) {
	var uid string
	uid, err = global.IceRedis.Get(userPhoneKey + phone).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			ret, err = cacheUser(map[string]interface{}{"phone": phone})
			if err != nil {
				global.IceRedis.Set(userPhoneKey+phone, 0, time.Duration(helper.RandRangeInt(10, 60))*time.Second)
			} else {
				global.IceRedis.Set(userPhoneKey+phone, ret.Id, time.Duration(7 * 3600 * 24)*time.Second)
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
func GetUserById(uid uint64, field []string) (*entity.User, error) {
	var ret = new(entity.User)
	var err error
	redisCache := global.IceRedis.HMGet(userCacheKey+strconv.FormatUint(uid, 10), field...).Val()
	if redisCache[0] == nil {
		ret, err = cacheUser(map[string]interface{}{"id": uid})
		if err != nil {
			global.IceRedis.HSet(userCacheKey+strconv.FormatUint(uid, 10), "ID", 0)
			global.IceRedis.Expire(userCacheKey+strconv.FormatUint(uid, 10), time.Duration(helper.RandRangeInt(10, 99))*time.Second)
		}
	} else {
		err = helper.Slice2Struct(field, redisCache, ret)
	}
	return ret, err
}

/**
 * 创建用户并删除请求缓存
 */
func CreateUser(u *entity.User) (err error) {
	err = dao.CreateUser(u)
	if err == nil {
		global.IceRedis.Del(userPhoneKey+u.Phone)
	}
	return
}

func UpdateUserById(uid uint64, data map[string]interface{}) error {
	return dao.UpdateUserById(uid, data)
}

/**
 * 获取并缓存用户信息返回指定字段
 */
func cacheUser(where map[string]interface{}) (*entity.User, error) {
	u, err := dao.GetUserByWhere(where)
	if err == nil {
		global.IceRedis.HMSet(userCacheKey+strconv.FormatUint(u.Id, 10), helper.Struct2Map(u))
	}
	return u, err
}
