package casbin

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/casbin/casbin"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Config casbin需要的配置
type Config struct {
	model string
	F     func(string, string) ([]Permission, string, error) //基于token的Authorization来进行数据获取 第一个是Authorization 第二个是appid
	Open  bool                                               // 0代表非开放 1代表开放接口
}

// Permission 需要的权限结果集合
type Permission struct {
	ID     string
	Action string
	Method string
}

//Authorizer 权限结构
type Authorizer struct {
	enforcer *casbin.Enforcer
}

//RestAuth rest权限校验
func RestAuth(c Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if c.model == "" {
			c.model = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)
`
		}
		var e = casbin.NewEnforcer(casbin.NewModel(c.model))
		a := &Authorizer{enforcer: e}
		permissions, userID, err := c.F(ctx.GetHeader("Authorization"), ctx.GetHeader("appid"))
		if err != nil {
			ctx.JSON(http.StatusForbidden, "Auth Fail"+err.Error())
		}
		for _, v := range permissions {
			e.AddPermissionForUser(userID, v.Action, v.Method)
		}
		if !c.Open && !a.checkPermission(userID, ctx.Request.Method, ctx.Request.URL.Path) {
			ctx.JSON(http.StatusForbidden, "Auth Fail")
		} else if c.Open && a.checkPermission(ctx.GetHeader("appid"), ctx.Request.Method, ctx.Request.URL.Path) {
			ctx.Next()
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
