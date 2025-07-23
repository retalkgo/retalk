package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/retalkgo/retalk/server/common"
)

type HealthzHandler struct{}

var Healthz = &HealthzHandler{}

// 健康检查
//
//	@summary		健康检查
//	@description	健康检查
//	@tags			健康检查
//	@success		200	{boolean}	boolean	"返回健康状态"
//	@router			/healthz [GET]
func (h *HealthzHandler) Healthz(ctx context.Context, c *app.RequestContext) {
	common.RespSuccess(c, true)
}
