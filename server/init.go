package server

import (
	h "retalk/server/handler"
	m "retalk/server/middleware"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	// 初始化中间件
	m.Logger(app)
	m.Cors(app)

	// 路由注册

	// Swagger
	h.Swagger(app)

	// 主页路由
	h.Home(app)

	// api路由
	api := app.Group("/api")
	h.Home(api)

	// 404路由
	h.NotFound(app)
}
