package controllers

import (
	"encoding/json"
	"github.com/Herts/ray-ui/models"
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

// @router /list [get]
// @Success 200 {[]*models.RemoteServer} []*models.RemoteServer
func (c *ServerController) ListAllServers() {
	servers := models.GetAllServers()

	c.Data["json"] = response{Data: servers}
	c.ServeJSON()

}
