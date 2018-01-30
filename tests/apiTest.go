package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/xwinie/glue"
)

//TestAPI 测试基础方法
func TestAPI(t *testing.T, appid, signature, timestamp, Authorization, method, path string, pathargs ...interface{}) *httpexpect.Request {
	// gin handle
	handler := glue.Engine()
	server := httptest.NewServer(handler)
	defer server.Close()
	e := httpexpect.New(t, server.URL)
	if Authorization == "" {
		return e.Request(method, path, pathargs).
			WithHeader("appid", appid).
			WithHeader("signature", signature).
			WithHeader("timestamp", timestamp)
	}
	return e.Request(method, path, pathargs).
		WithHeader("appid", appid).
		WithHeader("signature", signature).
		WithHeader("timestamp", timestamp).
		WithHeader("Authorization", Authorization)

}
