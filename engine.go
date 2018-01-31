package glue

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xwinie/glue/lib/middleware/sign"
)

//Engine 获取engine
func Engine() *gin.Engine {
	s := gin.Default()

	s.Use(sign.New(sign.Config{getAppSecret, 100}))
	s.GET("/a", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return s
}

//Hander 获取hander
func Hander() http.Handler {
	return Engine()
}

func getAppSecret(appid string) string {
	// get appsecret by appid
	// maybe store in configure, maybe in database
	return "1"
}
