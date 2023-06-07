package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/swagger"
)

func Swagger(router fiber.Router) {
	router.All("/swagger/*", swagger.HandlerDefault)
}
