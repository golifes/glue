package auth

import (
	"strconv"

	"github.com/xwinie/glue/core"
)

func findRoleIDByUser(userID string) ([]string, error) {
	return findRoleIDByUserID(userID)
}

func userAllotRole(userID string, roleIds []string) (responseEntity core.ResponseEntity) {
	sysUserRole := new([]SysUserRole)
	G, _ := core.NewGUID(2)
	for _, value := range roleIds {
		m := new(SysUserRole)
		id, _ := G.NextID()
		m.ID = strconv.FormatInt(id, 10)
		m.UserID = userID
		m.RoleID = value
		*sysUserRole = append(*sysUserRole, *m)
	}
	err := deleteUserRole(userID)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(ParameterError, getMsg(ParameterError)))
	}
	err = insertUserRole(*sysUserRole)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(ParameterError, getMsg(ParameterError)))
	}
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/user/"+userID+"/role", "self", "GET", "根据用户id获取角色"))
	hateoas.AddLinks(links)
	type data struct {
		*core.Hateoas
	}
	d := &data{&hateoas}
	return *responseEntity.Build(d)
}

func findRoleByUserIDService(userID string) (responseEntity core.ResponseEntity) {
	u, err := findRoleByUserID(userID)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(QueryError, getMsg(QueryError)))
	}
	type data struct {
		Roles interface{}
	}
	d := &data{u}
	return *responseEntity.Build(d)
}
