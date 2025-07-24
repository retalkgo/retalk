package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/retalkgo/retalk/server/handler"
)

func registerNotFoundRoute(group *route.RouterGroup) {
	group.Any("/*any", handler.NotFound.NotFound)
}
