package core

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var app *echo.Echo

// 初始化echo实例
func NewEcho() *echo.Echo {
	app = echo.New()
	return app
}

//Hander 获取hander
func Hander() http.Handler {
	return app
}

// 开启服务
func Start(prot string) {
	app.Logger.Fatal(app.Start(prot))
}

// 打印请求异常信息
func Recover() {

	app.Use(middleware.Recover())
}

// 是否开启debug
func SetDebug(on bool) {
	app.Debug = on
}

// 获取debug状态
func Debug() bool {
	return app.Debug
}

// 打印请求信息
func Logger() {

	app.Use(middleware.Logger())
}

// 开启gzip压缩
func Gzip() {

	app.Use(middleware.Gzip())
}

// 设置Body大小
func BodyLimit(str string) {

	app.Use(middleware.BodyLimit(str))
}

// 自动添加末尾斜杠
func AddTrailingSlash() {

	app.Use(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
}

// 自动删除末尾斜杠
func RemoveTrailingSlash() {

	app.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
}
