package auth

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/xwinie/glue/core"
)

//SysUserController 用户ctl
type SysUserController struct {
}

//UserByPage 分页获取数据
func (c *SysUserController) UserByPage() func(echo.Context) error {
	return func(c echo.Context) error {
		pageSize := c.QueryParam("perPage")
		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil {
			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
		counts := findUserCountByPageService()
		page := core.NewPaginator(c.Request(), pageSizeInt, counts)
		response := findUserByPageService(page)
		return c.JSON(response.StatusCode, response.Data)
	}
}

//Post 创建用户
func (c *SysUserController) Post() func(echo.Context) error {
	return func(c echo.Context) error {
		var json SysUser
		if err := c.Bind(&json); err == nil {
			response := createUser(json)
			return c.JSON(response.StatusCode, response.Data)
		} else {
			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
	}
}

//Put 修改数据
func (c *SysUserController) Put() func(echo.Context) error {
	return func(c echo.Context) error {
		var json map[string]interface{}
		ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		if err := c.Bind(&json); err == nil {
			response := updateUserService(ID, json)
			return c.JSON(response.StatusCode, response.Data)
		} else {
			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
	}
}

//Delete 删除数据
func (c *SysUserController) Delete() func(echo.Context) error {
	return func(c echo.Context) error {
		ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		response := deleteUserService(ID)
		return c.JSON(response.StatusCode, response.Data)
	}
}

//Get 根据账号获取数据
func (c *SysUserController) Get() func(echo.Context) error {
	return func(c echo.Context) error {
		response := findUserByAccountService(c.Param("account"))
		return c.JSON(response.StatusCode, response.Data)
	}
}

func (c *SysUserController) GetRoleByUserID() func(echo.Context) error {
	return func(c echo.Context) error {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		response := findRoleByUserIDService(id)
		return c.JSON(response.StatusCode, response.Data)
	}
}

func (c *SysUserController) UserAllotRole() func(echo.Context) error {
	return func(c echo.Context) error {
		type json struct {
			RoleId []int64
		}
		var d json
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		err = c.Bind(&d)
		if err == nil {
			response := userAllotRole(id, d.RoleId)
			return c.JSON(response.StatusCode, response.Data)
		} else {
			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
	}
}
