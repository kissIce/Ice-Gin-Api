package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"ice/core/typestruct"
	"ice/utils/snowflake"
)

var (
	IceDb        *gorm.DB
	IceVp        *viper.Viper
	IceRedis     *redis.Client
	IceLog       *zap.Logger
	IceSnowflake *snowflake.Worker
	IceConfig    typestruct.Server
)
