package handler

import (
	"net/http"

	"github.com/retalkgo/retalk/internal/i18n"
	"github.com/retalkgo/retalk/server/common"

	"github.com/gofiber/fiber/v2"
)

func NotFound(router fiber.Router) {
	router.All("*", func(c *fiber.Ctx) error {
		return common.RespError(c, i18n.I18n("notFound"), nil, http.StatusNotFound)
	})
}
