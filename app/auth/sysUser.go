package auth

import (
	"time"

	"github.com/xwinie/glue/lib/db"
	. "github.com/xwinie/glue/lib/utils"
)

type SysUser struct {
	Id           int64
	Account      string
	Name         string
	UserType     int8 //0是第三方用户1是self
	Password     string
	Salt         string
	DeleteStatus int8
	Created      time.Time
	Updated      time.Time
	Locked       int8
}

//CheckEqualPassword Md5(Md5(Sha1("12345") + Sha1("passwod")) + salt)
func (u SysUser) CheckEqualPassword(password string) bool {
	return u.Password == Md5(password+u.Salt)
}

//EncryptionPassword 加密密码
func (u SysUser) EncryptionPassword(password string) string {
	return Md5(password + u.Salt)
}

func findUserAllColums(account string) (user SysUser, err error) {
	o := db.New()
	_, err = o.Table(&user).Where("account = ?", account).Get(&user)
	return user, err
}
