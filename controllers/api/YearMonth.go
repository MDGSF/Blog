package api

import (
	"github.com/MDGSF/Blog/controllers/base"
	"github.com/MDGSF/Blog/models"
	"github.com/astaxie/beego"
)

// YearMonthController main controller
type YearMonthController struct {
	base.Controller
}

// Get main controller get
func (c *YearMonthController) Get() {

	beego.Info("YearMonthController get", c.Ctx.Input.Params(), c.Ctx.Request.Form, c.Ctx.Request.PostForm)

	YearMonthName := c.Ctx.Request.Form.Get("yearmonthname")
	if len(YearMonthName) == 0 {
		beego.Error("no tag name")
		return
	}

	postsArr, ok := models.MonthPosts[YearMonthName]
	if !ok {
		beego.Error("invalid tag name =", YearMonthName)
		return
	}

	c.TplName = "front/listArticleTitle.html"
	c.Data["TitleName"] = YearMonthName
	c.Data["Articles"] = postsArr
}
