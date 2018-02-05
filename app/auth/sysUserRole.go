package auth

import (
	"time"

	"github.com/xwinie/glue/core"
)

//SysUserRole 用户角色
type SysUserRole struct {
	ID           int64     `xorm:"pk bigint 'id'"`
	RoleID       int64     `xorm:"bigint notnull 'role_id'"`
	UserID       int64     `xorm:"bigint notnull 'user_id'"`
	DeleteStatus int8      `xorm:"tinyint default(0) notnull"`
	Created      time.Time `xorm:"timestamp created notnull"`
	Updated      time.Time `xorm:"timestamp updated  notnull"`
	Locked       int8      `xorm:"tinyint default(0) notnull"`
}

func findRoleIDByUserID(userID int64) ([]int64, error) {
	var roleIds []int64
	o := core.New()
	err := o.Table("sys_user_role").Cols("role_id").Where("user_id = ?", userID).Find(&roleIds)
	return roleIds, err
}
func findRoleByUserID(userID int64) ([]QuerySysRole, error) {
	var m []QuerySysRole
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

func deleteUserRole(userId int64) error {
	o := core.New()
	_, err := o.Where("user_id = ?", userId).Delete(new(SysUserRole))
	return err
}
