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
	adminCacheKey = "adminInfo:"
	adminPhoneKey = "adminPhone:"
)

func GetAdminByPhone(phone string, field []string) (ret *entity.Admin, err error) {
	var adminId string
	adminId, err = global.IceRedis.Get(adminPhoneKey + phone).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			ret, err = cacheAdmin(map[string]interface{}{"phone": phone})
			if err != nil {
				global.IceRedis.Set(adminPhoneKey+phone, 0, time.Duration(helper.RandRangeInt(10, 60))*time.Second)
			} else {
				global.IceRedis.Set(adminPhoneKey+phone, ret.Id, time.Duration(cacheExpire)*time.Second)
			}
		}
	} else {
		adminId, _ := strconv.ParseUint(adminId, 10, 64)
		if adminId == 0 {
			err = gorm.ErrRecordNotFound
		} else {
			ret, err = GetAdminById(adminId, field)
		}
	}
	return
}

/**
 * 根据用户id从缓存查找用户数据
 */
func GetAdminById(uid uint64, field []string) (*entity.Admin, error) {
	var ret = new(entity.Admin)
	var err error
	redisCache := global.IceRedis.HMGet(adminCacheKey+strconv.FormatUint(uid, 10), field...).Val()
	if redisCache[0] == nil {
		ret, err = cacheAdmin(map[string]interface{}{"id": uid})
		if err != nil {
			global.IceRedis.HSet(adminCacheKey+strconv.FormatUint(uid, 10), "ID", 0)
			global.IceRedis.Expire(adminCacheKey+strconv.FormatUint(uid, 10), time.Duration(helper.RandRangeInt(10, 99))*time.Second)
		}
	} else {
		err = helper.Slice2Struct(field, redisCache, ret)
	}
	return ret, err
}

/**
 * 获取并缓存用户信息返回指定字段
 */
func cacheAdmin(where map[string]interface{}) (*entity.Admin, error) {
	admin, err := dao.GetAdmin(where)
	if err == nil {
		global.IceRedis.HMSet(adminCacheKey+strconv.FormatUint(admin.Id, 10), helper.Struct2Map(admin))
	}
	return admin, err
}

func GetAdminList() (adminList []entity.Admin, err error) {
	return dao.GetAdminList()
}

/**
 * 创建用户并删除请求缓存
 */
func AddAdmin(u *entity.Admin) (err error) {
	err = dao.AddAdmin(u)
	if err == nil {
		global.IceRedis.Del(adminPhoneKey+u.Phone)
	}
	return
}

func EditAdmin(admin *entity.Admin) error {
	return dao.EditAdmin(admin)
}

func DelAdmin(admin *entity.Admin) error {
	return dao.DelAdmin(admin)
}