package auth

import "github.com/xwinie/glue/lib/middleware/casbin"

func PermissionByMultiRole(roleIds []int64, resType int8) ([]casbin.Permission, int64, error) {
	return permissionByMultiRole(roleIds, resType)
}
