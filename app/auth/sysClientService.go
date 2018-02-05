package auth

import (
	"strconv"

	"github.com/xwinie/glue/core"
)

//GetClientService 获取客户端信息
func GetClientService(clientID string) (SysClient, error) {
	return getClient(clientID)
}

func findClientByClientIDService(clientID string) (responseEntity core.ResponseEntity) {
	u, err := getClient(clientID)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(QueryError, getMsg(QueryError)))
	}
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/client/"+u.ID, "self", "DELETE", "根据id锁客户端信息"))
	links.Add(core.LinkTo("/v1/client/"+u.ID, "self", "PUT", "根据id修改客户端信息"))
	hateoas.AddLinks(links)
	type data struct {
		Client SysClient
		*core.Hateoas
	}
	d := &data{u, &hateoas}
	return *responseEntity.Build(d)
}

func findClientCountByPageService() int64 {
	num, err := clientCountByPage()
	if err != nil {
		return 0
	}
	return num
}

func findClientByPageService(p *core.Paginator) (responseEntity core.ResponseEntity) {
	m, err := clientByPage(p.PerPageNums, p.Offset())
	var hateoas core.HateoasTemplate
	var links core.Links
	links.Add(core.LinkTo("/v1/client/{code}", "self", "GET", "根据编码获取客户端信息"))
	links.Add(core.LinkTo("/v1/client/{id}", "self", "DELETE", "根据id锁客户端信息"))
	links.Add(core.LinkTo("/v1/client/{id}", "self", "PUT", "根据id修改客户端信息"))
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
		Clients []*QuerySysClient
		Total   int64
		*core.HateoasTemplate
	}
	d := &data{m, p.Nums(), &hateoas}
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(QueryError, getMsg(QueryError)))
	}
	return *responseEntity.Build(d)
}

func updateClientService(id string, m *SysClient) (responseEntity core.ResponseEntity) {
	m.ID = id
	m.ClientID = ""
	m.Secret = ""
	m.VerifySecret = ""
	err := m.update()
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(UpdateClientError, getMsg(UpdateClientError)))
	}
	u, _ := finClientByID(id)
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/client/"+u.ClientID, "self", "GET", "根据编码获取客户端信息"))
	links.Add(core.LinkTo("/v1/client/"+id, "self", "DELETE", "根据id锁客户端信息"))
	links.Add(core.LinkTo("/v1/client/"+id, "self", "PUT", "根据id修改客户端信息"))
	hateoas.AddLinks(links)
	type data struct {
		*core.Hateoas
	}
	d := &data{&hateoas}
	return *responseEntity.BuildPostAndPut(d)
}

func deleteClientService(id string, lock int8) (responseEntity core.ResponseEntity) {
	err := updateClientLock(id, lock)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(DeleteClientError, getMsg(DeleteClientError)))
	}
	return *responseEntity.BuildDelete(core.BuildEntity(Success, getMsg(Success)))
}
func createClient(u SysClient) (responseEntity core.ResponseEntity) {
	G, _ := core.NewGUID(1)
	id, _ := G.NextID()
	u.ID = strconv.FormatInt(id, 10)
	u.ClientID = strconv.FormatInt(id, 10)
	u.Secret = core.RandStringByLen(10)
	u.VerifySecret = core.RandStringByLen(10)

	err := u.insert()
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(CreateUserError, getMsg(CreateUserError)))
	}
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/client", "self", "GET", "根据用户账号分页获取客户端信息"))
	links.Add(core.LinkTo("/v1/client/"+strconv.FormatInt(id, 10), "self", "PUT", "根据id修改客户端信息"))
	hateoas.AddLinks(links)
	type data struct {
		*core.Hateoas
	}
	d := &data{&hateoas}
	return *responseEntity.BuildPostAndPut(d)
}
