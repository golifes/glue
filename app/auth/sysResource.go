package auth

import (
	"time"

	"github.com/xwinie/glue/core"
	"github.com/xwinie/glue/core/middleware/casbin"
)

//SysResource 资源
type SysResource struct {
	ID           string    `xorm:"pk bigint 'id'" json:"Id"`
	Code         string    `xorm:"varchar(100) unique notnull"`
	Name         string    `xorm:"varchar(100) notnull"`
	Action       string    `xorm:"varchar(100) notnull"`
	Method       string    `xorm:"varchar(100)"`
	IsOpen       int8      `xorm:"tinyint default(0) notnull"` //0非开放1开放
	ResType      int8      `xorm:"tinyint default(0) notnull"` //0代表是接口1代表菜单
	ParentID     string    `xorm:"bigint 'parent_id' default(0) notnull"`
	DeleteStatus int8      `xorm:"tinyint default(0) notnull"`
	Created      time.Time `xorm:"datetime created notnull"`
	Updated      time.Time `xorm:"timestamp updated notnull"`
}

//OpenPermission 根据角色获取权限
func openPermission() ([]casbin.Permission, error) {
	permiss := new([]casbin.Permission)
	o := core.New()
	err := o.Table(SysResource{}).Cols("code", "action", "method").Where("is_open = ?", 1).Find(permiss)
	return *permiss, err
}

func resourceCountByPage() (num int64, err error) {
	o := core.New()
	num, err = o.Table("sys_resource").Count()
	return num, err
}

func resourceByPage(pageSize int, offset int) (m []*SysResource, err error) {
	o := core.New()
	err = o.Table("sys_resource").Limit(pageSize, offset).Find(&m)
	return m, err
}
func findResourceByCode(code string) (m SysResource, err error) {
	o := core.New()
	_, err = o.Where("code = ?", code).Get(&m)
	return m, err
}
