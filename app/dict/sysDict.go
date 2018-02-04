package dict

import "time"

type SysDict struct {
	Id         int64     `xorm:"pk bigint 'id'"`
	Code       string    `xorm:"varchar(50) unique notnull"`
	City       string    `xorm:"varchar(50) notnull"`
	Name       string    `xorm:"varchar(200) notnull"`
	Level      int       `xorm:"tinyint default(0) notnull"`
	ParentId   int64     `xorm:"bigint default(0) notnull"`
	ParentCode string    `xorm:"varchar(50) notnull"`
	Type       string    `xorm:"varchar(50) notnull"`
	Status     int8      `xorm:"tinyint default(0) notnull"`
	Created    time.Time `xorm:"timestamp created notnull"`
	Updated    time.Time `xorm:"timestamp updated  notnull"`
}
