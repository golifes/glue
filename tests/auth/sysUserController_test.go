package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
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
	repos := e.WithJSON(values).Expect().Status(http.StatusCreated).JSON().Object()
	Convey("Subject: 创建用户\n", t, func() {
		So(repos.Value("_links").Array().First().Object().Value("href").String().Raw(), ShouldEqual, "/v1/user/1234567")
	})
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
	RequestURL := "/v1/user/" + user.ID
	signature := sign.Signature("Lx1b8JoZoE", method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	repos := e.WithJSON(values).Expect().Status(http.StatusCreated).JSON().Object()
	Convey("Subject: 修改用户\n", t, func() {
		So(repos.Value("_links").Array().First().Object().Value("href").String().Raw(), ShouldEqual, "/v1/user/1234567")
	})
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
	repos := e.WithQueryString(values.Encode()).Expect().Status(http.StatusOK).JSON().Object()
	Convey("Subject: 分页获取用户\n", t, func() {
		So(repos.Value("Users").Array().First().Object().Value("Account").String().Raw(), ShouldEqual, "12345")
	})
}
func TestUserByAccount(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	RequestURL := "/v1/user/1234567"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	repos := e.Expect().Status(http.StatusOK).JSON().Object()
	Convey("Subject: 根据账号获取用户\n", t, func() {
		So(repos.Value("Account").String().Raw(), ShouldEqual, "1234567")
	})
}

func TestUserDelete(t *testing.T) {
	o := core.New()
	var user auth.SysUser
	o.Table("sys_user").Where("account = ?", "1234567").Get(&user)
	method := "DELETE"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	RequestURL := "/v1/user/" + user.ID
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	repos := e.Expect().Status(http.StatusNoContent).JSON().Object()
	Convey("Subject: 根据id删除用户\n", t, func() {
		So(repos.Value("code").Raw(), ShouldEqual, 100000)
	})
}

func TestUserAllotRole(t *testing.T) {
	user := new(auth.SysUser)
	user.ID = "10"
	salt := core.RandStringByLen(6)
	user.Account = "10"
	user.Name = "测试员工10"
	user.Password = core.Md5(core.Md5(core.Sha1("12345")+core.Sha1("Password")) + salt)
	user.Salt = salt
	role := []auth.SysRole{
		{ID: "10", Code: "10", Name: "测试管理员10"},
		{ID: "11", Code: "11", Name: "测试管理员11"},
	}
	o := core.New()
	o.Insert(user)
	o.Insert(role)
	method := "POST"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := map[string]interface{}{
		"roleId": []string{"10", "11"},
	}
	jsonValue, _ := json.Marshal(values)
	RequestURL := "/v1/user/" + strconv.FormatInt(10, 10) + "/role"
	signature := sign.Signature("Lx1b8JoZoE", method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	repos := e.WithJSON(values).Expect().Status(http.StatusOK).JSON().Object()
	Convey("Subject: 根据用户分配角色\n", t, func() {
		So(repos.Value("_links").Array().First().Object().Value("href").String().Raw(), ShouldEqual, "/v1/user/10/role")
	})
}

func TestRoleByUserID(t *testing.T) {
	method := "GET"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	RequestURL := "/v1/user/10/role"
	signature := sign.Signature("Lx1b8JoZoE", method, nil, RequestURL, timestamp)
	tokin := tests.Tokin(t)
	e := tests.TestAPI(t, method, RequestURL, "app1", signature, timestamp, tokin)
	repos := e.Expect().Status(http.StatusOK).JSON().Object()
	Convey("Subject: 根据用户id获取角色\n", t, func() {
		So(repos.Value("Roles").Array().First().Object().Value("Code").Raw(), ShouldEqual, "10")
	})
}
