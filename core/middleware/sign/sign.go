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

	"github.com/labstack/echo"
)

//Config 需要的结构配置 参考beego的apiauth
type Config struct {
	F       func(string) string
	Timeout int
}

//New 签名认证
func New(conf Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			req := ctx.Request()
			if req.Header.Get("appid") == "" {
				return ctx.JSON(http.StatusForbidden, "miss appid header")
			}
			appsecret := conf.F(req.Header.Get("appid"))
			if appsecret == "" {
				return ctx.JSON(http.StatusForbidden, "not exist this appid")
			}
			clientSignature := req.Header.Get("signature")
			if clientSignature == "" {
				return ctx.JSON(http.StatusForbidden, "miss signature header")
			}
			if req.Header.Get("timestamp") == "" {
				return ctx.JSON(http.StatusForbidden, "miss timestamp header")
			}
			u, err := time.Parse("2006-01-02 15:04:05", req.Header.Get("timestamp"))
			if err != nil {
				return ctx.JSON(http.StatusForbidden, "timestamp format is error, should 2006-01-02 15:04:05")
			}
			t := time.Now()
			if t.Sub(u).Seconds() > float64(conf.Timeout) {
				return ctx.JSON(http.StatusForbidden, "timeout! the request time is long ago, please try again")
			}
			var requestURL string
			var body []byte
			if req.Method == echo.GET {
				requestURL = req.RequestURI
			} else {
				requestURL = req.URL.Path
				body, err = getRawBody(req)
				newBody := ioutil.NopCloser(bytes.NewBuffer(body))
				req.Body = newBody
			}
			serviceSignature := Signature(appsecret, req.Method, body, requestURL, req.Header.Get("timestamp"))
			// fmt.Println("serviceSignature:", serviceSignature)
			// fmt.Println("clientSignature:", clientSignature)
			if clientSignature != serviceSignature {
				return ctx.JSON(http.StatusForbidden, "Signature Failed")
			}
			return next(ctx)
		}
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

func getRawBody(r *http.Request) ([]byte, error) {
	body := r.Body
	defer body.Close()
	rawBody, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return rawBody, nil
}
