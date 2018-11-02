package crm

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/xwinie/glue/core"
)

//CustomerController ctl
type CustomerController struct {
}

//ClientByPage 分页获取
func (c *CustomerController) CustomerByPage() func(echo.Context) error {
	return func(c echo.Context) error {
		pageSize := c.QueryParam("perPage")
		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil {
			return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
		}
		counts := findCustomerCountByPageService()
		page := core.NewPaginator(c.Request(), pageSizeInt, counts)
		response := findCustomerByPageService(page)
		return c.JSON(response.StatusCode, response.Data)
	}
}

//Put 修改数据
func (c *CustomerController) Put() func(echo.Context) error {
	return func(c echo.Context) error {
		var json Customer
		err := c.Bind(&json)
		if err == nil {
			response := updateCustomerService(c.Param("id"), &json)
			return c.JSON(response.StatusCode, response.Data)
		}
		return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常"))
	}
}

//Delete 删除数据
func (c *CustomerController) Delete() func(echo.Context) error {
	return func(c echo.Context) error {
		i64, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		response := deleteCustomerService(i64)
		return c.JSON(response.StatusCode, response.Data)
	}
}

//Get
func (c *CustomerController) Get() func(echo.Context) error {
	return func(c echo.Context) error {
		response := findCustomerByCustomerIdService(c.Param("id"))
		return c.JSON(response.StatusCode, response.Data)
	}
}

//Post
func (c *CustomerController) Post() func(echo.Context) error {
	return func(c echo.Context) error {
		var json Customer
		err := c.Bind(&json)
		if err == nil {
			response := createCustomer(json)
			return c.JSON(response.StatusCode, response.Data)
		}
		return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常:"+err.Error()))

	}
}

//PostTag
func (c *CustomerController) PostTag() func(echo.Context) error {
	return func(c echo.Context) error {
		type json struct {
			TagName []string
		}
		var d json
		err := c.Bind(&d)
		if err == nil {
			response := createCustomerTagBycustomerId(c.Param("id"), d.TagName)
			return c.JSON(response.StatusCode, response.Data)
		}
		return c.JSON(http.StatusBadRequest, core.BuildEntity(http.StatusBadRequest, "请求异常:"+err.Error()))
	}
}

//GetTag
func (c *CustomerController) GetTag() func(echo.Context) error {
	return func(c echo.Context) error {
		response := findCustomerTagByCustomerIdService(c.Param("id"))
		return c.JSON(response.StatusCode, response.Data)
	}
}
