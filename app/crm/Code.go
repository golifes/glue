package crm

import (
	"net/http"

	conf "github.com/xwinie/glue/app"
)

//常量
const (
	Unauthorized           = http.StatusUnauthorized
	Success                = conf.Crm
	ParameterError         = conf.Crm + 1
	QueryError             = conf.Crm + 2
	UpdateCustomerError    = conf.Crm + 3
	DeleteCustomerError    = conf.Crm + 4
	CreateCustomerError    = conf.Crm + 5
	UpdateCustomerTagError = conf.Crm + 6
	DeleteCustomerTagError = conf.Crm + 7
	CreateCustomerTagError = conf.Crm + 8
)

//Msg 错误信息
var Msg = map[int]string{
	Unauthorized:           "没有权限",
	Success:                "成功",
	ParameterError:         "参数异常",
	QueryError:             "查询异常",
	UpdateCustomerError:    "修改顾客信息异常",
	DeleteCustomerError:    "删除顾客信息异常",
	CreateCustomerError:    "创建顾客信息异常",
	UpdateCustomerTagError: "修改顾客标签异常",
	DeleteCustomerTagError: "删除顾客标签异常",
	CreateCustomerTagError: "创建顾客标签异常",
}

func getMsg(code int) string {
	return Msg[code]
}
