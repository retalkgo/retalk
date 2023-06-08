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

	// apidoc
	h.ApiDoc(app)

	// 主页
	h.Home(app)

	// api
	api := app.Group("/api")
	h.Home(api)

	// 评论
	comment := api.Group("/comment")
	h.CommentGetAll(comment)

	// 404路由
	h.NotFound(app)
}
