package router

import "github.com/cloudwego/hertz/pkg/app/server"

func RegisterRoutes(h *server.Hertz) {
	api := h.Group("/api")

	registerHealtzRoute(api)
	registerNotFoundRoute(api)
}
