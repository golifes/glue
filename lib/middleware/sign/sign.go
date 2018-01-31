package sign

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//Config 需要的结构配置 参考beego的apiauth
type Config struct {
	F       func(string) string
	Timeout int
}

//New 签名认证
func New(conf Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("appid") == "" {
			ctx.JSON(http.StatusForbidden, "miss appid header")
		}
		appsecret := conf.F(ctx.GetHeader("appid"))
		if appsecret == "" {
			ctx.JSON(http.StatusForbidden, "not exist this appid")
		}
		clientSignature := ctx.GetHeader("signature")
		if clientSignature == "" {
			ctx.JSON(http.StatusForbidden, "miss signature header")
		}
		if ctx.GetHeader("timestamp") == "" {
			ctx.JSON(http.StatusForbidden, "miss timestamp header")
		}
		u, err := time.Parse("2006-01-02 15:04:05", ctx.GetHeader("timestamp"))
		if err != nil {
			ctx.JSON(http.StatusForbidden, "timestamp format is error, should 2006-01-02 15:04:05")
		}
		t := time.Now()
		if t.Sub(u).Seconds() > float64(conf.Timeout) {
			ctx.JSON(http.StatusForbidden, "timeout! the request time is long ago, please try again")
		}
		var requestURL string
		var body []byte
		if ctx.Request.Method == "GET" {
			requestURL = ctx.Request.RequestURI
		} else {
			requestURL = ctx.Request.URL.Path
			body, err = ctx.GetRawData()
			newBody := ioutil.NopCloser(bytes.NewBuffer(body))
			ctx.Request.Body = newBody
		}
		serviceSignature := Signature(appsecret, ctx.Request.Method, body, requestURL, ctx.GetHeader("timestamp"))
		// fmt.Println("serviceSignature:", serviceSignature)
		// fmt.Println("clientSignature:", clientSignature)
		if clientSignature != serviceSignature {
			ctx.JSON(http.StatusForbidden, "Signature Failed")
		}
		ctx.Next()
	}
}

// Signature used to generate signature with the appsecret/method/params/RequestURI
func Signature(appSecret, method string, body []byte, RequestURL string, timestamp string) (result string) {
	stringToSign := fmt.Sprintf("%v\n%v\n%v\n%v\n", method, string(body), RequestURL, timestamp)
	// fmt.Println(1111, stringToSign)
	sha256 := sha256.New
	hash := hmac.New(sha256, []byte(appSecret))
	hash.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
