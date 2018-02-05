package auth

import (
	"strconv"
	"time"

	"github.com/xwinie/glue/core"
)

//SysUser 用户
type SysUser struct {
	ID           string    `xorm:"pk bigint 'id'" json:"Id"`
	Account      string    `xorm:"unique varchar(100) notnull"`
	Name         string    `xorm:"varchar(200) notnull"`
	UserType     int8      `xorm:"tinyint default(0) notnull"` //0是第三方用户1是self
	Password     string    `xorm:"varchar(200) notnull"`
	Salt         string    `xorm:"varchar(200) notnull"`
	DeleteStatus int8      `xorm:"tinyint default(0) notnull"`
	Created      time.Time `xorm:"datetime created notnull"`
	Updated      time.Time `xorm:"timestamp updated  notnull"`
	Locked       int8      `xorm:"tinyint default(0) notnull"`
}

//QuerySysUser 列表查询结果结构
type QuerySysUser struct {
	ID           string `xorm:"'id'" json:"Id"`
	Account      string
	Name         string
	UserType     int8
	DeleteStatus int8
	Created      time.Time
	Updated      time.Time
	Locked       int8
}

//CheckEqualPassword Md5(Md5(Sha1("12345") + Sha1("passwod")) + salt)
func (u SysUser) CheckEqualPassword(password string) bool {
	return u.Password == core.Md5(password+u.Salt)
}

//EncryptionPassword 加密密码
func (u SysUser) EncryptionPassword(password string) string {
	return core.Md5(password + u.Salt)
}

func (u SysUser) insert() error {
	o := core.New()
	_, err := o.Insert(u)
	return err
}

func (u SysUser) update() error {
	i64, err := strconv.ParseInt(u.ID, 10, 64)
	if err != nil {
		return err
	}
	o := core.New()
	_, err = o.Where("id = ?", i64).Update(u)
	return err
}
func (u SysUser) accountIsExist() (entity core.Entity) {
	o := core.New()
	has, err := o.Table(&u).Where("account = ?", u.Account).Exist()
	if has || err != nil {
		return entity.New(UserIsExist, getMsg(UserIsExist))
	}
	return entity.New(Success, getMsg(Success))
}

func deleteUser(ID int64) error {
	o := core.New()
	_, err := o.Table("sys_user").Where("id = ?", ID).Update(map[string]interface{}{"delete_status": 1})
	return err
}
func findUserAllColums(account string) (user SysUser, err error) {
	o := core.New()
	_, err = o.Table(&user).Where("account = ?", account).Get(&user)
	return user, err
}
func findUserByAccount(account string) (user QuerySysUser, err error) {
	o := core.New()
	_, err = o.Table("sys_user").Where("account = ?", account).Get(&user)
	return user, err
}

func findUserByID(id string) (user SysUser, err error) {
	i64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return user, err
	}
	o := core.New()
	_, err = o.Table(&user).Where("id = ?", i64).Get(&user)
	return user, err
}
func userCountByPage() (num int64, err error) {
	o := core.New()
	num, err = o.Table("sys_user").Count()
	return num, err
}

func userByPage(pageSize int, offset int) (users []*QuerySysUser, err error) {
	o := core.New()
	err = o.Table("sys_user").Limit(pageSize, offset).Find(&users)
	return users, err
}
