package sign

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/xwinie/glue/core/middleware/sign"
	"github.com/xwinie/glue/tests"
)

func TestSignatureGet(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := make(url.Values)
	values.Add("sort", "asc")
	RequestURL := "/a"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL+"?"+values.Encode(), timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	repos := e.WithJSON(values).Expect().Status(http.StatusForbidden).JSON()
	Convey("Subject: 数字签名\n", t, func() {
		So(repos.String().Raw(), ShouldEqual, "Signature Failed")
	})
}
