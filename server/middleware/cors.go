package middleware

import (
	"github.com/gofiber/fiber/v2"
	c "github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors(app *fiber.App) {
	app.Use(c.New(c.Config{
		AllowOrigins: "*",
	}))
}
