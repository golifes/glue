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
