package glue

import (
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"github.com/xwinie/glue/app/auth"

	"github.com/gin-gonic/gin"
	"github.com/xwinie/glue/lib/db"
	"github.com/xwinie/glue/lib/middleware/casbin"
	"github.com/xwinie/glue/lib/middleware/sign"
)

//Engine 获取engine
func Engine() *gin.Engine {
	dbconfig := db.Config{}
	dbconfig.DbHost = "192.168.3.253"
	dbconfig.DbName = "yunhao"
	dbconfig.DbUser = "yunhao"
	dbconfig.DbPassword = "1qazZAQ!"
	dbconfig.DbPort = "3306"
	dbconfig.DbType = "mysql"
	dbconfig.DbCharset = "utf8mb4"
	dbconfig.DbPath = []string{"/Users/bobo/go/src/github.com/xwinie/glue/app"}

	s := gin.Default()
	err := db.Connect(dbconfig)
	if err != nil {
		log.Fatal("init db error:", err.Error())
	}
	s.Use(sign.New(sign.Config{F: getAppSecret, Timeout: 100}))

	s.GET("/a", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s.POST("/v1/login", new(auth.LoginController).Post())

	s.POST("/form_post", func(c *gin.Context) {
		type Login struct {
			Message string
			Nick    string
		}
		var json Login
		if err := c.ShouldBindJSON(&json); err == nil {
			c.JSON(200, gin.H{
				"status":  "posted",
				"message": json.Message,
				"nick":    json.Nick,
			})
		}

	})
	return s
}

//Hander 获取hander
func Hander() http.Handler {
	return Engine()
}

func getAppSecret(appid string) string {
	// get appsecret by appid
	client, _ := auth.GetClientService(appid)
	return client.Secret
}

func getRestAutz(authorization, appid string) ([]casbin.Permission, string, error) {
	client, _ := auth.GetClientService(appid)
	type Claims struct {
		Exp      float64
		Iat      int64
		Issuer   string
		UserId   int64
		userType int8
		Account  string
		UserName string
		Role     []int64
	}
	token, err := casbin.ParseToken(authorization, client.VerifySecret)
	if err != nil {
		return nil, "", err
	}
	claimsMap, ok := token.Claims.(jwt.MapClaims)
	var claims Claims
	err = mapstructure.Decode(claimsMap, &claims)
	if !ok || err != nil {
		return nil, "", err
	}
	permiss, _, err := auth.PermissionByMultiRole(claims.Role, 1)
	return permiss, string(claims.UserId), err
}
