package dict

import (
	"time"

	"github.com/xwinie/glue/core"
)

//SysDict 数据字典
type SysDict struct {
	ID         int64     `xorm:"pk bigint 'id'"`
	Code       string    `xorm:"varchar(50) unique notnull"`
	City       string    `xorm:"varchar(50) notnull"`
	Name       string    `xorm:"varchar(200) notnull"`
	Level      int       `xorm:"tinyint default(0) notnull"`
	ParentID   int64     `xorm:"bigint default(0) 'parent_id' notnull"`
	ParentCode string    `xorm:"varchar(50) notnull"`
	Type       string    `xorm:"varchar(50) notnull"`
	Status     int8      `xorm:"tinyint default(0) notnull"`
	Created    time.Time `xorm:"timestamp created notnull"`
	Updated    time.Time `xorm:"timestamp updated  notnull"`
}

func (m SysDict) insert() error {
	o := core.New()
	_, err := o.Insert(m)
	return err
}
func updateDict(ID int64, m map[string]interface{}) error {
	o := core.New()
	_, err := o.Table("sys_dict").Where("id = ?", ID).Update(m)
	return err
}
