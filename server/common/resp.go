package common

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Resp struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func RespSuccess(c *fiber.Ctx, msg string, data interface{}) error {
	return c.Status(http.StatusOK).JSON(&Resp{
		Success: true,
		Msg:     msg,
		Data:    data,
	})
}

func RespError(c *fiber.Ctx, msg string, data interface{}, code int) error {
	return c.Status(code).JSON(&Resp{
		Success: false,
		Msg:     msg,
		Data:    data,
	})
}
