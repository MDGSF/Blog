package controllers

import (
	"github.com/astaxie/beego"
)

type LayuiController struct {
	beego.Controller
}

func (c *LayuiController) Get() {
	beego.Debug("LayuiController get")

	c.TplName = "layui.tpl"
}
