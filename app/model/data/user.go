package data

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/mitchellh/mapstructure"
	"ice/app/model/dao"
	"ice/app/model/entity"
	"ice/global"
	"ice/helper"
	"math/rand"
	"strconv"
	"time"
)

func GetUserByPhone(phone string) (user *entity.User, err error) {
	var uid string
	var u *entity.User
	uid, err = global.IceRedis.Get("user:" + phone).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = dao.GetUserByPhone(u)
			if err != nil {
				global.IceRedis.Set("user:"+phone, 0, time.Duration(rand.Intn(20)+20)*time.Second)
			} else {
				global.IceRedis.Set("user:"+phone, u.Id, time.Duration(30*3600*24)*time.Second)
				global.IceRedis.HMSet("userInfo:" + strconv.FormatUint(u.Id, 10), map[string]interface{}{"id":1, "name":"username"})
			}
		}
	} else {
		uid, _ := strconv.ParseUint(uid, 10, 64)
		if uid == 0 {
			err = errors.New("用户不存在")
		} else {
			u, err = GetUserById(uid)
		}
	}
	return u, err
}

func GetUserById(uid uint64) (user *entity.User, err error) {
	var ret map[string]string
	u := &entity.User{
		Id: uid,
	}
	ret = global.IceRedis.HGetAll("userInfo:" + strconv.FormatUint(uid, 10)).Val()
	fmt.Println(len(ret))
	if len(ret) < 1 {
		u.Id = uid
		err = dao.GetUserById(u)
		if err != nil {
			global.IceRedis.HSet("userInfo:"+strconv.FormatUint(uid, 10), "ID", 0)
			global.IceRedis.Expire("userInfo:"+strconv.FormatUint(uid, 10), time.Duration(rand.Intn(20)+20)*time.Second)
		} else {
			global.IceRedis.HMSet("userInfo:"+strconv.FormatUint(uid, 10), helper.Struct2Map(user))
		}
	} else {
		err = mapstructure.Decode(ret, &user)
		if user.Id == 0 {
			err = errors.New("用户不存在")
		}
	}
	return
}
