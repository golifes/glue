package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xwinie/glue/core"
)

//SysRoleController 角色ctl
type SysRoleController struct {
}

//UserByPage 分页获取数据
func (c *SysRoleController) RoleByPage() func(*gin.Context) {
	return func(c *gin.Context) {
		pageSize := c.Query("perPage")
		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "请求异常"})
		}
		counts := findRoleCountByPageService()
		page := core.NewPaginator(c.Request, pageSizeInt, counts)
		response := findRoleByPageService(page)
		c.JSON(response.StatusCode, response.Data)
	}
}

//Post
func (c *SysRoleController) Post() func(*gin.Context) {
	return func(c *gin.Context) {
		var json SysRole
		if err := c.ShouldBindJSON(&json); err == nil {
			response := createRole(json)
			c.JSON(response.StatusCode, response.Data)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "请求异常"})
		}
	}
}

//Put 修改数据
func (c *SysRoleController) Put() func(*gin.Context) {
	return func(c *gin.Context) {
		var json map[string]interface{}
		ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		if err := c.ShouldBindJSON(&json); err == nil {
			response := updateRoleService(ID, json)
			c.JSON(response.StatusCode, response.Data)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "请求异常"})
		}
	}
}

//Delete 删除数据
func (c *SysRoleController) Delete() func(*gin.Context) {
	return func(c *gin.Context) {
		ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		response := deleteRoleService(ID)
		c.JSON(response.StatusCode, response.Data)
	}
}

//Get
func (c *SysRoleController) Get() func(*gin.Context) {
	return func(c *gin.Context) {
		response := findRoleByCodeService(c.Param("code"))
		c.JSON(response.StatusCode, response.Data)
	}
}

//GetResourceByRoleID
func (c *SysRoleController) GetResourceByRoleID() func(*gin.Context) {
	return func(c *gin.Context) {

		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		response := findResourceByRoleIdService(id)
		c.JSON(response.StatusCode, response.Data)
	}
}

//RoleAllotResource
func (c *SysRoleController) RoleAllotResource() func(*gin.Context) {
	return func(c *gin.Context) {
		var json []int64
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		if err := c.ShouldBindJSON(&json); err == nil {
			response := roleAllotResource(id, json)
			c.JSON(response.StatusCode, response.Data)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "请求异常"})
		}
	}
}
