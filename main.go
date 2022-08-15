package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phial3/go-server-template/global"
	"github.com/phial3/go-server-template/internal/model"
	"github.com/phial3/go-server-template/pkg/logger"
	"github.com/phial3/go-server-template/pkg/setting"
	"github.com/phial3/go-server-template/routers"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	//参数配置初始化
	err := setupSetting()
	if err != nil {
		fmt.Printf("init setting error: %v", err)
	}
	//数据库连接初始化
	err = setupDBEngine()
	if err != nil {
		fmt.Printf("init db fail. %v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           router,
		ReadHeaderTimeout: global.ServerSetting.ReadTimeOut,
		WriteTimeout:      global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes:    1 << 20,
	}
	s.ListenAndServe()
}

// setupSetting 初始化配置参数
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	global.JWTSetting.Expire *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
