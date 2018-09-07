package controllers

import (
	"github.com/MDGSF/Blog/models"
	"github.com/astaxie/beego"
)

// IndexController main controller
type IndexController struct {
	CommonController
}

// Get main controller get
func (c *IndexController) Get() {
	beego.Info("IndexController get")

	c.TplName = "index.tpl"
	c.Data["Website"] = "MDGSF Blog"
	c.Data["Email"] = "1342042894@qq.com"
	c.Data["Author"] = "huangjian"
	c.Data["Posts"] = models.AllPosts
}
