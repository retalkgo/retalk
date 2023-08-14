package server

import (
	_ "github.com/retalkgo/retalk/docs"
	"github.com/retalkgo/retalk/internal/version"

	"github.com/gofiber/fiber/v2"
)

//	@Title						Retalk API
//	@Version					1.0
//	@Description				Retalk 后端 API 文档
//	@BasePath					/
//	@Contact.name				API 支持
//	@Contact.email				retalk@redish101.top
//	@License.name				GPL-3.0
//	@SecurityDefinitions.apikey	ApiKeyAuth
//	@In							header
//	@Name						Authorization
//	@Description				"将 "Bearer TOKEN" 设置为初始化时设置的 API 密钥"

func Start() {
	app := fiber.New(fiber.Config{
		AppName:      "Retalk " + version.Version + "-" + version.CommitHash,
		ServerHeader: "retalk",
		Prefork:      false,
	})
	Init(app)
	app.Listen(":3000")
}
