package auth

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	. "github.com/xwinie/glue/lib/response"
)

//LoginData 登录from请求数据字段
type LoginData struct {
	UserName string
	Password string
}

func loginService(loginData *LoginData, appID string) (responseEntity ResponseEntity) {
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if r := recover(); r != nil {
			responseEntity.BuildError(BuildEntity(QueryError, GetMsg(QueryError)))
			return
		}
	}()

	user, err := findUserAllColums(loginData.UserName)

	if &user.Id == nil && err != nil {
		return *responseEntity.BuildError(BuildEntity(NotFoundUser, GetMsg(NotFoundUser)))
	}
	if !user.CheckEqualPassword(loginData.Password) {
		return *responseEntity.NewBuild(http.StatusUnauthorized, BuildEntity(Unauthorized, GetMsg(Unauthorized)))
	}

	roleID, err1 := findRoleIDByUserID(user.Id)
	if err1 != nil {
		return *responseEntity.BuildError(BuildEntity(NotFoundUserRole, GetMsg(NotFoundUserRole)))

	}
	client, err := GetClientService(appID)
	if err != nil {
		return *responseEntity.BuildError(BuildEntity(GenerateTokenError, GetMsg(GenerateTokenError)))
	}
	exp := time.Now().Add(time.Hour * time.Duration(1)).Unix()
	token, tokenErr := getToken(appID, client.VerifySecret, user, roleID, exp)

	if tokenErr != nil {
		return *responseEntity.BuildError(BuildEntity(GenerateTokenError, GetMsg(GenerateTokenError)))
	} else {
		type data struct {
			Account string
			Name    string
			Token   string
			Exp     int64
			Id      int64
		}
		d := &data{user.Account, user.Name, token, exp, user.Id}
		return *responseEntity.BuildPostAndPut(d)
	}

}

func getToken(appID string, key string, user SysUser, userRoleID []int64, exp int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = time.Now().Unix()
	claims["Issuer"] = appID
	claims["userId"] = user.Id
	claims["account"] = user.Account
	claims["userName"] = user.Name
	claims["userType"] = user.UserType
	claims["role"] = userRoleID
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(key))
	return tokenString, err
}
