package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/user/add", &controllers.UserController{}, "post:CreateUser")

	beego.Router("/api/server/listAll", &controllers.ServerController{}, "get:ListAllServers")
}
