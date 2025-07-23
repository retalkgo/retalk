package handler

import (
	"net/http"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/stretchr/testify/assert"
)

func TestHealthz(t *testing.T) {
	r := server.Default().Engine

	r.GET("/healthz", Healthz.Healthz)
	resp := ut.PerformRequest(r, "GET", "/healthz", nil).Result()

	assert.Equal(t, resp.StatusCode(), http.StatusOK)
	assert.True(t, string(resp.Body()) == "true")
}
