package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/retalkgo/retalk/server/handler"
)

func registerHealtzRoute(group *route.RouterGroup) {
	group.GET("/healthz", handler.Healthz.Healthz)
}
