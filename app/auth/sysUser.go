package auth

import (
	"time"

	"github.com/xwinie/glue/lib/db"
	"github.com/xwinie/glue/lib/utils"
)

//SysUser 用户
type SysUser struct {
	ID           int64     `xorm:"pk bigint 'id'"`
	Account      string    `xorm:"unique notnull"`
	Name         string    `xorm:"varchar(200) notnull"`
	UserType     int8      `xorm:"tinyint default(0) notnull"` //0是第三方用户1是self
	Password     string    `xorm:"varchar(200) notnull"`
	Salt         string    `xorm:"varchar(200) notnull"`
	DeleteStatus int8      `xorm:"tinyint default(0) notnull"`
	Created      time.Time `xorm:"timestamp created notnull"`
	Updated      time.Time `xorm:"timestamp updated  notnull"`
	Locked       int8      `xorm:"tinyint default(0) notnull"`
}

//CheckEqualPassword Md5(Md5(Sha1("12345") + Sha1("passwod")) + salt)
func (u SysUser) CheckEqualPassword(password string) bool {
	return u.Password == utils.Md5(password+u.Salt)
}

//EncryptionPassword 加密密码
func (u SysUser) EncryptionPassword(password string) string {
	return utils.Md5(password + u.Salt)
}

func findUserAllColums(account string) (user SysUser, err error) {
	o := db.New()
	_, err = o.Table(&user).Where("account = ?", account).Get(&user)
	return user, err
}
