package common

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/stretchr/testify/assert"
)

func TestRespSuccess(t *testing.T) {
	c := ut.CreateUtRequestContext("GET", "/", nil)

	// 定义要发送的测试数据
	testData := map[string]string{"message": "success"}

	// 调用 RespSuccess 函数
	RespSuccess(c, testData)

	// 使用 testify/assert 断言
	assert.Equal(t, http.StatusOK, c.Response.StatusCode(), "状态码应该为 200")

	// 验证响应体
	var respBody map[string]string
	err := json.Unmarshal(c.Response.Body(), &respBody)
	assert.NoError(t, err, "响应体应该可以被正确解析为 JSON")
	assert.Equal(t, testData, respBody, "响应体内容应该与测试数据一致")
}

func TestRespSuccessWithCode(t *testing.T) {
	c := ut.CreateUtRequestContext("GET", "/", nil)
	testData := "data"
	testCode := http.StatusCreated // 201

	RespSuccessWithCode(c, testCode, testData)

	assert.Equal(t, testCode, c.Response.StatusCode(), "状态码应该为指定的 201")

	var respBody string
	err := json.Unmarshal(c.Response.Body(), &respBody)
	assert.NoError(t, err)
	assert.Equal(t, testData, respBody, "响应体内容应该与测试数据一致")
}

func TestRespFailure(t *testing.T) {
	c := ut.CreateUtRequestContext("GET", "/", nil)
	testError := errors.New("这是一个测试错误")

	RespFailure(c, testError)

	assert.Equal(t, http.StatusOK, c.Response.StatusCode(), "状态码应该为 200")

	var respBody FailureResp
	err := json.Unmarshal(c.Response.Body(), &respBody)
	assert.NoError(t, err)
	assert.Equal(t, testError.Error(), respBody.Error, "响应体中的错误信息应该与测试错误一致")
}

func TestRespInvalidParams(t *testing.T) {
	c := ut.CreateUtRequestContext("GET", "/", nil)

	RespInvalidParams(c)

	assert.Equal(t, http.StatusOK, c.Response.StatusCode(), "状态码应该为 200")

	var respBody FailureResp
	err := json.Unmarshal(c.Response.Body(), &respBody)
	assert.NoError(t, err)
	assert.Equal(t, "无效的参数", respBody.Error, "响应体中的错误信息应为 '无效的参数'")
}
