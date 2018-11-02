package crm

import (
	"strconv"
	"time"

	"github.com/xwinie/glue/core"
)

//Customer 顾客
type Customer struct {
	ID           string    `xorm:"pk bigint 'id'" json:"Id"`
	Phone        string    `xorm:"varchar(20) notnull"`
	Gender       string    `xorm:"bigint "`
	Age          string    `xorm:"tinyint"`
	City         string    `xorm:"bigint"`
	SourceType   string    `xorm:"bigint"`
	Name         string    `xorm:"varchar(200) notnull"`
	CompanyName  string    `xorm:"varchar(200) notnull"`
	CompanyId    string    `xorm:"bigint"`
	DeleteStatus int8      `xorm:"tinyint default(0) notnull"`
	Created      time.Time `xorm:"datetime created notnull"`
	Updated      time.Time `xorm:"timestamp updated  notnull"`
}

//QueryCustomer 结构
type QueryCustomer struct {
	ID          string `xorm:"pk bigint 'id'" json:"Id"`
	Phone       string `xorm:"varchar(20) notnull"`
	Gender      string `xorm:"bigint "`
	Age         string `xorm:"tinyint"`
	City        string `xorm:"bigint"`
	SourceType  string `xorm:"bigint"`
	Name        string `xorm:"varchar(200) notnull"`
	CompanyName string `xorm:"varchar(200) notnull"`
	CompanyId   string `xorm:"bigint"`
}

func (u Customer) insert() error {
	o := core.New()
	_, err := o.Insert(u)
	return err
}

func (u Customer) update() error {
	i64, err := strconv.ParseInt(u.ID, 10, 64)
	if err != nil {
		return err
	}
	o := core.New()
	_, err = o.Where("id = ?", i64).Update(u)
	return err
}

func deleteCustomer(ID int64) error {
	o := core.New()
	_, err := o.Table("customer").Where("id = ?", ID).Update(map[string]interface{}{"delete_status": 1})
	return err
}
func getCustomer(id string) (Customer, error) {
	c := new(Customer)
	o := core.New()
	_, err := o.Table(c).Where("id = ?", id).Get(c)
	return *c, err
}
func customerCountByPage() (num int64, err error) {
	o := core.New()
	num, err = o.Table("customer").Count()
	return num, err
}
func customerByPage(pageSize int, offset int) (m []*QueryCustomer, err error) {
	o := core.New()
	err = o.Table("customer").Limit(pageSize, offset).Find(&m)
	return m, err
}

//CustomerTag 顾客标签
type CustomerTag struct {
	ID           int64     `xorm:"pk bigint 'id'"`
	CustomerID   string    `xorm:"bigint notnull 'customer_id'"`
	TagType      int64     `xorm:"bigint "`
	TagName      string    `xorm:"varchar(200) notnull"`
	DeleteStatus int8      `xorm:"tinyint default(0) notnull"`
	Created      time.Time `xorm:"datetime created notnull"`
	Updated      time.Time `xorm:"timestamp updated  notnull"`
}

func findCustomerTag(customerId string) (u []string, err error) {
	o := core.New()
	err = o.Cols("tag_name").Table("customer_tag").Where("customer_id = ?", customerId).Find(&u)
	return u, err
}
func insertCustomerTag(m []CustomerTag) error {
	o := core.New()
	_, err := o.Insert(&m)
	return err
}

func deleteCustomerTagByCustomerId(customerId string) error {
	o := core.New()
	_, err := o.Where("customer_id = ?", customerId).Delete(new(CustomerTag))
	return err
}
