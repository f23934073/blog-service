package main

import (
	"log"
	"net/http"
	"time"

	"blog-service/global"
	"blog-service/internal/model"
	"blog-service/internal/routers"
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

// @title 部落格後端系統
// @version 1.0
// @description Steve 的部落格設計後端
func main() {
	//init
	if err := setupSetting(); err != nil {
		panic(err)
	}
	if err := setupDBEngine(); err != nil {
		panic(err)
	}
	if err := setupLogger(); err != nil {
		panic(err)
	}

	gin.SetMode(global.ServerSetting.RunMode)

	router := routers.NewRouters()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

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

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
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
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
