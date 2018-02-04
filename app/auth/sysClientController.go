package auth

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/xwinie/glue/core"
)

//SysClientController 角色ctl
type SysClientController struct {
}

//RoleByPage 分页获取数据
func (c *SysClientController) RoleByPage() func(echo.Context) error {
	return func(c echo.Context) error {
		pageSize := c.QueryParam("perPage")
		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil {
			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
		counts := findClientCountByPageService()
		page := core.NewPaginator(c.Request(), pageSizeInt, counts)
		response := findClientByPageService(page)
		return c.JSON(response.StatusCode, response.Data)
	}
}

//Put 修改数据
func (c *SysClientController) Put() func(echo.Context) error {
	return func(c echo.Context) error {
		var json map[string]interface{}
		ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		if err := c.Bind(&json); err == nil {
			response := updateClientService(ID, json)
			return c.JSON(response.StatusCode, response.Data)
		} else {
			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
	}
}

//Delete 删除数据
func (c *SysClientController) Delete() func(echo.Context) error {
	return func(c echo.Context) error {
		ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		var lock int8
		if err := c.Bind(&lock); err == nil {
			response := deleteClientService(ID, lock)
			return c.JSON(response.StatusCode, response.Data)
		} else {
			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
	}
}

//Get
func (c *SysClientController) Get() func(echo.Context) error {
	return func(c echo.Context) error {
		response := findClientByClientIDService(c.Request().Header.Get("appid"))
		return c.JSON(response.StatusCode, response.Data)
	}
}
