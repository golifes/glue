package auth

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/xwinie/glue/core"
)

//SysRoleController 角色ctl
type SysRoleController struct {
}

//UserByPage 分页获取数据
func (c *SysRoleController) RoleByPage() func(echo.Context) error {
	return func(c echo.Context) error {
		pageSize := c.QueryParam("perPage")
		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil {
			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
		counts := findRoleCountByPageService()
		page := core.NewPaginator(c.Request(), pageSizeInt, counts)
		response := findRoleByPageService(page)
		return c.JSON(response.StatusCode, response.Data)
	}
}

//Post
func (c *SysRoleController) Post() func(echo.Context) error {
	return func(c echo.Context) error {
		var json SysRole
		if err := c.Bind(&json); err == nil {
			response := createRole(json)
			return c.JSON(response.StatusCode, response.Data)
		} else {
			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
	}
}

//Put 修改数据
func (c *SysRoleController) Put() func(echo.Context) error {
	return func(c echo.Context) error {
		var json map[string]interface{}
		ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		if err := c.Bind(&json); err == nil {
			response := updateRoleService(ID, json)
			return c.JSON(response.StatusCode, response.Data)
		} else {
			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
	}
}

//Delete 删除数据
func (c *SysRoleController) Delete() func(echo.Context) error {
	return func(c echo.Context) error {
		ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		response := deleteRoleService(ID)
		return c.JSON(response.StatusCode, response.Data)
	}
}

//Get
func (c *SysRoleController) Get() func(echo.Context) error {
	return func(c echo.Context) error {
		response := findRoleByCodeService(c.Param("code"))
		return c.JSON(response.StatusCode, response.Data)
	}
}

//GetResourceByRoleID
func (c *SysRoleController) GetResourceByRoleID() func(echo.Context) error {
	return func(c echo.Context) error {

		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		response := findResourceByRoleIdService(id)
		return c.JSON(response.StatusCode, response.Data)
	}
}

//RoleAllotResource
func (c *SysRoleController) RoleAllotResource() func(echo.Context) error {
	return func(c echo.Context) error {
		type json struct {
			ResourceId []int64
		}
		var d json
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err = c.Bind(&d); err == nil {
			response := roleAllotResource(id, d.ResourceId)
			return c.JSON(response.StatusCode, response.Data)
		} else {

			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
	}
}
