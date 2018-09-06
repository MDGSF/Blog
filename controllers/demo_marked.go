package controllers

import (
	"io/ioutil"

	"github.com/astaxie/beego"
)

// MarkedController main controller
type MarkedController struct {
	beego.Controller
}

// Get main controller get
func (c *MarkedController) Get() {
	beego.Info("MarkedController get")

	mkPost, err := ioutil.ReadFile("test/post1.md")
	if err != nil {
		beego.Error("read failed, err =", err)
		return
	}

	c.Data["Post"] = string(mkPost)
	c.TplName = "demo/marked.tpl"
}
