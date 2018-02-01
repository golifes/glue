package auth

import (
	"time"

	"github.com/xwinie/glue/lib/db"
)

//SysUserRole 用户角色
type SysUserRole struct {
	Id           int64
	RoleId       int64
	UserId       int64
	DeleteStatus int8
	Created      time.Time
	Updated      time.Time
	Locked       int8
}

func findRoleIDByUserID(userID int64) ([]int64, error) {
	var roleIds []int64
	o := db.New()
	err := o.Table("sys_user_role").Cols("role_id").Where("user_id = ?", userID).Find(&roleIds)
	return roleIds, err
}
