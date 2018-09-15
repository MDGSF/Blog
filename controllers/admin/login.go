package admin

import (
	"github.com/MDGSF/Blog/controllers/base"
	"github.com/MDGSF/Blog/modules/auth"
	"github.com/astaxie/beego"
)

// LoginController controller
type LoginController struct {
	base.Controller
}

// Get controller get
func (c *LoginController) Get() {
	beego.Info("LoginController get")

	c.TplName = "admin/basic/login.html"
}

// Post controller Post
func (c *LoginController) Post() {
	beego.Info("LoginController post", c.Ctx.Input.Params(), c.Ctx.Request.Form, c.Ctx.Request.PostForm)

	username := c.Ctx.Request.Form.Get("form-username")
	password := c.Ctx.Request.Form.Get("form-password")

	beego.Info("username, password =", username, password)

	if len(username) == 0 || len(password) == 0 {
		beego.Error("Invalid username or password")
		return
	}

	if !auth.IsUserExist(username) {
		beego.Error("username already exist in db.")
		return
	}

	c.TplName = "admin/basic/login.html"
}
