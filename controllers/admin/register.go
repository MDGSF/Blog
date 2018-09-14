package admin

import (
	"github.com/MDGSF/Blog/controllers/base"
	"github.com/MDGSF/Blog/modules/auth"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// RegisterController controller
type RegisterController struct {
	base.Controller
}

// Get controller get
func (c *RegisterController) Get() {
	beego.Info("RegisterController get")

	c.TplName = "admin/login/index.html"
}

// Post controller Post
func (c *RegisterController) Post() {
	beego.Info("RegisterController post", c.Ctx.Input.Params(), c.Ctx.Request.Form, c.Ctx.Request.PostForm)

	username := c.Ctx.Request.Form.Get("form-username")
	password := c.Ctx.Request.Form.Get("form-password")
	email := c.Ctx.Request.Form.Get("form-email")

	beego.Info("username, password =", username, password)

	if len(username) == 0 || len(password) == 0 {
		beego.Error("Invalid username or password")
		return
	}

	valid := validation.Validation{}
	if v := valid.Email(email, "email"); !v.Ok {
		beego.Error("invalid email")
		return
	}

	if !auth.IsUserExist(username) {
		beego.Error("username already exist in db.")
		return
	}

	c.TplName = "admin/login/index.html"
}
