package middleware

import (
	"fmt"
	"time"

	"github.com/retalkgo/retalk/internal/i18n"
	l "github.com/retalkgo/retalk/internal/logger"

	"github.com/gofiber/fiber/v2"
)

var startTime int64

func Logger(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		startTime = time.Now().UnixMicro()
		err := c.Next()
		l.Info(fmt.Sprintf("%s %s, %s: %d %s", c.Method(), c.Path(), i18n.I18n("processingTime"), time.Now().UnixMicro()-startTime, i18n.I18n("microsecond")))
		return err
	})
}
