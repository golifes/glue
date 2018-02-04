package auth

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/xwinie/glue/core"
)

//SysResourceController 用户资源
type SysResourceController struct {
}

//MenusByUserID 获取用户默认菜单
func (c *SysResourceController) MenusByUserID() func(echo.Context) error {
	return func(c echo.Context) error {
		userID := c.Param("userId")
		i64, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
		response := menuByUserIDService(i64)
		return c.JSON(response.StatusCode, response.Data)
	}
}

func (c *SysResourceController) ResourceByPage() func(echo.Context) error {
	return func(c echo.Context) error {
		pageSize := c.QueryParam("perPage")
		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil {
			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
		counts := findResourceCountByPageService()
		page := core.NewPaginator(c.Request(), pageSizeInt, counts)
		response := findResourceByPageService(page)
		return c.JSON(response.StatusCode, response.Data)
	}
}
func (c *SysResourceController) ResourceByCode() func(echo.Context) error {
	return func(c echo.Context) error {
		response := findResourceByCodeService(c.Param("code"))
		return c.JSON(response.StatusCode, response.Data)
	}
}
