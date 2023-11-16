package common

import "github.com/gofiber/fiber/v2"

type Resp struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    any    `json:"data"`
}

func RespSuccess(c *fiber.Ctx, msg string, data any) error {
	return c.JSON(Resp{
		Success: true,
		Msg:     msg,
		Data:    data,
	})
}

func RespError(c *fiber.Ctx, code int, msg string) error {
	return c.Status(code).JSON(Resp{
		Success: false,
		Msg:     msg,
	})
}
