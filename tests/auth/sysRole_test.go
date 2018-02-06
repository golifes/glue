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

func TestRolePost(t *testing.T) {
	method := "POST"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := map[string]interface{}{
		"Code": "1234567",
		"Name": "测试角色",
	}
	jsonValue, _ := json.Marshal(values)
	RequestURL := "/v1/role"
	signature := sign.Signature("Lx1b8JoZoE", method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	e.WithJSON(values).Expect().Status(http.StatusCreated)
}
func TestRolePut(t *testing.T) {
	o := core.New()
	var m auth.SysRole
	o.Table("sys_role").Where("code = ?", "1234567").Get(&m)
	method := "PUT"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := map[string]interface{}{
		"Code": "12345678",
		"Name": "测试角色2",
	}
	jsonValue, _ := json.Marshal(values)
	RequestURL := "/v1/role/" + m.ID
	signature := sign.Signature("Lx1b8JoZoE", method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	e.WithJSON(values).Expect().Status(http.StatusCreated)
}
func TestRoleByPage(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := make(url.Values)
	values.Add("p", "1")
	values.Add("perPage", "10")
	RequestURL := "/v1/role"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL+"?"+values.Encode(), timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	e.WithQueryString(values.Encode()).Expect().Status(http.StatusOK)
}
func TestRoleByAccount(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	RequestURL := "/v1/role/1234567"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	e.Expect().Status(http.StatusOK)
}

func TestRoleDelete(t *testing.T) {
	o := core.New()
	var m auth.SysRole
	o.Table("sys_role").Where("code = ?", "1234567").Get(&m)
	method := "DELETE"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	RequestURL := "/v1/role/" + m.ID
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	e.Expect().Status(http.StatusNoContent)
}

func TestRoleAllotResource(t *testing.T) {
	role := new(auth.SysRole)
	role.ID = "100"
	role.Code = "100"
	role.Name = "管理员100"
	resource := []auth.SysResource{
		{ID: "1", Code: "1", Action: "/v1/test1", Method: "POST", Name: "测试", IsOpen: 0, ResType: 0, ParentID: "0"},
		{ID: "2", Code: "2", Action: "/v1/test2", Method: "POST", Name: "测试", IsOpen: 0, ResType: 0, ParentID: "0"},
	}
	o := core.New()
	o.Insert(role)
	o.Insert(resource)
	method := "POST"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := map[string]interface{}{
		"ResourceId": []string{"1", "2"},
	}
	jsonValue, _ := json.Marshal(values)
	RequestURL := "/v1/role/" + strconv.FormatInt(100, 10) + "/resource"
	signature := sign.Signature("Lx1b8JoZoE", method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	e.WithJSON(values).Expect().Status(http.StatusOK)
}

func TestResourceByRoleID(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	RequestURL := "/v1/role/100/resource"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	e.Expect().Status(http.StatusOK)
}
