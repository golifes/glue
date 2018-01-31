package sign

import (
	"bytes"
	"encoding/json"
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
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL+"?"+values.Encode(), timestamp)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, "")
	e.WithQueryString(values.Encode()).Expect().Status(http.StatusOK)
}

func TestSignaturePost(t *testing.T) {
	method := "POST"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := map[string]interface{}{
		"message": "1111",
		"nick":    "测试",
	}
	jsonValue, _ := json.Marshal(values)
	RequestURL := "/form_post"
	signature := sign.Signature("Lx1b8JoZoE", method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, "")
	e.WithJSON(values).Expect().Status(http.StatusOK)
}
