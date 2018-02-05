package router

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
	"github.com/xwinie/glue/app/auth"
	"github.com/xwinie/glue/core/middleware/casbin"
	"github.com/xwinie/glue/core/middleware/sign"
)

//Routers 系统路由
func Routers(s *echo.Echo) {
	s.Use(sign.New(sign.Config{F: getAppSecret, Timeout: 100}))
	s.Use(casbin.RestAuth(casbin.Config{OpenF: getOpenRestAutz, F: getRestAutz}))

	v1 := s.Group("/v1")
	v1.POST("/login", new(auth.LoginController).Post())
	user := v1.Group("/user")
	usrCtl := new(auth.SysUserController)
	user.POST("", usrCtl.Post())
	user.PUT("/:id", usrCtl.Put())
	user.GET("", usrCtl.UserByPage())
	user.GET("/:account", usrCtl.Get())
	user.DELETE("/:id", usrCtl.Delete())
	user.GET("/:id/role", usrCtl.GetRoleByUserID())
	user.POST("/:id/role", usrCtl.UserAllotRole())
	role := v1.Group("/role")
	roleCtl := new(auth.SysRoleController)
	role.POST("", roleCtl.Post())
	role.PUT("/:id", roleCtl.Put())
	role.GET("", roleCtl.RoleByPage())
	role.GET("/:code", roleCtl.Get())
	role.DELETE("/:id", roleCtl.Delete())
	role.GET("/:id/resource", roleCtl.GetResourceByRoleID())
	role.POST("/:id/resource", roleCtl.RoleAllotResource())
	resource := v1.Group("/resource")
	resourceCtl := new(auth.SysResourceController)
	resource.GET("", resourceCtl.ResourceByPage())
	resource.GET("/:code", resourceCtl.ResourceByCode())
	v1.GET("/menus/:userId", resourceCtl.MenusByUserID())
	client := v1.Group("/client")
	clientCtl := new(auth.SysClientController)
	client.GET("", clientCtl.ClientByPage())
	client.GET("/default", clientCtl.Get())
	client.PUT("/:id", clientCtl.Put())
	client.DELETE("/:id", clientCtl.Delete())
	client.POST("", clientCtl.Post())
	s.GET("/a", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
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
		UserId   string
		userType int8
		Account  string
		UserName string
		Role     []string
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
