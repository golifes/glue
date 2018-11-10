package auth

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/xwinie/glue/core/middleware/sign"
	"github.com/xwinie/glue/tests"
)

func TestResourcerByPage(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := make(url.Values)
	values.Add("p", "1")
	values.Add("perPage", "10")
	RequestURL := "/v1/resource"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL+"?"+values.Encode(), timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	repos := e.WithQueryString(values.Encode()).Expect().Status(http.StatusOK).JSON().Object()
	Convey("Subject: 分页获取资源信息\n", t, func() {
		So(repos.Value("Resources").Array().First().Object().Value("Action").String().Raw(), ShouldEqual, "/v1/login")
	})
}
func TestMenusByUserID(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	RequestURL := "/v1/menus/1"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	repos := e.Expect().Status(http.StatusOK).JSON()
	Convey("Subject: 根据用户id获取菜单信息\n", t, func() {
		So(repos.Array().First().Object().Value("Code").Raw(), ShouldEqual, "10034")
	})
}

func TestResourceByCode(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	RequestURL := "/v1/resource/10000"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	repos := e.Expect().Status(http.StatusOK).JSON().Object()
	Convey("Subject: 根据资源编号获取资源信息\n", t, func() {
		So(repos.Value("Action").String().Raw(), ShouldEqual, "/v1/login")
	})
}
