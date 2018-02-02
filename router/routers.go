package router

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/xwinie/glue/app/auth"
	"github.com/xwinie/glue/lib/middleware/casbin"
	"github.com/xwinie/glue/lib/middleware/sign"
)

//Routers 系统路由
func Routers(s *gin.Engine) {
	s.Use(sign.New(sign.Config{F: getAppSecret, Timeout: 100}))
	s.Use(casbin.RestAuth(casbin.Config{OpenF: getOpenRestAutz, F: getRestAutz}))
	v1 := s.Group("/v1")
	{
		v1.POST("/login", new(auth.LoginController).Post())
		user := v1.Group("/user")
		{
			user.GET("/", new(auth.SysResourceController).MenusByUserID())
		}
	}
	s.GET("/a", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

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
		UserID   int64
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
	var permiss []casbin.Permission

	permiss, _, err = auth.PermissionByMultiRole(claims.Role, 0)
	return permiss, string(claims.UserID), err
}

func getOpenRestAutz() ([]casbin.Permission, error) {
	permiss, err := auth.OpenPermission()
	return permiss, err
}
