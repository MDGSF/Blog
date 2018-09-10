package controllers

import (
	"strconv"

	"github.com/MDGSF/Blog/models"
	"github.com/MDGSF/Blog/u"
	"github.com/astaxie/beego"
)

// IndexController main controller
type IndexController struct {
	CommonController
}

// Get main controller get
func (c *IndexController) Get() {
	beego.Info("IndexController get", c.Ctx.Input.Params(), c.Ctx.Request.Form, c.Ctx.Request.PostForm)

	pageCount := len(models.AllPosts)
	pageLimit := 10
	var curPage int
	var curPageIndex int

	strCurPage := c.Ctx.Request.Form.Get("p")
	if len(strCurPage) > 0 {
		clientCurPage, _ := strconv.Atoi(strCurPage)
		if clientCurPage < 0 {
			curPage = 0
		} else {
			curPage = clientCurPage
		}
	} else {
		curPage = 0
	}

	if curPage == 0 {
		curPageIndex = 0
	} else {
		curPageIndex = curPage - 1
	}

	c.TplName = "index.tpl"

	start := curPageIndex * pageLimit
	end := start + pageLimit
	if start > pageCount {
		beego.Error(pageCount, pageLimit, curPageIndex, start, end)
		return
	}
	if end > pageCount {
		end = pageCount
	}

	beego.Info(pageCount, pageLimit, curPageIndex, curPage, start, end)

	c.Data["Posts"] = models.AllPosts[start:end]

	p := u.NewPaginator(c.Ctx.Request, pageLimit, pageCount)
	c.Data["paginator"] = p
}
