package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//SysResourceController 用户资源
type SysResourceController struct {
}

//MenusByUserID 获取用户默认菜单
func (c *SysResourceController) MenusByUserID() func(*gin.Context) {
	return func(c *gin.Context) {
		userID := c.Param("userId")
		i64, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "请求异常"})
		}
		response := menuByUserIDService(i64)
		c.JSON(response.StatusCode, response.Data)
	}
}
