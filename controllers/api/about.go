package api

import (
	"github.com/MDGSF/Blog/controllers/base"
	"github.com/astaxie/beego"
)

// AboutController main controller
type AboutController struct {
	base.Controller
}

// Get main controller get
func (c *AboutController) Get() {
	beego.Info("AboutController get")

	c.TplName = "front/about.html"
	c.Data["IsAbout"] = true
}
