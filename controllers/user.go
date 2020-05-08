package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Herts/ray-ui/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
)

type UserController struct {
	beego.Controller
}

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// @Param body body true models.User "New user"
// @router /create [post]
// @Success 200 {response} response
func (c *UserController) CreateUser() {
	var newUser models.User
	logs.Debug(string(c.Ctx.Input.RequestBody))
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &newUser)
	if err != nil {
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}
	newUser.Email = strings.ReplaceAll(newUser.Email, " ", "")
	u := models.GetUser(newUser.Email)
	if u.Email == newUser.Email {
		c.Data["json"] = response{Message: fmt.Sprintf("User email %s exists", newUser.Email)}
		c.ServeJSON()
		return
	}

	if len(newUser.UserId) != 36 {
		c.Data["json"] = response{
			Message: "User id is not correct",
		}
		c.ServeJSON()
		return
	}

	models.AddUser(&newUser)
	c.Data["json"] = response{Message: fmt.Sprintf("User %s creation success", newUser.Email)}
	c.ServeJSON()
}

// @router /listData [get]
// @Success 200 {response} response
func (c *UserController) ListAllUserData() {
	uds := models.GetAllUserData()
	tableData := [][]interface{}{}
	for _, ud := range uds {
		tableData = append(tableData, []interface{}{ud.Region, ud.Index, ud.Email, ud.Date,
			ud.UpDataConsumed, ud.DownDataConsumed})
	}
	c.Data["json"] = response{Data: tableData}
	c.ServeJSON()
}
