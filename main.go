package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zxgit/gin-blog-project/global"
	"github.com/zxgit/gin-blog-project/internal/dbClient"
	"github.com/zxgit/gin-blog-project/internal/routers"
	"github.com/zxgit/gin-blog-project/pkg/logger"
	"github.com/zxgit/gin-blog-project/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init()  {
	err:=setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v",err)
	}
	err = setupLogger()
	if err!=nil{
		log.Fatalf("init.setupLogger %v",err)
	}
	dbClient.MysqlInit()
	dbClient.InitRedisClient()
}

func setupSetting() error{
	setting,err := setting.NewSetting()
	if err != nil {
		return err
	}
	err =setting.ReadSection("Server",&global.ServerSetting)
	if err != nil {
		return err
	}
	err =setting.ReadSection("App",&global.AppSetting)
	if err != nil {
		return err
	}
	err =setting.ReadSection("Database",&global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *=time.Second
	global.ServerSetting.WriteTimeout *=time.Second
	return nil
}

func main() {
	fmt.Println("dbtype:",global.DatabaseSetting.DBType)
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:          ":"+global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}



func setupLogger() error{
	fileName :=global.AppSetting.LogSavePath+"/"+global.AppSetting.LogFileName+global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: fileName,
		MaxSize: 600,
		MaxAge: 10,
		LocalTime: true,
	},"",log.LstdFlags).WithCaller(2)
	return nil
}
