package casbin

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/casbin/casbin"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//Config casbin需要的配置
type Config struct {
	OpenF func() ([]Permission, error)
	F     func(string, string) ([]Permission, string, error) //基于token的Authorization来进行数据获取 第一个是Authorization 第二个是appid
}

// Permission 需要的权限结果集合
type Permission struct {
	Code   string
	Action string
	Method string
}

const modal = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)
`

//Authorizer 权限结构
type Authorizer struct {
	enforcer *casbin.Enforcer
}

//RestAuth rest权限校验
func RestAuth(c Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			req := ctx.Request()
			defer func() {
				// recover from panic if one occured. Set err to nil otherwise.
				if r := recover(); r != nil {
					ctx.JSON(http.StatusForbidden, r)
					return
				}
			}()

			var e = casbin.NewEnforcer(casbin.NewModel(modal))
			a := &Authorizer{enforcer: e}
			//判断是否开放接口
			openPermiss, err := c.OpenF()

			if err != nil {
				ctx.JSON(http.StatusForbidden, "Auth Fail"+err.Error())
			}
			for _, v := range openPermiss {
				e.AddPermissionForUser(req.Header.Get("appid"), v.Action, v.Method)
			}
			if a.checkPermission(req.Header.Get("appid"), req.Method, req.URL.Path) {

				return next(ctx)
			}

			if req.Header.Get("Authorization") == "" {
				ctx.JSON(http.StatusForbidden, "miss Authorization header")

			}
			//权限校验
			permissions, userID, err1 := c.F(req.Header.Get("Authorization"), req.Header.Get("appid"))
			if err1 != nil {
				ctx.JSON(http.StatusForbidden, "Auth Fail"+err.Error())
			}
			for _, v := range permissions {
				e.AddPermissionForUser(userID, v.Action, v.Method)
			}
			if !a.checkPermission(userID, req.Method, req.URL.Path) {
				ctx.JSON(http.StatusForbidden, "Auth Fail")
			}
			return next(ctx)
		}
	}

}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *Authorizer) checkPermission(user, method, path string) bool {
	return a.enforcer.Enforce(user, path, method)
}

//ParseToken 解析token
func ParseToken(authString, secret string) (*jwt.Token, error) {
	if strings.Split(authString, " ")[1] == "" {
		return nil, errors.New("AuthString invalid,Token:" + authString)
	}

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		return nil, errors.New("AuthString invalid,Token:" + authString)
	}
	tokenString := kv[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("That's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errors.New("Token is either expired or not active yet")
			} else {
				return nil, errors.New("Couldn‘t handle this token")
			}
		} else {
			return nil, errors.New("Parse token is error")
		}
	}
	if !token.Valid {
		return nil, errors.New("Token invalid:" + authString)
	}

	return token, nil
}
