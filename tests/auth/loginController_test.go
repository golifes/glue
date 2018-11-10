package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/xwinie/glue/core"
	"github.com/xwinie/glue/core/middleware/sign"
	"github.com/xwinie/glue/tests"
)

func TestLoginPost(t *testing.T) {
	method := "POST"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := map[string]interface{}{
		"UserName": "12345",
		"Password": core.Md5(core.Sha1("12345") + core.Sha1("Password")),
	}
	jsonValue, _ := json.Marshal(values)
	body := bytes.NewBuffer(jsonValue)
	RequestURL := "/v1/login"
	signature := sign.Signature("Lx1b8JoZoE", method, body.Bytes(), RequestURL, timestamp)

	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, "")
	repos := e.WithJSON(values).Expect().Status(http.StatusCreated).JSON().Object()
	Convey("Subject: 用户登录\n", t, func() {
		Convey("Accoutn Should Be 12345", func() {
			So(repos.Value("Account").String().Raw(), ShouldEqual, "12345")
		})
	})
}
