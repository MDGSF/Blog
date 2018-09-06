package controllers

import (
	"github.com/astaxie/beego"
)

// LoginController login controller
type LoginController struct {
	beego.Controller
}

// Get login controller get
func (c *LoginController) Get() {
	beego.Info("LoginController get, c.Data =", c.Data)

	if c.Ctx.Request.Form == nil {
		c.Ctx.Request.ParseForm()
	}

	beego.Info("LoginController get, c.Ctx.Request.Form =", c.Ctx.Request.Form)

	c.TplName = "demo/login.tpl"
}

// Post login controller post
func (c *LoginController) Post() {
	beego.Info("LoginController post, c.Data =", c.Data)

	if c.Ctx.Request.Form == nil {
		c.Ctx.Request.ParseForm()
	}

	beego.Info("LoginController post, c.Ctx.Request.Form =", c.Ctx.Request.Form)

	// verify user name and password here.

	c.TplName = "index.tpl"
}
