package server

import (
	"strconv"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/retalkgo/retalk/internal/config"
	"github.com/retalkgo/retalk/internal/db"

	"github.com/retalkgo/retalk/internal/store"
	"github.com/retalkgo/retalk/internal/version"
	"github.com/retalkgo/retalk/server/handler"
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
func Start() {
	logrus.Infof("retalk %s", version.Version)

	config := config.LaunchConfig()

	if config.Dev {
		logrus.SetLevel(logrus.DebugLevel)
	}

	err := store.Init(db.DB(), &config.Cache)
	if err != nil {
		logrus.Fatalf("[STORE] 初始化储存层时失败: %v", err)
	}

	listenAddr := config.Server.Host + ":" + strconv.Itoa(config.Server.Port)

	h := server.New(
		server.WithHostPorts(listenAddr),
	)

	registerRoutes(h)

	logrus.Infof("[HTTP] 在 http://%s 启动服务", listenAddr)

	h.Spin()
}

func registerRoutes(app *server.Hertz) {
	app.GET("/healthz", handler.Healthz.Healthz)

	app.GET("/*any", handler.NotFound.NotFound)
}
