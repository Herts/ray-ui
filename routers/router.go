// @APIVersion 1.0.0
// @Title ray-ui API
package routers

import (
	"github.com/Herts/ray-ui/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/server",
			beego.NSInclude(&controllers.ServerController{}),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(&controllers.UserController{}),
		),
		beego.NSNamespace("/config",
			beego.NSInclude(&controllers.V2RayConfigController{}),
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/index.html", &controllers.MainController{})
	beego.Router("/tables.html", &controllers.MainController{}, "get:GetTable")
	//beego.Router("/api/user/add", &controllers.UserController{}, "post:CreateUser")
	//
	//beego.Router("/api/server/listAll", &controllers.ServerController{}, "get:ListAllServers")
	//beego.Router("/", &controllers.UserServerController{})
}
