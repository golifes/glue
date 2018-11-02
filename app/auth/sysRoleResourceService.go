package auth

import (
	"strconv"

	"github.com/xwinie/glue/core"
	"github.com/xwinie/glue/core/middleware/casbin"
)

//PermissionByMultiRole 根据角色获取权限
func PermissionByMultiRole(roleIds interface{}, resType int8) ([]casbin.Permission, error) {
	return permissionByMultiRole(roleIds, resType)
}

func roleAllotResource(roleID string, resourceIds []string) (responseEntity core.ResponseEntity) {
	roleResources := new([]SysRoleResource)
	G, _ := core.NewGUID(2)
	for _, value := range resourceIds {
		m := new(SysRoleResource)
		id, _ := G.NextID()
		m.ID = id
		resourceInt64, _ := strconv.ParseInt(value, 10, 64)
		m.ResourceID = resourceInt64
		roleInt64, _ := strconv.ParseInt(roleID, 10, 64)
		m.RoleID = roleInt64
		*roleResources = append(*roleResources, *m)
	}
	err := deleteRoleResource(roleID)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(ParameterError, getMsg(ParameterError)))
	}
	err = insertRoleResource(*roleResources)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(ParameterError, getMsg(ParameterError)))
	}
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/role/"+roleID+"/resource", "self", "GET", "根据用户id获取角色"))
	hateoas.AddLinks(links)
	type data struct {
		*core.Hateoas
	}
	d := &data{&hateoas}
	return *responseEntity.Build(d)
}

func findResourceByRoleIDService(roleID string) (responseEntity core.ResponseEntity) {
	u, err := findResourceByRoleID(roleID)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(QueryError, getMsg(QueryError)))
	}
	type data struct {
		Resources interface{}
	}
	d := &data{u}
	return *responseEntity.Build(d)
}
