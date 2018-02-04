package auth

import (
	"github.com/xwinie/glue/core"
	"github.com/xwinie/glue/core/middleware/casbin"
)

//OpenPermission 根据角色获取权限
func OpenPermission() ([]casbin.Permission, error) {
	return openPermission()
}

// 获取用户菜单
func menuByUserIDService(userID int64) (responseEntity core.ResponseEntity) {
	roleID, _ := findRoleIDByUserID(userID)
	menus, err := findResourceByMultiRole(roleID, 1)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(QueryError, getMsg(QueryError)))
	}
	return *responseEntity.Build(menus)
}

func findResourceCountByPageService() int64 {
	num, err := resourceCountByPage()
	if err != nil {
		return 0
	}
	return num
}

func findResourceByPageService(p *core.Paginator) (responseEntity core.ResponseEntity) {
	resources, err := resourceByPage(p.PerPageNums, p.Offset())
	type data struct {
		Resources []*SysResource
		Total     int64
	}
	d := &data{resources, p.Nums()}
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(QueryError, getMsg(QueryError)))
	}
	return *responseEntity.Build(d)
}
