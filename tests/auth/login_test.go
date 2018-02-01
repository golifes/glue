package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"
	"yh-foundation-backend/cores"

	"github.com/xwinie/glue/lib/middleware/sign"
	. "github.com/xwinie/glue/lib/utils"
	"github.com/xwinie/glue/tests"
)

func TestLoginPost(t *testing.T) {
	method := "POST"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := map[string]interface{}{
		"UserName": "12345",
		"Password": Md5(cores.Sha1("12345") + Sha1("Password")),
	}
	jsonValue, _ := json.Marshal(values)
	RequestURL := "/v1/login"
	signature := sign.Signature("Lx1b8JoZoE", method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, "")
	e.WithJSON(values).Expect().Status(http.StatusCreated)
}
