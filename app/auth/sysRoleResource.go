package auth

import (
	"time"

	"github.com/xwinie/glue/core"
	"github.com/xwinie/glue/core/middleware/casbin"
)

//SysRoleResource 角色资源
type SysRoleResource struct {
	ID           int64     `xorm:"pk bigint 'id'"`
	RoleID       int64     `xorm:" bigint  notnull 'role_id'"`
	ResourceID   int64     `xorm:" bigint notnull 'resource_id'"`
	DeleteStatus int8      `xorm:"tinyint default(0) notnull"`
	Created      time.Time `xorm:"timestamp created notnull"`
	Updated      time.Time `xorm:"timestamp updated  notnull"`
	Locked       int8      `xorm:"tinyint default(0) notnull"`
}

type findRoleResource struct {
	code       string
	resourceID int64 `xorm:" 'resource_id'"`
	action     string
	method     string
	roleID     int64 `xorm:"'role_id'"`
}

func permissionByMultiRole(roleIds interface{}, resType int8) (resource []casbin.Permission, err error) {
	o := core.New()
	err = o.Table("sys_role_resource").Alias("rr").
		Join("INNER", []string{"sys_resource", "r"}, "r.id=rr.resource_id").
		And("r.res_type=?", resType).
		In("rr.role_id", roleIds).
		Cols("r.code", "r.action", "r.method").Find(&resource)
	return resource, err
}

func findResourceByMultiRole(roleIds []int64, resType int8) (resource []findRoleResource, err error) {
	o := core.New()
	err = o.Table("sys_role_resource").Alias("rr").Join("INNER", []string{"sys_resource", "r"}, "r.id=rr.resource_id").
		In("rr.role_id", roleIds).
		And("r.res_type=?", resType).
		Cols("r.code", "r.id resource_id", "r.action", "r.method", "rr.id role_Id").
		Find(&resource)
	return resource, err
}

func findResourceByRoleId(roleId int64) (resource []SysResource, err error) {
	o := core.New()
	err = o.Table("sys_role_resource").Alias("rr").Join("INNER", []string{"sys_resource", "r"}, "r.id=rr.resource_id").
		Where("rr.role_id=?", roleId).
		Cols("r.*").
		Find(&resource)
	return resource, err
}

func insertRoleResource(m []SysRoleResource) error {
	o := core.New()
	_, err := o.Insert(&m)
	return err
}

func deleteRoleResource(roleId int64) error {
	o := core.New()
	_, err := o.Where("role_id = ?", roleId).Delete(new(SysRoleResource))
	return err
}
