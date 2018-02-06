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

//ClientByPage 分页获取
func (c *SysClientController) ClientByPage() func(echo.Context) error {
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
		var json SysClient
		err := c.Bind(&json)
		if err == nil {
			response := updateClientService(c.Param("id"), &json)
			return c.JSON(response.StatusCode, response.Data)
		}
		return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
	}
}

//Delete 删除数据
func (c *SysClientController) Delete() func(echo.Context) error {
	return func(c echo.Context) error {
		response := deleteClientService(c.Param("id"), 1)
		return c.JSON(response.StatusCode, response.Data)
	}
}

//Get 根据默认应用获取应用信息
func (c *SysClientController) Get() func(echo.Context) error {
	return func(c echo.Context) error {
		response := findClientByClientIDService(c.Request().Header.Get("appid"))
		return c.JSON(response.StatusCode, response.Data)
	}
}

//Post 创建应用
func (c *SysClientController) Post() func(echo.Context) error {
	return func(c echo.Context) error {
		var json SysClient
		err := c.Bind(&json)
		if err == nil {
			response := createClient(json)
			return c.JSON(response.StatusCode, response.Data)
		}
		return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常:"+err.Error()))

	}
}
