package controllers

import (
	"github.com/MDGSF/Blog/models"
	"github.com/astaxie/beego"
)

// PostController main controller
type PostController struct {
	CommonController
}

// Get main controller get
func (c *PostController) Get() {
	beego.Info("PostController get", c.Ctx.Input.Param(":splat"))

	postFileName := c.Ctx.Input.Param(":splat")
	stPost, ok := models.AllPostsFileName[postFileName]
	if !ok {
		return
	}

	c.TplName = "post.tpl"
	c.Data["Content"] = string(stPost.Content)
}
