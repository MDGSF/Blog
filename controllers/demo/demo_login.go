package demo

import (
	"github.com/MDGSF/Blog/u"
	"github.com/astaxie/beego"
)

// LoginController login controller
type LoginController struct {
	beego.Controller
}

// Get login controller get
func (c *LoginController) Get() {
	if c.Ctx.Request.Form == nil {
		c.Ctx.Request.ParseForm()
	}

	beego.Info("LoginController get, c.Data =", c.Data)
	beego.Info("LoginController get, c.Ctx.Request.Form =", c.Ctx.Request.Form)

	c.TplName = "demo/login.tpl"
}

// Post login controller post
func (c *LoginController) Post() {
	if c.Ctx.Request.Form == nil {
		c.Ctx.Request.ParseForm()
	}

	beego.Info("LoginController post, c.Data =", c.Data)
	beego.Info("LoginController post, c.Ctx.Request.Form =", c.Ctx.Request.Form)

	// verify user name and password here.

	sess, _ := u.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	//username := sess.Get("username")
	sess.Set("username", c.Ctx.Request.Form["username"])

	c.TplName = "index.tpl"
}
