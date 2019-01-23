package middleware

var ControllerMap map[string]string

func init() {
	ControllerMap = make(map[string]string)
	ControllerMap["/user"] = "用户操作"
	ControllerMap["/company"] = "单位操作"
	//ControllerMap["/menu"] = "获取菜单"
	ControllerMap["/rolemenu"] = "获取角色菜单"
	ControllerMap["/users"] = "用户操作"
	ControllerMap["/companys"] = "单位操作"
	ControllerMap["/roles"] = "角色操作"
	//ControllerMap["/data"] = "数据字典"
	ControllerMap["/notice"] = "通知公告"
	ControllerMap["/database"] = "查重数据库"
	ControllerMap["/file"] = "文件管理"
	ControllerMap["/system"] = "系统管理"
}
