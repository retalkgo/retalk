package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/retalkgo/retalk/server"
)

var srv http.Handler

func init() {
	app := fiber.New()
	server.InitVercel(app)
	srv = adaptor.FiberApp(app)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	srv.ServeHTTP(w, r)
}
