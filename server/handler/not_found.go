package handler

import (
	"net/http"

	"github.com/retalkgo/retalk/server/common"

	"github.com/gofiber/fiber/v2"
)

func NotFound(router fiber.Router) {
	router.All("*", func(c *fiber.Ctx) error {
		return common.RespError(c, "404 Not Found", nil, http.StatusNotFound)
	})
}
