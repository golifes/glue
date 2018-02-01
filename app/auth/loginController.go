package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//LoginController 用户登录结构
type LoginController struct {
}

//Post 用户登录post
func (c *LoginController) Post() func(*gin.Context) {
	return func(c *gin.Context) {
		var loginData loginData
		if err := c.ShouldBindJSON(&loginData); err == nil {
			response := loginService(&loginData, c.GetHeader("appid"))
			c.JSON(response.StatusCode, response.Data)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "请求异常"})
		}
	}
}
