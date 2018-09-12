package api

import (
	"github.com/MDGSF/Blog/controllers/base"
	"github.com/astaxie/beego"
)

// CategoriesController main controller
type CategoriesController struct {
	base.Controller
}

// Get main controller get
func (c *CategoriesController) Get() {
	beego.Info("CategoriesController get")

	c.TplName = "front/categories.html"
	c.Data["IsCategory"] = true
}
