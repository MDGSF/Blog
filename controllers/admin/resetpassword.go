package admin

import (
	"github.com/MDGSF/Blog/controllers/base"
	"github.com/astaxie/beego"
)

// ResetPasswordController controller
type ResetPasswordController struct {
	base.Controller
}

// Get controller get
func (c *ResetPasswordController) Get() {
	beego.Info("ResetPasswordController get")

	c.TplName = "admin/basic/resetpassword.html"
}
