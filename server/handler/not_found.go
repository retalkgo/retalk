package handler

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/retalkgo/retalk/server/common"
)

type NotFoundHandler struct{}

var NotFound = &NotFoundHandler{}

func (h *NotFoundHandler) NotFound(ctx context.Context, c *app.RequestContext) {
	common.RespFailure(c, errors.New("未知的路由"))
}
