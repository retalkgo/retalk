package server

import (
	_ "github.com/retalkgo/retalk/docs"
	"github.com/retalkgo/retalk/internal/version"

	"github.com/gofiber/fiber/v2"
)

//	@Title			Retalk API
//	@Version		1.0
//	@Description	Retalk 后端 API 文档
//	@BasePath		/

//	@Contact.name	API 支持
//	@Contact.email	retalk@redish101.top

//	@License.name	GPL-3.0

func Start() {
	app := fiber.New(fiber.Config{
		AppName:      "Retalk " + version.Version + "-" + version.CommitHash,
		ServerHeader: "retalk",
		Prefork:      true,
	})
	Init(app)
	app.Listen(":3000")
}
