package auth

import (
	"github.com/labstack/echo"
)

//SysDictController ctl
type SysDictController struct {
}

//Get
func (c *SysDictController) Get() func(echo.Context) error {
	return func(c echo.Context) error {
		response := findDictByTypeService(c.Param("type"))
		return c.JSON(response.StatusCode, response.Data)
	}
}
