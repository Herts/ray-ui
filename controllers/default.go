package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *MainController) GetTable() {
	c.Layout = "layout.html"
	c.TplName = "tables.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["SideBar"] = "sidebar.html"
	c.LayoutSections["PagePlugins"] = "datatable-script.html"
	c.LayoutSections["PageCss"] = "datatable-css.html"
	c.LayoutSections["PageCustomScripts"] = "datatable-userdata.html"
}
