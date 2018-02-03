package auth

import (
	"fmt"
	"strconv"

	"github.com/xwinie/glue/core"
	"github.com/xwinie/glue/core/middleware/casbin"
)

//PermissionByMultiRole 根据角色获取权限
func PermissionByMultiRole(roleIds interface{}, resType int8) ([]casbin.Permission, error) {
	return permissionByMultiRole(roleIds, resType)
}

func roleAllotResource(roleId int64, resourceIds []int64) (responseEntity core.ResponseEntity) {
	roleResources := make([]SysRoleResource, len(resourceIds))
	G, _ := core.NewGUID(2)
	for index, value := range resourceIds {
		m := new(SysRoleResource)
		id, _ := G.NextID()
		m.ID = id
		m.ResourceID = value
		m.RoleID = roleId
		fmt.Printf("arr[%d]=%d \n", index, value)
		roleResources = append(roleResources, *m)
	}
	err := deleteRoleResource(roleId)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(ParameterError, getMsg(ParameterError)))
	}
	err = insertRoleResource(roleResources)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(ParameterError, getMsg(ParameterError)))
	}
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/role/"+strconv.FormatInt(roleId, 10)+"/resource", "self", "GET", "根据用户id获取角色"))
	hateoas.AddLinks(links)
	type data struct {
		*core.Hateoas
	}
	d := &data{&hateoas}
	return *responseEntity.Build(d)
}

func findResourceByRoleIdService(roleId int64) (responseEntity core.ResponseEntity) {
	u, err := findResourceByRoleId(roleId)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(QueryError, getMsg(QueryError)))
	}
	return *responseEntity.Build(u)
}
