package api

import (
	"github.com/MDGSF/Blog/controllers/base"
	"github.com/MDGSF/Blog/modules/models"
	"github.com/astaxie/beego"
)

// ClassificationController main controller
type ClassificationController struct {
	base.Controller
}

// Get main controller get
func (c *ClassificationController) Get() {
	beego.Info("ClassificationController get")

	c.TplName = "front/classification.html"
	c.Data["IsCategory"] = true
	c.Data["classificationCategory"] = models.PostsCategory
	c.Data["classificationYearMonth"] = models.MonthPosts
}
