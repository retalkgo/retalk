package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/retalkgo/retalk/internal/auth"
	"github.com/retalkgo/retalk/server/common"
)

func AdminLogin(router fiber.Router) {
	router.Get("/login", func(c *fiber.Ctx) error {
		token, _ := auth.GenerateToken("redish101", "12346")
		return common.RespSuccess(c, "", token)
	})
	router.Get("/test", func(c *fiber.Ctx) error {
		token := c.Query("token", "")
		fmt.Println(token)
		t, _ := auth.Verify(token)
		return common.RespSuccess(c, "", t)
	})
}