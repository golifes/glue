package auth

import (
	"time"

	"github.com/xwinie/glue/core"
)

//SysRole 角色
type SysRole struct {
	ID           int64     `xorm:"pk bigint 'id'" json:"Id"`
	Code         string    `xorm:"varchar(100) unique notnull"`
	Name         string    `xorm:"varchar(200)  notnull"`
	Description  string    `xorm:"varchar(250)"`
	DeleteStatus int8      `xorm:"tinyint default(0) notnull"`
	Created      time.Time `xorm:"timestamp created notnull"`
	Updated      time.Time `xorm:"timestamp updated  notnull"`
	Locked       int8      `xorm:"tinyint default(0) notnull"`
}

//QuerySysRole 查询返回字符串
type QuerySysRole struct {
	ID           string `json:"Id"`
	Code         string
	Name         string
	Description  string
	DeleteStatus int8
	Created      time.Time
	Updated      time.Time
	Locked       int8
}

func (u SysRole) insert() error {
	o := core.New()
	_, err := o.Insert(u)
	return err
}
func (u SysRole) codeIsExist() (entity core.Entity) {
	o := core.New()
	has, err := o.Table(&u).Where("code = ?", u.Code).Exist()
	if has || err != nil {
		return entity.New(RoleIsExist, getMsg(RoleIsExist))
	}
	return entity.New(Success, getMsg(Success))
}

func deleteRole(ID int64) error {
	o := core.New()
	_, err := o.Table("sys_role").Where("id = ?", ID).Update(map[string]interface{}{"delete_status": 1})
	return err
}
func (u *SysRole) update() error {
	o := core.New()
	_, err := o.Where("id = ?", u.ID).Update(u)
	return err
}

func roleCountByPage() (num int64, err error) {
	o := core.New()
	num, err = o.Table("sys_role").Count()
	return num, err
}

func roleByPage(pageSize int, offset int) (roles []*SysRole, err error) {
	o := core.New()
	err = o.Table("sys_role").Limit(pageSize, offset).Find(&roles)
	return roles, err
}

func findRoleByCode(code string) (role SysRole, err error) {
	o := core.New()
	_, err = o.Table("sys_role").Where("code = ?", code).Get(&role)
	return role, err
}
func findRoleByID(id int64) (u SysRole, err error) {
	o := core.New()
	_, err = o.Table(&u).Where("id = ?", id).Get(&u)
	return u, err
}
