package auth

import (
	"strconv"

	"github.com/xwinie/glue/core"
)

func findRoleIDByUser(userID int64) ([]int64, error) {
	return findRoleIDByUserID(userID)
}

func userAllotRole(userId int64, roleIds []int64) (responseEntity core.ResponseEntity) {
	sysUserRole := new([]SysUserRole)
	G, _ := core.NewGUID(2)
	for _, value := range roleIds {
		m := new(SysUserRole)
		id, _ := G.NextID()
		m.ID = id
		m.UserID = userId
		m.RoleID = value
		*sysUserRole = append(*sysUserRole, *m)
	}
	err := deleteUserRole(userId)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(ParameterError, getMsg(ParameterError)))
	}
	err = insertUserRole(*sysUserRole)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(ParameterError, getMsg(ParameterError)))
	}
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/user/"+strconv.FormatInt(userId, 10)+"/role", "self", "GET", "根据用户id获取角色"))
	hateoas.AddLinks(links)
	type data struct {
		*core.Hateoas
	}
	d := &data{&hateoas}
	return *responseEntity.Build(d)
}

func findRoleByUserIDService(userID int64) (responseEntity core.ResponseEntity) {
	u, err := findRoleByUserID(userID)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(QueryError, getMsg(QueryError)))
	}
	return *responseEntity.Build(u)
}
