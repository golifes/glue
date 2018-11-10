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
		err := c.Bind(&json)
		if err == nil {
			response := createUser(json)
			return c.JSON(response.StatusCode, response.Data)
		}
		return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常:"+err.Error()))

	}
}

//Put 修改数据
func (c *SysUserController) Put() func(echo.Context) error {
	return func(c echo.Context) error {
		var json SysUser
		err := c.Bind(&json)
		if err == nil {
			response := updateUserService(c.Param("id"), &json)
			return c.JSON(response.StatusCode, response.Data)
		}
		return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常:"+err.Error()))

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
		c.SetParamNames("account") //put 和get same router,这样两种解决方法，1是现在这种重新设置param，2是用param("id"),即重复的param
		if account := c.Param("account"); account != "" {
			response := findUserByAccountService(account)
			return c.JSON(response.StatusCode, response.Data)
		}
		return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
	}
}

//GetRoleByUserID 根据用户获取角色
func (c *SysUserController) GetRoleByUserID() func(echo.Context) error {
	return func(c echo.Context) error {
		response := findRoleByUserIDService(c.Param("id"))
		return c.JSON(response.StatusCode, response.Data)
	}
}

//UserAllotRole 用户分配角色
func (c *SysUserController) UserAllotRole() func(echo.Context) error {
	return func(c echo.Context) error {
		type json struct {
			RoleID []string `query:"RoleId"`
		}
		var d json
		err := c.Bind(&d)
		if err == nil {
			response := userAllotRole(c.Param("id"), d.RoleID)
			return c.JSON(response.StatusCode, response.Data)
		}
		return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
	}
}
