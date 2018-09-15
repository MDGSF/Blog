package admin

import (
	"fmt"

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

	c.TplName = "admin/gentelella-1.4.0/production/register.html"
}

// Post controller Post
func (c *RegisterController) Post() {
	beego.Info("RegisterController post", c.Ctx.Input.Params(), c.Ctx.Request.Form, c.Ctx.Request.PostForm)

	username := c.Ctx.Request.Form.Get("name")
	password := c.Ctx.Request.Form.Get("password")
	email := c.Ctx.Request.Form.Get("email")

	beego.Info("username, password, email =", username, password, email)

	if len(username) == 0 || len(password) == 0 {
		beego.Error("Invalid username or password")
		return
	}

	valid := validation.Validation{}
	if v := valid.Email(email, "email"); !v.Ok {
		beego.Error("invalid email")
		return
	}

	if auth.IsUserExist(username) {
		strError := "username already exist in db."
		beego.Error(strError)
		c.TplName = "admin/basic/errormsg.html"
		c.Data["error"] = strError
		return
	}

	if auth.IsEmailExist(email) {
		strError := "email already exist in db."
		beego.Error(strError)
		c.TplName = "admin/basic/errormsg.html"
		c.Data["error"] = strError
		return
	}

	if err := auth.RegisterUser(username, password, email); err != nil {
		strError := fmt.Sprintf("register to db failed, err = %v", err)
		beego.Error(strError)
		c.TplName = "admin/basic/errormsg.html"
		c.Data["error"] = "register user failed."
		return
	}

	// c.Redirect("admin/gentelella-1.4.0/production/login.html", 301)
	c.TplName = "admin/gentelella-1.4.0/production/login.html"
}
