package server

import (
	"github.com/retalkgo/retalk/internal/db"
	"github.com/retalkgo/retalk/internal/logger"
	h "github.com/retalkgo/retalk/server/handler"
	m "github.com/retalkgo/retalk/server/middleware"

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

	// 服务器初始化
	h.Init(api)

	// 评论
	comment := api.Group("/comment")
	h.CommentGetAll(comment)
	h.CommentGetByPath(comment)
	h.CommentAdd(comment)

	// 404路由
	h.NotFound(app)
}

func InitVercel(app *fiber.App) {
	_, err := db.InitDB()
	logger.Info(err)
}