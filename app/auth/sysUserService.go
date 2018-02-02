package auth

import (
	. "github.com/xwinie/glue/lib/hateoas"
	. "github.com/xwinie/glue/lib/response"
	"github.com/xwinie/glue/lib/utils"
)

//FindUserCountByPageService 分页获取总数
func FindUserCountByPageService() int64 {
	num, err := userCountByPage()
	if err != nil {
		return 0
	}
	return num
}

//FindUserByPageService 分页获取用户信息
func FindUserByPageService(p *utils.Paginator) (responseEntity ResponseEntity) {
	users, err := userByPage(p.PerPageNums, p.Offset())
	var hateoas HateoasTemplate
	var links Links
	links.Add(LinkTo("/v1/user/{account}", "self", "GET", "根据用户账号获取用户信息"))
	links.Add(LinkTo("/v1/user/{id}", "self", "DELETE", "根据id删除用户信息"))
	links.Add(LinkTo("/v1/user/{id}", "self", "PUT", "根据id修改用户信息"))
	links.Add(LinkTo(p.PageLinkFirst(), "first", "GET", ""))
	links.Add(LinkTo(p.PageLinkLast(), "last", "GET", ""))
	if p.HasNext() {
		links.Add(LinkTo(p.PageLinkNext(), "next", "GET", ""))
	}
	if p.HasPrev() {
		links.Add(LinkTo(p.PageLinkPrev(), "prev", "GET", ""))
	}
	hateoas.AddLinks(links)
	type data struct {
		Users []*QuerySysUser
		Total int64
		*HateoasTemplate
	}
	d := &data{users, p.Nums(), &hateoas}
	if err != nil {
		return *responseEntity.BuildError(BuildEntity(QueryError, getMsg(QueryError)))
	}
	return *responseEntity.Build(d)
}
