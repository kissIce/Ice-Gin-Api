package core

import (
	"fmt"
	"github.com/go-redis/redis"
	"ice/global"
)

func initCache()  {
	switch global.IceConfig.System.Cachetype {
	case "redis":
		initRedis()
	}
}

func initRedis()  {
	redisCfg := global.IceConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr: redisCfg.Uri + ":" + redisCfg.Port,
		Password: redisCfg.Password,
		DB: redisCfg.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("redis 链接失败，错误原因: %v \n", err)
	} else {
		global.IceRedis = client
	}
}