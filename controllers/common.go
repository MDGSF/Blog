package controllers

import (
	"github.com/astaxie/beego"
)

// CommonController main controller
type CommonController struct {
	beego.Controller
}

// Prepare main controller Prepare
func (c *CommonController) Prepare() {
	c.Layout = "HomeLayout.tpl"
}
