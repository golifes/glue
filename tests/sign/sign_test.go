package sign

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/xwinie/glue/lib/middleware/sign"
	"github.com/xwinie/glue/tests"
)

func TestSignatureGet(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := make(url.Values)
	values.Add("sort", "asc")
	RequestURL := "/a"
	signature := sign.Signature("1", method, nil, RequestURL+"?"+values.Encode(), timestamp)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, "")
	e.WithQueryString(values.Encode()).Expect().Status(http.StatusOK)
}
