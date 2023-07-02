package handler

import "github.com/gofiber/fiber/v2"

func ApiDoc(router fiber.Router) {
	router.Static("/apidoc/*", "apidoc")
}
