package tests

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/xwinie/glue"
)

func testHander(t *testing.T) *httpexpect.Expect {
	handler := glue.Hander()
	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
	return e
}

//TestAPI 测试基础方法
func TestAPI(t *testing.T, method, path, appid, signature, timestamp, authorization string) *httpexpect.Request {
	e := testHander(t)

	headers := make(map[string]string, 4)
	if authorization != "" {
		headers["Authorization"] = authorization
	}
	headers["appid"] = appid
	headers["signature"] = signature
	headers["timestamp"] = timestamp

	return e.Request(method, path).WithHeaders(headers)
}
