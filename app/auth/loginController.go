package auth

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/xwinie/glue/core"
)

//LoginController 用户登录结构
type LoginController struct {
}

//Post 用户登录post
func (c *LoginController) Post() func(echo.Context) error {
	return func(c echo.Context) error {
		login := new(loginData)
		if err := c.Bind(login); err == nil {
			response := loginService(login, c.Request().Header.Get("appid"))
			return c.JSON(response.StatusCode, response.Data)
		} else {
			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
	}
}
