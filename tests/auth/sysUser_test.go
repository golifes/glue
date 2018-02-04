package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/xwinie/glue/app/auth"

	"github.com/xwinie/glue/core"
	"github.com/xwinie/glue/core/middleware/sign"
	"github.com/xwinie/glue/tests"
)

func TestUserPost(t *testing.T) {
	method := "POST"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := map[string]interface{}{
		"Account":  "1234567",
		"Password": core.Md5(core.Sha1("1234567") + core.Sha1("Password")),
		"Name":     "测试",
	}
	jsonValue, _ := json.Marshal(values)
	RequestURL := "/v1/user"
	signature := sign.Signature("Lx1b8JoZoE", method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	e.WithJSON(values).Expect().Status(http.StatusCreated)
}
func TestUserPut(t *testing.T) {
	o := core.New()
	var user auth.SysUser
	o.Table("sys_user").Where("account = ?", "1234567").Get(&user)
	method := "PUT"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := map[string]interface{}{
		"Account": "12345678",
		"Name":    "测试2",
	}
	jsonValue, _ := json.Marshal(values)
	RequestURL := "/v1/user/" + strconv.FormatInt(user.ID, 10)
	signature := sign.Signature("Lx1b8JoZoE", method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	e.WithJSON(values).Expect().Status(http.StatusCreated)
}
func TestUserByPage(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := make(url.Values)
	values.Add("p", "1")
	values.Add("perPage", "10")
	RequestURL := "/v1/user"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL+"?"+values.Encode(), timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	e.WithQueryString(values.Encode()).Expect().Status(http.StatusOK)
}
func TestUserByAccount(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	RequestURL := "/v1/user/1234567"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	e.Expect().Status(http.StatusOK)
}

func TestUserDelete(t *testing.T) {
	o := core.New()
	var user auth.SysUser
	o.Table("sys_user").Where("account = ?", "1234567").Get(&user)
	method := "DELETE"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	RequestURL := "/v1/user/" + strconv.FormatInt(user.ID, 10)
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	e.Expect().Status(http.StatusNoContent)
}
