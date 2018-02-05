package auth

import (
	"strconv"

	"github.com/xwinie/glue/core"
)

//SysClient 客户端管理
type SysClient struct {
	ID           string `xorm:"pk bigint 'id'" json:"Id"`
	ClientID     string `xorm:"varchar(100) notnull unique 'client_id'"`
	Name         string `xorm:"varchar(200) notnull"`
	Secret       string `xorm:"varchar(200) notnull"`
	VerifySecret string `xorm:"varchar(200) notnull"`
	Locked       int8   `xorm:"tinyint default(0) notnull"`
}

//QuerySysClient 结构
type QuerySysClient struct {
	ID       string `xorm:"'id'" json:"Id"`
	ClientID string `xorm:"'client_id'" json:"ClientId"`
	Name     string
	Locked   int8
}

//getClient 获取客户端信息
func getClient(clientID string) (SysClient, error) {
	client := new(SysClient)
	o := core.New()
	_, err := o.Table(client).Where("client_id = ?", clientID).Get(client)
	return *client, err
}

func (u SysClient) insert() error {
	o := core.New()
	_, err := o.Insert(u)
	return err
}

func (u SysClient) update() error {
	o := core.New()
	i64, err := strconv.ParseInt(u.ID, 10, 64)
	if err != nil {
		return err
	}
	_, err = o.Where("id = ?", i64).Update(u)
	return err
}

func (u SysClient) codeIsExist() (entity core.Entity) {
	o := core.New()
	has, err := o.Table(&u).Where("client_id = ?", u.ClientID).Exist()
	if has || err != nil {
		return entity.New(ClientIsExist, getMsg(ClientIsExist))
	}
	return entity.New(Success, getMsg(Success))
}

func updateClientLock(ID string, locked int8) error {
	i64, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return err
	}
	o := core.New()
	_, err = o.Table("sys_client").Where("id = ?", i64).Update(map[string]interface{}{"locked": locked})
	return err
}
func clientCountByPage() (num int64, err error) {
	o := core.New()
	num, err = o.Table("sys_client").Count()
	return num, err
}

func clientByPage(pageSize int, offset int) (m []*QuerySysClient, err error) {
	o := core.New()
	err = o.Table("sys_client").Limit(pageSize, offset).Find(&m)
	return m, err
}

func finClientByID(id string) (u SysClient, err error) {
	o := core.New()
	i64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return u, err
	}
	_, err = o.Table(&u).Where("id = ?", i64).Get(&u)
	return u, err
}
