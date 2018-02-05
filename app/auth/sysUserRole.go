package auth

import (
	"strconv"
	"time"

	"github.com/xwinie/glue/core"
)

//SysUserRole 用户角色
type SysUserRole struct {
	ID           int64     `xorm:"pk bigint 'id'"`
	RoleID       int64     `xorm:"bigint notnull 'role_id'"`
	UserID       int64     `xorm:"bigint notnull 'user_id'"`
	DeleteStatus int8      `xorm:"tinyint default(0) notnull"`
	Created      time.Time `xorm:"datetime created notnull"`
	Updated      time.Time `xorm:"timestamp updated  notnull"`
	Locked       int8      `xorm:"tinyint default(0) notnull"`
}

func findRoleIDByUserID(userID string) ([]int64, error) {
	var roleIds []int64
	int64, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, err
	}
	o := core.New()
	err = o.Table("sys_user_role").Cols("role_id").Where("user_id = ?", int64).Find(&roleIds)
	return roleIds, err
}
func findRoleByUserID(userID string) ([]SysRole, error) {
	var m []SysRole
	o := core.New()
	int64, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, err
	}
	err = o.Table("sys_user_role").Alias("ur").Join("INNER", []string{"sys_role", "r"}, "r.id=ur.role_id").
		Where("ur.user_id=?", int64).
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
	int64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	_, err = o.Where("user_id = ?", int64).Delete(new(SysUserRole))
	return err
}
