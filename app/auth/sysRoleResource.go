package auth

import (
	"time"

	"github.com/xwinie/glue/lib/middleware/casbin"
)

//SysRoleResource 角色资源
type SysRoleResource struct {
	Id           int64
	RoleId       int64
	ResourceId   int64
	DeleteStatus int8
	Created      time.Time
	Updated      time.Time
	Locked       int8
}

func permissionByMultiRole(roleIds []int64, resType int8) (resource []casbin.Permission, num int64, err error) {
	return resource, num, err
}
