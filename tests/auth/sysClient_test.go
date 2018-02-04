package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"

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
	e.WithQueryString(values.Encode()).Expect().Status(http.StatusOK)
}

func TestClientPut(t *testing.T) {
	client := new(auth.SysClient)
	client.ID = 10
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
	e.WithJSON(values).Expect().Status(http.StatusCreated)
}

func TestClientDelete(t *testing.T) {
	method := "DELETE"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	RequestURL := "/v1/client/10"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	e.Expect().Status(http.StatusNoContent)
}
func TestClientByDefault(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	RequestURL := "/v1/client/default"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	e.Expect().Status(http.StatusOK)
}
