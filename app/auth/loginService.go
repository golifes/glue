package auth

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/xwinie/glue/lib/response"
)

//LoginData 登录from请求数据字段
type loginData struct {
	UserName string
	Password string
}

func loginService(loginData *loginData, appID string) (responseEntity response.ResponseEntity) {
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if r := recover(); r != nil {
			responseEntity.BuildError(response.BuildEntity(QueryError, getMsg(QueryError)))
			return
		}
	}()

	user, err := findUserAllColums(loginData.UserName)

	if &user.ID == nil && err != nil {
		return *responseEntity.BuildError(response.BuildEntity(NotFoundUser, getMsg(NotFoundUser)))
	}
	if !user.CheckEqualPassword(loginData.Password) {
		return *responseEntity.NewBuild(http.StatusUnauthorized, response.BuildEntity(Unauthorized, getMsg(Unauthorized)))
	}

	roleID, err1 := findRoleIDByUserID(user.ID)
	if err1 != nil {
		return *responseEntity.BuildError(response.BuildEntity(NotFoundUserRole, getMsg(NotFoundUserRole)))

	}
	client, err := GetClientService(appID)
	if err != nil {
		return *responseEntity.BuildError(response.BuildEntity(GenerateTokenError, getMsg(GenerateTokenError)))
	}
	exp := time.Now().Add(time.Hour * time.Duration(1)).Unix()
	token, tokenErr := getToken(appID, client.VerifySecret, user, roleID, exp)

	if tokenErr != nil {
		return *responseEntity.BuildError(response.BuildEntity(GenerateTokenError, getMsg(GenerateTokenError)))
	}
	type data struct {
		Account string
		Name    string
		Token   string
		Exp     int64
		ID      int64
	}
	d := &data{user.Account, user.Name, token, exp, user.ID}
	return *responseEntity.BuildPostAndPut(d)
}

func getToken(appID string, key string, user SysUser, userRoleID []int64, exp int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = time.Now().Unix()
	claims["Issuer"] = appID
	claims["userId"] = user.ID
	claims["account"] = user.Account
	claims["userName"] = user.Name
	claims["userType"] = user.UserType
	claims["role"] = userRoleID
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(key))
	return tokenString, err
}
