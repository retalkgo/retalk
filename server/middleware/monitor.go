package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/retalkgo/retalk/internal/config"
	"github.com/retalkgo/retalk/internal/i18n"
)

func Monitor(config *config.ConfigSchema, app *fiber.App) {
	if config.Status.Enable {
		var title string
		if config.Status.Title != "" {
			title = config.Status.Title
		} else {
			title = "retalk status"
		}
		app.Get("/status", monitor.New(monitor.Config{
			Title:   title,
			FontURL: "https://jsd.onmicrosoft.cn/npm/@recomponent/react@0.1.4/dist/recomponent-global.css",
		}))
	} else {
		app.Get("/status", func(c *fiber.Ctx) error {
			return c.SendString(i18n.I18n("dashboard_not_enable"))
		})
	}
}
