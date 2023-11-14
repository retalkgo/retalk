package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/retalkgo/retalk/internal/config"
)

func Start() {
	config.InitConfig()

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
	if c.Status.Enable {
		var title string
		if c.Status.Title != "" {
			title = c.Status.Title
		} else {
			title = "retalk status"
		}
		app.Get("/status", monitor.New(monitor.Config{
			Title:   title,
			FontURL: "https://jsd.onmicrosoft.cn/npm/@recomponent/react@0.1.4/dist/recomponent-global.css",
		}))
	} else {
		app.Get("/status", func(c *fiber.Ctx) error {
			return c.SendString("未")
		})
	}

	// Listen
	app.Listen(fmt.Sprintf(":%d", c.Server.Port))
}
