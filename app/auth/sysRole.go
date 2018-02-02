package auth

import "time"

//SysRole 角色
type SysRole struct {
	ID           int64     `xorm:"pk bigint 'id'"`
	Code         string    `xorm:"varchar(100) unique notnull"`
	Name         string    `xorm:"varchar(200)  notnull"`
	Description  string    `xorm:"varchar(255)"`
	DeleteStatus int8      `xorm:"tinyint default(0) notnull"`
	Created      time.Time `xorm:"timestamp created notnull"`
	Updated      time.Time `xorm:"timestamp updated  notnull"`
	Locked       int8      `xorm:"tinyint default(0) notnull"`
}
