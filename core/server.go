package core

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ice/global"
	"ice/utils/snowflake"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer()  {
	// 初始化数据库
	initDB()
	if global.IceConfig.System.Cache {
		initCache()
	}
	// 初始化雪花算法
	global.IceSnowflake, _ = snowflake.NewWorker(global.IceConfig.SnowFlake.Workid)
	// 初始化路由
	router := initRouter()
	// 初始化服务
	addr := fmt.Sprintf(":%d", global.IceConfig.System.Addr)
	s := initServer(addr, router)
	time.Sleep(10 * time.Microsecond)
	global.IceLog.Info("server run success on ", zap.String("address", addr))
	global.IceLog.Error(s.ListenAndServe().Error())
	// 程序结束前关闭数据库
	db, _ := global.IceDb.DB()
	defer db.Close()
}

func initServer(addr string, gen *gin.Engine) server {
	s := endless.NewServer(addr, gen)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 8 << 20
	return s
}
