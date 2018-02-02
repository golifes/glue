package auth

import (
	"github.com/xwinie/glue/lib/middleware/casbin"
	"github.com/xwinie/glue/lib/response"
)

//OpenPermission 根据角色获取权限
func OpenPermission() ([]casbin.Permission, error) {
	return openPermission()
}

//MenuByUserIDService 获取用户菜单
func MenuByUserIDService(userID int64) (responseEntity response.ResponseEntity) {
	roleID, _ := findRoleIDByUserID(userID)
	menus, err := findResourceByMultiRole(roleID, 1)
	if err != nil {
		return *responseEntity.BuildError(response.BuildEntity(QueryError, getMsg(QueryError)))
	}
	return *responseEntity.Build(menus)
}
