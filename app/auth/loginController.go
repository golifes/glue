package auth

import (
	"github.com/gin-gonic/gin"
)

//LoginController 用户登录结构
type LoginController struct {
	gin.Context
}

//Post 用户post
func (c *LoginController) Post() {
	var loginData LoginData
	if err := c.ShouldBindJSON(&loginData); err == nil {
		response := loginService(&loginData, c.GetHeader("appid"))
		c.JSON(response.StatusCode, response.Data)
	}

}
