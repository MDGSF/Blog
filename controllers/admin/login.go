package admin

import (
	"github.com/MDGSF/Blog/controllers/base"
	"github.com/astaxie/beego"
)

// LoginController main controller
type LoginController struct {
	base.Controller
}

// Get main controller get
func (c *LoginController) Get() {
	beego.Info("LoginController get")

	c.TplName = "admin/login/index.html"
}
