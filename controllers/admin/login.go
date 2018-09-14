package admin

import (
	"github.com/MDGSF/Blog/controllers/base"
	"github.com/astaxie/beego"
)

// LoginController controller
type LoginController struct {
	base.Controller
}

// Get controller get
func (c *LoginController) Get() {
	beego.Info("LoginController get")

	c.TplName = "admin/login/index.html"
}

// Post controller Post
func (c *LoginController) Post() {
	beego.Info("LoginController post", c.Ctx.Input.Params(), c.Ctx.Request.Form, c.Ctx.Request.PostForm)

	userName := c.Ctx.Request.Form.Get("form-username")
	passWord := c.Ctx.Request.Form.Get("form-password")

	beego.Info("userName, passWord =", userName, passWord)

	c.TplName = "admin/login/index.html"
}
