package handler

import (
	"github.com/retalkgo/retalk/server/common"

	"github.com/gofiber/fiber/v2"
)

//	@Summary		首页
//	@Description	输出欢迎信息以验证安装
//	@Tags			首页
//	@Success		200	{object}	common.Resp{msg=string}
//	@Router			/ [get]
func Home(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return common.RespSuccess(c, "欢迎使用Retalk", nil)
	})
}
