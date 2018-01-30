package glue

import "github.com/gin-gonic/gin"

//Engine 获取engine
func Engine() *gin.Engine {
	s := gin.Default()
	return s
}
