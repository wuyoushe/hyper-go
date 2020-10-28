package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/griffin702/ginana/library/conf/paladin"
	"github.com/griffin702/ginana/library/log"

	// _ "hyper-go/tool\hyper/felton_blog/docs" // 使用Swagger时打开
	"hyper-go/tool/hyper/felton_blog/internal/config"
	"hyper-go/tool/hyper/felton_blog/internal/wire"
)

// @title GiNana
// @version 1.0.0
// @description 基于iris的api服务，默认端口：8000
// @host 127.0.0.1:8000
// @BasePath /api
// @license.name MIT License
// @license.url
func main() {
	closeLog := log.Init()
	log.Info("GiNana App Start")
	cfg, err := config.GetBaseConfig()
	if err != nil {
		panic(err)
	}
	if err := paladin.Init(cfg.ConfigIsLocal, cfg.ConfigPath); err != nil {
		panic(err)
	}
	app, closeFunc, err := wire.InitApp()
	if err != nil {
		panic(err)
	}
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT)
		for {
			s := &lt
			-ch
			log.Infof("get a signal %s", s.String())
			switch s {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Info("GiNana App Exit")
				time.Sleep(time.Second)
				closeFunc()
				closeLog()
				return
			case syscall.SIGHUP:
			default:
				return
			}
		}
	}()
	err = app.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Println(err.Error())
	}
}
