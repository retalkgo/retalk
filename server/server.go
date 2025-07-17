package server

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/retalkgo/retalk/internal/config"
	"github.com/retalkgo/retalk/internal/version"
	"github.com/sirupsen/logrus"
)

// @title						retalk api
// @version					1.0
// @description				retalk api
// @license.name				GPL 3.0
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func Start() {
	logrus.Infof("retalk %s", version.Version)

	config := config.Config()

	fiberApp := fiber.New(fiber.Config{
		AppName:      "retalk",
		ServerHeader: "retalk",
	})

	listenAddr := config.Host + ":" + strconv.Itoa(config.Port)
	logrus.Printf("[HTTP] 在 http://%s 启动服务", listenAddr)
	fiberApp.Listen(listenAddr, fiber.ListenConfig{
		DisableStartupMessage: true,
	})
}
