package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/xwinie/glue/app/auth"
	"github.com/xwinie/glue/core"
	"github.com/xwinie/glue/core/middleware/sign"
	"github.com/xwinie/glue/tests"
)

func TestClientrByPage(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := make(url.Values)
	values.Add("p", "1")
	values.Add("perPage", "10")
	RequestURL := "/v1/client"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL+"?"+values.Encode(), timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	repos := e.WithQueryString(values.Encode()).Expect().Status(http.StatusOK).JSON().Object()
	Convey("Subject: 分页获取客户端\n", t, func() {
		So(repos.Value("Clients").Array().First().Object().Value("ClientId").String().Raw(), ShouldEqual, "app1")
	})
}

func TestClientPut(t *testing.T) {
	client := new(auth.SysClient)
	client.ID = "10"
	client.ClientID = "app10"
	client.Name = "测试app10"
	client.Secret = "Lx1b8JoZoE"
	client.VerifySecret = "Lx1b8JoZoE"
	o := core.New()
	o.Insert(client)
	method := "PUT"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := map[string]interface{}{
		"ClientID": "12345678",
		"Name":     "测试角色2",
	}
	jsonValue, _ := json.Marshal(values)
	RequestURL := "/v1/client/10"
	signature := sign.Signature("Lx1b8JoZoE", method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	repos := e.WithJSON(values).Expect().Status(http.StatusCreated).JSON().Object()
	Convey("Subject: 根据id修改客户端信息\n", t, func() {
		So(repos.Value("_links").Array().First().Object().Value("href").String().Raw(), ShouldEqual, "/v1/client/app10")
	})
}

func TestClientDelete(t *testing.T) {
	method := "DELETE"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	RequestURL := "/v1/client/10"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	repos := e.Expect().Status(http.StatusNoContent).JSON().Object()
	Convey("Subject: 根据id删除客户端信息\n", t, func() {
		So(repos.Value("code").Raw(), ShouldEqual, 100000)
	})
}
func TestClientByDefault(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	RequestURL := "/v1/client/default"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	repos := e.Expect().Status(http.StatusOK).JSON().Object()
	Convey("Subject: 根据id获取客户端信息\n", t, func() {
		So(repos.Value("Client").Object().Value("ClientID").String().Raw(), ShouldEqual, "app1")
	})
}
