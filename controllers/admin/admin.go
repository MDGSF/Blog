package admin

import (
	"github.com/MDGSF/Blog/controllers/base"
	"github.com/astaxie/beego"
)

// AdminController controller
type AdminController struct {
	base.Controller
}

// Get controller get
func (c *AdminController) Get() {
	beego.Info("AdminController get")

	c.TplName = "admin/gentelella-1.4.0/production/index.html"
}
