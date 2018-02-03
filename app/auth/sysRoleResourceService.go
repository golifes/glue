package auth

import "github.com/xwinie/glue/core/middleware/casbin"

//PermissionByMultiRole 根据角色获取权限
func PermissionByMultiRole(roleIds interface{}, resType int8) ([]casbin.Permission, error) {
	return permissionByMultiRole(roleIds, resType)
}
