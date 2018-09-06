package controllers

import (
	"github.com/astaxie/beego"
)

// MainController main controller
type MainController struct {
	beego.Controller
}

// Get main controller get
func (c *MainController) Get() {
	beego.Info("MainController get")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
