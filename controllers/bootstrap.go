package controllers

import (
	"github.com/astaxie/beego"
)

type BootstrapController struct {
	beego.Controller
}

func (c *BootstrapController) Get() {
	beego.Debug("BootstrapController get")

	c.TplName = "bootstrap.tpl"
}
