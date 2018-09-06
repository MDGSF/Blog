package controllers

import (
	"github.com/astaxie/beego"
)

// MainController main controller
type MainController struct {
	beego.Controller
}

// Init main controller init
// func (c *MainController) Init(ct *context.Context, controllerName, actionName string, app interface{}) {
// 	beego.Info("MainController Init")
// }

// Prepare main controller Prepare
func (c *MainController) Prepare() {
	beego.Info("MainController Prepare")
}

// Get main controller get
func (c *MainController) Get() {
	beego.Info("MainController get")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// Post main controller Post
func (c *MainController) Post() {
	beego.Info("MainController Post")
}

// Delete main controller Delete
func (c *MainController) Delete() {
	beego.Info("MainController Delete")
}

// Put main controller Put
func (c *MainController) Put() {
	beego.Info("MainController Put")
}

// Head main controller Head
func (c *MainController) Head() {
	beego.Info("MainController Head")
}

// Patch main controller Patch
func (c *MainController) Patch() {
	beego.Info("MainController Patch")
}

// Options main controller Options
func (c *MainController) Options() {
	beego.Info("MainController Options")
}

// Finish main controller Finish
func (c *MainController) Finish() {
	beego.Info("MainController Finish")
}
