package notFound

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//NotFound 没有找到的页面
func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
	}
}
