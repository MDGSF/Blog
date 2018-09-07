package controllers

import (
	"github.com/astaxie/beego"
)

// CommonController main controller
type CommonController struct {
	beego.Controller
}

// Prepare main controller Prepare
func (c *CommonController) Prepare() {

	if c.Ctx.Request.Form == nil {
		c.Ctx.Request.ParseForm()
	}

	c.Layout = "HomeLayout.tpl"
	c.Data["Website"] = "MDGSF Blog"
	c.Data["Email"] = "1342042894@qq.com"
	c.Data["Author"] = "huangjian"
}
