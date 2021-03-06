package api

import (
	"github.com/MDGSF/Blog/controllers/base"
	"github.com/MDGSF/Blog/modules/models"
	"github.com/astaxie/beego"
)

// PostController main controller
type PostController struct {
	base.Controller
}

// Get main controller get
func (c *PostController) Get() {
	beego.Info("PostController get", c.Ctx.Input.Param(":splat"))

	postFileName := c.Ctx.Input.Param(":splat")
	stPost, ok := models.AllPostsFileName[postFileName]
	if !ok {
		return
	}

	c.TplName = "front/post.html"
	c.Data["Content"] = string(stPost.Content)
	c.Data["Title"] = stPost.Title
}
