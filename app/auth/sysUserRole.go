package auth

import (
	"time"

	"github.com/xwinie/glue/core"
)

//SysUserRole 用户角色
type SysUserRole struct {
	ID           string    `xorm:"pk varchar(20) 'id'"`
	RoleID       string    `xorm:"varchar(20) notnull 'role_id'"`
	UserID       string    `xorm:"varchar(20) notnull 'user_id'"`
	DeleteStatus int8      `xorm:"tinyint default(0) notnull"`
	Created      time.Time `xorm:"timestamp created notnull"`
	Updated      time.Time `xorm:"timestamp updated  notnull"`
	Locked       int8      `xorm:"tinyint default(0) notnull"`
}

func findRoleIDByUserID(userID string) ([]string, error) {
	var roleIds []string
	o := core.New()
	err := o.Table("sys_user_role").Cols("role_id").Where("user_id = ?", userID).Find(&roleIds)
	return roleIds, err
}
func findRoleByUserID(userID string) ([]SysRole, error) {
	var m []SysRole
	o := core.New()
	err := o.Table("sys_user_role").Alias("ur").Join("INNER", []string{"sys_role", "r"}, "r.id=ur.role_id").
		Where("ur.user_id=?", userID).
		Cols("r.*").
		Find(&m)
	return m, err
}

func insertUserRole(m []SysUserRole) error {
	o := core.New()
	_, err := o.Insert(&m)
	return err
}

func deleteUserRole(id string) error {
	o := core.New()
	_, err := o.Where("user_id = ?", id).Delete(new(SysUserRole))
	return err
}
