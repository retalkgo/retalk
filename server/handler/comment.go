package handler

import "github.com/gofiber/fiber/v2"

func Comment(router fiber.Router) {
	router.Post("/comment")
}