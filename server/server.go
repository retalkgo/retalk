package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/retalkgo/retalk/internal/config"
	"github.com/retalkgo/retalk/server/handler"
	"github.com/retalkgo/retalk/server/middleware"
)

func Start() {
	c := config.Config()

	app := fiber.New(fiber.Config{
		Prefork:      true,
		AppName:      "retalk",
		ServerHeader: "retalk",
	})

	// Middleware

	// Logger
	app.Use(logger.New(logger.ConfigDefault))

	// Monitor
	middleware.Monitor(c, app)

	admin := app.Group("/admin")
	middleware.Auth(admin)

	handler.AdminLogin(admin)

	// Listen
	app.Listen(fmt.Sprintf(":%d", c.Server.Port))
}
