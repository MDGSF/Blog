package controllers

import (
	"github.com/astaxie/beego"
)

// BootstrapController bootstrap controller
type BootstrapController struct {
	beego.Controller
}

// Get bootstrap controller get
func (c *BootstrapController) Get() {
	beego.Debug("BootstrapController get")

	c.TplName = "bootstrap.tpl"
}
