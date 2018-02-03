package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xwinie/glue/core"
)

//SysClientController 角色ctl
type SysClientController struct {
}

//RoleByPage 分页获取数据
func (c *SysClientController) RoleByPage() func(*gin.Context) {
	return func(c *gin.Context) {
		pageSize := c.Query("perPage")
		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "请求异常"})
		}
		counts := findClientCountByPageService()
		page := core.NewPaginator(c.Request, pageSizeInt, counts)
		response := findClientByPageService(page)
		c.JSON(response.StatusCode, response.Data)
	}
}

//Put 修改数据
func (c *SysClientController) Put() func(*gin.Context) {
	return func(c *gin.Context) {
		var json map[string]interface{}
		ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		if err := c.ShouldBindJSON(&json); err == nil {
			response := updateClientService(ID, json)
			c.JSON(response.StatusCode, response.Data)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "请求异常"})
		}
	}
}

//Delete 删除数据
func (c *SysClientController) Delete() func(*gin.Context) {
	return func(c *gin.Context) {
		ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		var lock int8
		if err := c.ShouldBindJSON(&lock); err == nil {
			response := deleteClientService(ID, lock)
			c.JSON(response.StatusCode, response.Data)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "请求异常"})
		}
	}
}

//Get
func (c *SysClientController) Get() func(*gin.Context) {
	return func(c *gin.Context) {
		response := findClientByClientIDService(c.GetHeader("appid"))
		c.JSON(response.StatusCode, response.Data)
	}
}
