package controllers

import (
	"encoding/json"
	"github.com/Herts/ray-ui/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserServerController struct {
	beego.Controller
}

func (c *UserServerController) ListAllUserServers() {

}

func (c *UserServerController) CreateUserServer() {
	var us models.UserServer
	logs.Debug(string(c.Ctx.Input.RequestBody))
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &us)
	if err != nil {
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}

	models.SaveUserServer(&us)
}
