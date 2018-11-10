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

//RoleByPage 分页获取数据
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

//Post 创建角色
func (c *SysRoleController) Post() func(echo.Context) error {
	return func(c echo.Context) error {
		var json SysRole
		if err := c.Bind(&json); err == nil {
			response := createRole(json)
			return c.JSON(response.StatusCode, response.Data)
		}
		return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))

	}
}

//Put 修改数据
func (c *SysRoleController) Put() func(echo.Context) error {
	return func(c echo.Context) error {
		var json SysRole
		if err := c.Bind(&json); err == nil {
			response := updateRoleService(c.Param("id"), json)
			return c.JSON(response.StatusCode, response.Data)
		}
		return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
	}
}

//Delete 删除数据
func (c *SysRoleController) Delete() func(echo.Context) error {
	return func(c echo.Context) error {
		response := deleteRoleService(c.Param("id"))
		return c.JSON(response.StatusCode, response.Data)
	}
}

//Get 根据code获取角色
func (c *SysRoleController) Get() func(echo.Context) error {
	return func(c echo.Context) error {
		if id := c.Param("id"); id != "" {
			response := findRoleByCodeService(id)
			return c.JSON(response.StatusCode, response.Data)
		}
		return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
	}
}

//GetResourceByRoleID 根据角色获取资源
func (c *SysRoleController) GetResourceByRoleID() func(echo.Context) error {
	return func(c echo.Context) error {
		response := findResourceByRoleIDService(c.Param("id"))
		return c.JSON(response.StatusCode, response.Data)
	}
}

//RoleAllotResource  角色分配资源
func (c *SysRoleController) RoleAllotResource() func(echo.Context) error {
	return func(c echo.Context) error {
		type json struct {
			ResourceID []string `query:"ResourceId"`
		}
		var d json
		err := c.Bind(&d)
		if err == nil {
			response := roleAllotResource(c.Param("id"), d.ResourceID)
			return c.JSON(response.StatusCode, response.Data)
		}

		return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常:"+err.Error()))

	}
}
