package middleware

import "github.com/gofiber/fiber/v2"

func Auth(router fiber.Router) {
	router.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})
}
