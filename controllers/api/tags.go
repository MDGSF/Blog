package api

import (
	"github.com/MDGSF/Blog/controllers/base"
	"github.com/MDGSF/Blog/models"
	"github.com/astaxie/beego"
)

// TagsController main controller
type TagsController struct {
	base.Controller
}

// Get main controller get
func (c *TagsController) Get() {

	beego.Info("TagsController get", c.Ctx.Input.Params(), c.Ctx.Request.Form, c.Ctx.Request.PostForm)

	tagName := c.Ctx.Request.Form.Get("name")
	if len(tagName) == 0 {
		beego.Error("no tag name")
		return
	}

	postsArr, ok := models.AllPostsTags[tagName]
	if !ok {
		beego.Error("invalid tag name =", tagName)
		return
	}

	c.TplName = "front/listArticleTitle.html"
	c.Data["TitleName"] = tagName
	c.Data["Articles"] = postsArr

	c.Data["sidebarname"] = "Tags"
	c.Data["sideContentURL"] = "/tags"
	c.Data["sideArchives"] = models.AllPostsTags
}
