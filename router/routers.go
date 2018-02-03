package router

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/xwinie/glue/app/auth"
	"github.com/xwinie/glue/core/middleware/casbin"
	"github.com/xwinie/glue/core/middleware/sign"
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
			userCtl := new(auth.SysUserController)
			user.GET("", userCtl.UserByPage())
			user.GET("/:account", userCtl.Get())
			user.POST("", userCtl.Post())
			user.PUT("/:id", userCtl.Put())
			user.DELETE("/:id", userCtl.Delete())
		}
		role := v1.Group("/user")
		{
			ctl := new(auth.SysUserController)
			role.GET("", ctl.RoleByPage())
			role.GET("/:code", ctl.Get())
			role.POST("", ctl.Post())
			role.PUT("/:id", ctl.Put())
			role.DELETE("/:id", ctl.Delete())
		}
		v1.GET("/menus/:userId", new(auth.SysResourceController).MenusByUserID())
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
	var permiss []casbin.Permission

	permiss, err = auth.PermissionByMultiRole(claims.Role, 0)
	return permiss, claims.Account, err
}

func getOpenRestAutz() ([]casbin.Permission, error) {
	permiss, err := auth.OpenPermission()
	return permiss, err
}
