package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xwinie/glue/lib/utils"
)

//SysUserController 用户项目ctl
type SysUserController struct {
}

//UserByPage 分页获取数据
func (c *SysUserController) UserByPage() func(*gin.Context) {
	return func(c *gin.Context) {
		pageSize := c.Param("perPage")
		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "请求异常"})
		}
		counts := FindUserCountByPageService()
		page := utils.NewPaginator(c.Request, pageSizeInt, counts)
		response := FindUserByPageService(page)
		c.JSON(response.StatusCode, response.Data)
	}
}
