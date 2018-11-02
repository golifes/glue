package crm

import (
	"fmt"
	"strconv"

	"github.com/xwinie/glue/core"
)

func findCustomerByCustomerIdService(id string) (responseEntity core.ResponseEntity) {
	u, err := getCustomer(id)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(QueryError, getMsg(QueryError)))
	}
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/customer/"+u.ID, "self", "DELETE", "根据id删除顾客信息"))
	links.Add(core.LinkTo("/v1/customer/"+u.ID, "self", "PUT", "根据id修改顾客信息"))
	hateoas.AddLinks(links)
	type data struct {
		Customer Customer
		*core.Hateoas
	}
	d := &data{u, &hateoas}
	return *responseEntity.Build(d)
}

func findCustomerCountByPageService() int64 {
	num, err := customerCountByPage()
	if err != nil {
		return 0
	}
	return num
}

func findCustomerByPageService(p *core.Paginator) (responseEntity core.ResponseEntity) {
	m, err := customerByPage(p.PerPageNums, p.Offset())
	var hateoas core.HateoasTemplate
	var links core.Links
	links.Add(core.LinkTo("/v1/customer/{id}", "self", "GET", "根据id获取顾客信息"))
	links.Add(core.LinkTo("/v1/customer/{id}", "self", "DELETE", "根据id删除顾客信息"))
	links.Add(core.LinkTo("/v1/customer/{id}", "self", "PUT", "根据id修改顾客信息"))
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
		Customers []*QueryCustomer
		Total     int64
		*core.HateoasTemplate
	}
	d := &data{m, p.Nums(), &hateoas}
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(QueryError, getMsg(QueryError)))
	}
	return *responseEntity.Build(d)
}

func updateCustomerService(id string, m *Customer) (responseEntity core.ResponseEntity) {
	m.ID = id
	err := m.update()
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(UpdateCustomerError, getMsg(UpdateCustomerError)))
	}
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/customer/"+id, "self", "GET", "根据id获取顾客信息"))
	links.Add(core.LinkTo("/v1/customer/"+id, "self", "DELETE", "根据id锁顾客信息"))
	links.Add(core.LinkTo("/v1/customer/"+id, "self", "PUT", "根据id修改顾客信息"))
	hateoas.AddLinks(links)
	type data struct {
		*core.Hateoas
	}
	d := &data{&hateoas}
	return *responseEntity.BuildPostAndPut(d)
}

func deleteCustomerService(id int64) (responseEntity core.ResponseEntity) {
	err := deleteCustomer(id)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(DeleteCustomerError, getMsg(DeleteCustomerError)))
	}
	return *responseEntity.BuildDelete(core.BuildEntity(Success, getMsg(Success)))
}
func createCustomer(u Customer) (responseEntity core.ResponseEntity) {
	G, _ := core.NewGUID(1)
	id, _ := G.NextID()
	u.ID = strconv.FormatInt(id, 10)
	u.CompanyId = "0"
	err := u.insert()
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(CreateCustomerError, getMsg(CreateCustomerError)+err.Error()))
	}
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/customer", "self", "GET", "根据用户账号分页获取顾客信息"))
	links.Add(core.LinkTo("/v1/customer/"+strconv.FormatInt(id, 10), "self", "PUT", "根据id修改顾客信息"))
	hateoas.AddLinks(links)
	type data struct {
		*core.Hateoas
	}
	d := &data{&hateoas}
	return *responseEntity.BuildPostAndPut(d)
}

func createCustomerTagBycustomerId(customerId string, typeName []string) (responseEntity core.ResponseEntity) {
	customerTag := new([]CustomerTag)
	G, _ := core.NewGUID(3)
	for _, value := range typeName {
		m := new(CustomerTag)
		id, _ := G.NextID()
		m.ID = id
		m.CustomerID = customerId
		m.TagName = value
		*customerTag = append(*customerTag, *m)
	}
	err := deleteCustomerTagByCustomerId(customerId)
	if err != nil {
		fmt.Println(err)
		return *responseEntity.BuildError(core.BuildEntity(ParameterError, getMsg(ParameterError)))
	}
	err = insertCustomerTag(*customerTag)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(CreateCustomerTagError, getMsg(CreateCustomerTagError)))
	}
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/customer/"+customerId+"/tag", "self", "GET", "获取用户标签"))
	hateoas.AddLinks(links)
	type data struct {
		*core.Hateoas
	}
	d := &data{&hateoas}
	return *responseEntity.BuildPostAndPut(d)
}

func findCustomerTagByCustomerIdService(id string) (responseEntity core.ResponseEntity) {
	u, err := findCustomerTag(id)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(QueryError, getMsg(QueryError)))
	}
	var hateoas core.Hateoas
	var links core.Links
	links.Add(core.LinkTo("/v1/customer/"+id+"/tag", "self", "POST", "添加tag"))
	hateoas.AddLinks(links)
	type data struct {
		CustomerTags []string
		*core.Hateoas
	}
	d := &data{u, &hateoas}
	return *responseEntity.Build(d)
}
