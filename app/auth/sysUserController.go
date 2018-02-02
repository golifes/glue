package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xwinie/glue/core"
)

//SysUserController 用户项目ctl
type SysUserController struct {
}

//UserByPage 分页获取数据
func (c *SysUserController) UserByPage() func(*gin.Context) {
	return func(c *gin.Context) {
		pageSize := c.Query("perPage")
		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "请求异常"})
		}
		counts := findUserCountByPageService()
		page := core.NewPaginator(c.Request, pageSizeInt, counts)
		response := findUserByPageService(page)
		c.JSON(response.StatusCode, response.Data)
	}
}

//Post 创建用户
func (c *SysUserController) Post() func(*gin.Context) {
	return func(c *gin.Context) {
		var json SysUser
		if err := c.ShouldBindJSON(&json); err == nil {
			response := createUser(json)
			c.JSON(response.StatusCode, response.Data)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}

//Put 修改数据
func (c *SysUserController) Put() func(*gin.Context) {
	return func(c *gin.Context) {
		var json map[string]interface{}
		ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		if err := c.ShouldBindJSON(&json); err == nil {
			response := updateUserService(ID, json)
			c.JSON(response.StatusCode, response.Data)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}

//Delete 删除数据
func (c *SysUserController) Delete() func(*gin.Context) {
	return func(c *gin.Context) {
		ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		response := deleteUserService(ID)
		c.JSON(response.StatusCode, response.Data)
	}
}
