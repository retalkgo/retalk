package common

import (
	"errors"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

type FailureResp struct {
	Error string `json:"error"`
}

func RespSuccess(c *app.RequestContext, data any) {
	c.JSON(http.StatusOK, data)
}

func RespSuccessWithCode(c *app.RequestContext, code int, data any) {
	c.JSON(code, data)
}

func RespFailure(c *app.RequestContext, err error) {
	c.JSON(http.StatusOK, FailureResp{
		Error: err.Error(),
	})
}

func RespInvalidParams(c *app.RequestContext) {
	RespFailure(c, errors.New("无效的参数"))
}
