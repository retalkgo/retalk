package common

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Validator(c *fiber.Ctx, data interface{}) {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		RespError(c, "参数缺失", nil, http.StatusBadRequest)
	}
}