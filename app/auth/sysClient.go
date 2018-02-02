package auth

import (
	"github.com/xwinie/glue/lib/db"
)

//SysClient 客户端管理
type SysClient struct {
	ID           int64  `xorm:"pk bigint 'id'"`
	ClientID     string `xorm:"varchar(100) notnull unique 'client_id'"`
	Name         string `xorm:"varchar(200) notnull"`
	Secret       string `xorm:"varchar(200) notnull"`
	VerifySecret string `xorm:"varchar(200) notnull"`
	Locked       int8   `xorm:"tinyint default(0) notnull"`
}

//getClient 获取客户端信息
func getClient(clientID string) (SysClient, error) {
	client := new(SysClient)
	o := db.New()
	_, err := o.Table(client).Where("client_id = ?", clientID).Get(client)
	return *client, err
}
