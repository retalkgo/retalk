package server

import (
	"strconv"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/retalkgo/retalk/internal/config"
	"github.com/retalkgo/retalk/internal/db"

	"github.com/retalkgo/retalk/internal/store"
	"github.com/retalkgo/retalk/internal/version"
	"github.com/retalkgo/retalk/server/router"
	"github.com/sirupsen/logrus"
)

// 启动服务端
//
//	@title						retalk api
//	@version					1.0
//	@description				retalk api
//	@license.name				GPL 3.0
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@basePath					/api
func Start() {
	logrus.Infof("retalk %s", version.Version)

	config := config.LaunchConfig()

	if config.Dev {
		logrus.SetLevel(logrus.DebugLevel)
		hlog.SetLevel(hlog.LevelDebug)

		logrus.Infoln("以开发模式运行")
	} else {
		hlog.SetLevel(hlog.LevelInfo)
		hlog.SetSilentMode(true)
	}

	err := store.Init(db.DB(), &config.Cache)
	if err != nil {
		logrus.Fatalf("[STORE] 初始化储存层时失败: %v", err)
	}

	listenAddr := config.Server.Host + ":" + strconv.Itoa(config.Server.Port)

	h := server.New(
		server.WithHostPorts(listenAddr),
	)

	router.RegisterRoutes(h)

	logrus.Infof("[HTTP] 在 http://%s 启动服务", listenAddr)

	if config.Dev {
		h.Run()
	} else {
		h.Spin()
	}
}
