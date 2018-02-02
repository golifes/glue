package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/gavv/httpexpect"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xwinie/glue/lib/db"
	"github.com/xwinie/glue/lib/middleware/sign"
	"github.com/xwinie/glue/lib/utils"
	"github.com/xwinie/glue/migrate"
	"github.com/xwinie/glue/router"
)

//Engine 获取engine
func app() *gin.Engine {
	dbconfig := db.Config{}
	dbconfig.DbType = "sqlite3"
	dbconfig.DbCharset = "utf8mb4"
	dbconfig.DbPath = []string{"/Users/bobo/go/src/github.com/xwinie/glue/app"}

	s := gin.Default()
	err := db.Connect(dbconfig)
	if err != nil {
		log.Fatal("init db error:", err.Error())
	}

	router.Routers(s)
	o := db.New()
	migrate.Migrate(o)
	return s
}

//Hander 获取hander
func hander() http.Handler {
	return app()
}

func testHander(t *testing.T) *httpexpect.Expect {
	handler := hander()
	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
	return e
}

//TestAPI 测试基础方法
func TestAPI(t *testing.T, method, path, appid, signature, timestamp, tokenString string) *httpexpect.Request {
	e := testHander(t)

	headers := make(map[string]string, 4)
	if tokenString != "" {
		headers["Authorization"] = "Bearer " + tokenString
	}
	headers["appid"] = appid
	headers["signature"] = signature
	headers["timestamp"] = timestamp

	return e.Request(method, path).WithHeaders(headers)
}

//Tokin 获取token
func Tokin(t *testing.T) string {
	method := "POST"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	values := map[string]interface{}{
		"UserName": "12345",
		"Password": utils.Md5(utils.Sha1("12345") + utils.Sha1("Password")),
	}
	jsonValue, _ := json.Marshal(values)
	RequestURL := "/v1/login"
	signature := sign.Signature("Lx1b8JoZoE", method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	e := TestAPI(t, method, RequestURL, "app1", signature, timestamp, "")
	r := e.WithJSON(values).Expect().Status(http.StatusCreated).JSON().Object()
	fmt.Println(1111, r)
	return r.Value("Token").String().Raw()
}
