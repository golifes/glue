package auth

import (
	"github.com/xwinie/glue/lib/db"
)

//SysClient 客户端管理
type SysClient struct {
	Id           int64
	ClientId     string
	Name         string
	Secret       string
	VerifySecret string
	Locked       int8
}

//getClient 获取客户端信息
func getClient(clientID string) (SysClient, error) {
	client := new(SysClient)
	o := db.New()
	_, err := o.Table(client).Where("client_id = ?", clientID).Get(client)
	return *client, err
}
