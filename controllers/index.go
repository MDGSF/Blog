package controllers

import (
	"github.com/astaxie/beego"
)

// IndexController main controller
type IndexController struct {
	beego.Controller
}

// Get main controller get
func (c *IndexController) Get() {
	beego.Info("IndexController get")
	c.Data["Website"] = "MDGSF Blog"
	c.Data["Email"] = "1342042894@qq.com"
	c.TplName = "index.tpl"
}
