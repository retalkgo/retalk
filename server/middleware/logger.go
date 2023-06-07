package middleware

import (
	"fmt"
	"time"

	l "retalk/internal/logger"

	"github.com/gofiber/fiber/v2"
)

var (
	startTime int64
)

func Logger(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		startTime = time.Now().UnixMicro()
		err := c.Next()
		l.Info(fmt.Sprintf("%s %s, 处理时间: %d 微秒", c.Method(), c.Path(), time.Now().UnixMicro()-startTime))
		return err
	})
}
