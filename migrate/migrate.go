package migrate

import (
	"log"

	"github.com/xormplus/xorm"
	"github.com/xwinie/glue/app/auth"
	"github.com/xwinie/glue/core"
)

//Migrate 自动升级
func Migrate(o *xorm.Engine) {
	err := o.Sync2(new(auth.SysClient),
		new(auth.SysUser),
		new(auth.SysRole),
		new(auth.SysResource),
		new(auth.SysUserRole),
		new(auth.SysRoleResource))
	if err != nil {
		log.Fatal("init db error:", err.Error())
	}
	initAuthData()
}

func initAuthData() {
	G, _ := core.NewGUID(1)
	id, _ := G.NextID()
	client := new(auth.SysClient)
	client.ID = id
	client.ClientID = "app1"
	client.Name = "测试app"
	client.Secret = "Lx1b8JoZoE"
	client.VerifySecret = "Lx1b8JoZoE"
	user := new(auth.SysUser)
	userID, _ := G.NextID()
	user.ID = userID
	salt := core.RandStringByLen(6)
	user.Account = "12345"
	user.Name = "测试员工"
	user.Password = core.Md5(core.Md5(core.Sha1("12345")+core.Sha1("Password")) + salt)
	user.Salt = salt
	role := new(auth.SysRole)
	RoleID, _ := G.NextID()
	role.ID = RoleID
	role.Code = "1"
	role.Name = "管理员"
	userRole := new(auth.SysUserRole)
	userRole.RoleID = RoleID
	userRole.UserID = userID

	resource := []auth.SysResource{
		{ID: 10000, Code: "10000", Action: "/v1/login", Method: "POST", Name: "用户登录", IsOpen: 1},
		{ID: 10001, Code: "10001", Action: "/v1/user", Method: "POST", Name: "添加用户"},
		{ID: 10002, Code: "10002", Action: "/v1/user/:id", Method: "DELETE", Name: "删除用户"},
		{ID: 10003, Code: "10003", Action: "/v1/user/:id", Method: "PUT", Name: "修改用户"},
		{ID: 10004, Code: "10004", Action: "/v1/user/:account", Method: "GET", Name: "根据账号获取用户"},
		{ID: 10005, Code: "10005", Action: "/v1/user/:id/role", Method: "POST", Name: "给用户分配角色"},
		{ID: 10006, Code: "10006", Action: "/v1/role", Method: "POST", Name: "添加角色"},
		{ID: 10007, Code: "10007", Action: "/v1/role/:id", Method: "DELETE", Name: "删除角色"},
		{ID: 10008, Code: "10008", Action: "/v1/role/:id", Method: "PUT", Name: "修改角色"},
		{ID: 10009, Code: "10009", Action: "/v1/role/:code", Method: "GET", Name: "根据编码获取角色"},
		{ID: 10010, Code: "10010", Action: "/v1/role/:id/resource", Method: "POST", Name: "给角色分配资源"},
		{ID: 10011, Code: "10011", Action: "/v1/resource", Method: "POST", Name: "添加资源"},
		{ID: 10012, Code: "10012", Action: "/v1/resource/:id", Method: "DELETE", Name: "删除资源"},
		{ID: 10013, Code: "10013", Action: "/v1/resource/:id", Method: "PUT", Name: "修改资源"},
		{ID: 10014, Code: "10014", Action: "/v1/resource/:code", Method: "GET", Name: "根据编码获取资源"},
		{ID: 10015, Code: "10015", Action: "/v1/client", Method: "POST", Name: "添加应用"},
		{ID: 10016, Code: "10016", Action: "/v1/client/:id", Method: "DELETE", Name: "删除应用"},
		{ID: 10017, Code: "10017", Action: "/v1/client/:id", Method: "PUT", Name: "修改应用"},
		{ID: 10018, Code: "10018", Action: "/v1/client/:clientId", Method: "GET", Name: "根据应用id获取应用"},
		{ID: 10019, Code: "10019", Action: "/v1/user/:id", Method: "PUT", Name: "修改用户"},
		{ID: 10020, Code: "10020", Action: "/v1/role/:id/resource", Method: "GET", Name: "根据角色ID获取资源信息"},
		{ID: 10021, Code: "10021", Action: "/v1/user/:id/role", Method: "GET", Name: "根据ID获取角色信息"},
		{ID: 10022, Code: "10022", Action: "/v1/resource", Method: "GET", Name: "分页获取所有资源"},
		{ID: 10023, Code: "10023", Action: "/v1/user", Method: "GET", Name: "分页获取所有用户"},
		{ID: 10024, Code: "10024", Action: "/v1/role", Method: "GET", Name: "分页获取所有角色"},
		{ID: 10025, Code: "10025", Action: "/v1/dict", Method: "POST", Name: "添加数据字典"},
		{ID: 10026, Code: "10026", Action: "/v1/dict/:id", Method: "DELETE", Name: "删除数据字典"},
		{ID: 10027, Code: "10027", Action: "/v1/dict/:id", Method: "PUT", Name: "修改数据字典"},
		{ID: 10028, Code: "10028", Action: "/v1/dict", Method: "GET", Name: "根据分页获取数据字典"},
		{ID: 10029, Code: "10029", Action: "/v1/dict/:id", Method: "GET", Name: "根据ID获取数据字典"},
		{ID: 10030, Code: "10030", Action: "/v1/upload", Method: "POST", Name: "上传文件"},
		{ID: 10031, Code: "10031", Action: "/static", Method: "GET", Name: "获取文件"},
		{ID: 10032, Code: "10032", Action: "/images", Method: "GET", Name: "获取图片"},
		{ID: 10033, Code: "10033", Action: "/v1/menus/:userId", Method: "GET", Name: "根据用户获取菜单信息"},
		{ID: 10034, Code: "10034", Action: "/user", Method: "", Name: "用户列表", ResType: 1},
		{ID: 10035, Code: "10035", Action: "/role", Method: "", Name: "角色列表", ResType: 1},
		{ID: 10036, Code: "10036", Action: "/client", Method: "", Name: "客户端管理", ResType: 1},
		{ID: 10037, Code: "10037", Action: "/v1/app", Method: "GET", Name: "获取当前用户的应用"},
		{ID: 10038, Code: "10038", Action: "/v1/client", Method: "GET", Name: "获取当前用户的应用"},
	}
	roleResource := []auth.SysRoleResource{
		{ID: 10001, RoleID: RoleID, ResourceID: 10001},
		{ID: 10002, RoleID: RoleID, ResourceID: 10002},
		{ID: 10003, RoleID: RoleID, ResourceID: 10003},
		{ID: 10004, RoleID: RoleID, ResourceID: 10004},
		{ID: 10005, RoleID: RoleID, ResourceID: 10005},
		{ID: 10006, RoleID: RoleID, ResourceID: 10006},
		{ID: 10007, RoleID: RoleID, ResourceID: 10007},
		{ID: 10008, RoleID: RoleID, ResourceID: 10008},
		{ID: 10009, RoleID: RoleID, ResourceID: 10009},
		{ID: 10010, RoleID: RoleID, ResourceID: 10010},
		{ID: 10011, RoleID: RoleID, ResourceID: 10011},
		{ID: 10012, RoleID: RoleID, ResourceID: 10012},
		{ID: 10013, RoleID: RoleID, ResourceID: 10013},
		{ID: 10014, RoleID: RoleID, ResourceID: 10014},
		{ID: 10015, RoleID: RoleID, ResourceID: 10015},
		{ID: 10016, RoleID: RoleID, ResourceID: 10016},
		{ID: 10017, RoleID: RoleID, ResourceID: 10017},
		{ID: 10018, RoleID: RoleID, ResourceID: 10018},
		{ID: 10019, RoleID: RoleID, ResourceID: 10019},
		{ID: 10020, RoleID: RoleID, ResourceID: 10020},
		{ID: 10021, RoleID: RoleID, ResourceID: 10021},
		{ID: 10022, RoleID: RoleID, ResourceID: 10022},
		{ID: 10023, RoleID: RoleID, ResourceID: 10023},
		{ID: 10024, RoleID: RoleID, ResourceID: 10024},
		{ID: 10025, RoleID: RoleID, ResourceID: 10025},
		{ID: 10026, RoleID: RoleID, ResourceID: 10026},
		{ID: 10027, RoleID: RoleID, ResourceID: 10027},
		{ID: 10028, RoleID: RoleID, ResourceID: 10028},
		{ID: 10029, RoleID: RoleID, ResourceID: 10029},
		{ID: 10030, RoleID: RoleID, ResourceID: 10030},
		{ID: 10031, RoleID: RoleID, ResourceID: 10031},
		{ID: 10032, RoleID: RoleID, ResourceID: 10032},
		{ID: 10033, RoleID: RoleID, ResourceID: 10033},
	}
	o := core.New()
	o.Insert(client)
	o.Insert(user)
	o.Insert(role)
	o.Insert(userRole)
	o.Insert(&resource)
	o.Insert(&roleResource)
}
