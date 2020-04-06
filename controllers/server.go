package controllers

import (
	"../models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type ServerController struct {
	beego.Controller
}

func (c *ServerController) CreateServer() {
	var newServer models.RemoteServer
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &newServer)
	if err != nil {
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}

}

func (c *ServerController) ListAllServers() {
	servers := models.GetAllServers()

	c.Data["json"] = response{Data: servers}
	c.ServeJSON()

}
