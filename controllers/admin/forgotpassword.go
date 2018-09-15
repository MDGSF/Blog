package admin

import (
	"github.com/MDGSF/Blog/controllers/base"
	"github.com/astaxie/beego"
)

// ForgotPasswordController controller
type ForgotPasswordController struct {
	base.Controller
}

// Get controller get
func (c *ForgotPasswordController) Get() {
	beego.Info("ForgotPasswordController get")

	c.TplName = "admin/basic/forgotpassword.html"
}
