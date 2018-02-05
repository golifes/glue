package auth

import (
	"strconv"

	"github.com/xwinie/glue/core"
)

//findUserCountByPageService 分页获取总数
func findUserCountByPageService() int64 {
	num, err := userCountByPage()
	if err != nil {
		return 0
	}
	return num
}

//createUser 创建用户
func createUser(u SysUser) (responseEntity core.ResponseEntity) {
	isExist := u.accountIsExist()
	if isExist.Code != 100000 {
		return *responseEntity.BuildError(core.BuildEntity(UserIsExist, getMsg(UserIsExist)))
	}
	G, _ := core.NewGUID(1)
	id, _ := G.NextID()
	u.ID = strconv.FormatInt(id, 10)
	u.Salt = core.RandStringByLen(10)

	err := u.insert()
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(CreateUserError, getMsg(CreateUserError)))
	}
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/user/"+u.Account, "self", "GET", "根据用户账号获取用户信息"))
	links.Add(core.LinkTo("/v1/user/"+strconv.FormatInt(id, 10), "self", "DELETE", "根据id删除用户信息"))
	links.Add(core.LinkTo("/v1/user/"+strconv.FormatInt(id, 10), "self", "PUT", "根据id修改用户信息"))
	hateoas.AddLinks(links)
	type data struct {
		*core.Hateoas
	}
	d := &data{&hateoas}
	return *responseEntity.BuildPostAndPut(d)
}

func deleteUserService(id int64) (responseEntity core.ResponseEntity) {
	err := deleteUser(id)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(DeleteUserError, getMsg(DeleteUserError)))
	}
	return *responseEntity.BuildDelete(core.BuildEntity(Success, getMsg(Success)))
}

func updateUserService(id string, m *SysUser) (responseEntity core.ResponseEntity) {

	m.ID = id
	m.Account = ""
	err := m.update()
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(UpdateUserError, getMsg(UpdateUserError)+err.Error()))
	}
	u, _ := findUserByID(id)
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/user/"+u.Account, "self", "GET", "根据用户账号获取用户信息"))
	links.Add(core.LinkTo("/v1/user/"+id, "self", "DELETE", "根据id删除用户信息"))
	links.Add(core.LinkTo("/v1/user/"+id, "self", "PUT", "根据id修改用户信息"))
	hateoas.AddLinks(links)
	type data struct {
		*core.Hateoas
	}
	d := &data{&hateoas}
	return *responseEntity.BuildPostAndPut(d)
}
func findUserByAccountService(account string) (responseEntity core.ResponseEntity) {
	u, err := findUserByAccount(account)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(QueryError, getMsg(QueryError)))
	}
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/user/"+u.ID, "self", "DELETE", "根据id删除用户信息"))
	links.Add(core.LinkTo("/v1/user/"+u.ID, "self", "PUT", "根据id修改用户信息"))
	hateoas.AddLinks(links)
	type data struct {
		QuerySysUser
		*core.Hateoas
	}
	d := &data{u, &hateoas}
	return *responseEntity.Build(d)
}

//FindUserByPageService 分页获取用户信息
func findUserByPageService(p *core.Paginator) (responseEntity core.ResponseEntity) {
	users, err := userByPage(p.PerPageNums, p.Offset())
	var hateoas core.HateoasTemplate
	var links core.Links
	links.Add(core.LinkTo("/v1/user/{account}", "self", "GET", "根据用户账号获取用户信息"))
	links.Add(core.LinkTo("/v1/user/{id}", "self", "DELETE", "根据id删除用户信息"))
	links.Add(core.LinkTo("/v1/user/{id}", "self", "PUT", "根据id修改用户信息"))
	links.Add(core.LinkTo(p.PageLinkFirst(), "first", "GET", ""))
	links.Add(core.LinkTo(p.PageLinkLast(), "last", "GET", ""))
	if p.HasNext() {
		links.Add(core.LinkTo(p.PageLinkNext(), "next", "GET", ""))
	}
	if p.HasPrev() {
		links.Add(core.LinkTo(p.PageLinkPrev(), "prev", "GET", ""))
	}
	hateoas.AddLinks(links)
	type data struct {
		Users []*QuerySysUser
		Total int64
		*core.HateoasTemplate
	}
	d := &data{users, p.Nums(), &hateoas}
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(QueryError, getMsg(QueryError)))
	}
	return *responseEntity.Build(d)
}
