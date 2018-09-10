package controllers

import (
	"github.com/astaxie/beego"
)

// LayuiController layui controller
type LayuiController struct {
	beego.Controller
}

// Get layui controller get
func (c *LayuiController) Get() {
	beego.Info("LayuiController get, c.Data =", c.Data)
	c.TplName = "demo/layui.tpl"
}
