package handler

import (
	"retalk/server/common"

	"github.com/gofiber/fiber/v2"
)

// @Summary		首页
// @Description	欢迎页面
// @Tags			欢迎
// @Success		200	{object}	common.Resp
// @Route			/ [get]
func Home(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return common.RespSuccess(c, "欢迎使用Retalk", nil)
	})
}
