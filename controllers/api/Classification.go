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

	categorytype := c.GetString("categorytype")
	if categorytype == "postcategory" {
		c.Data["IsPostCategory"] = true
		c.Data["classificationCategory"] = models.PostsCategory
	} else if categorytype == "monthcategory" {
		c.Data["IsMonthCategory"] = true
		c.Data["classificationYearMonth"] = models.MonthPosts
	} else {
		c.Data["IsPostCategory"] = true
		c.Data["classificationCategory"] = models.PostsCategory
	}

	c.TplName = "front/classification.html"
	c.Data["IsCategory"] = true
}
