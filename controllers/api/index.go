package api

import (
	"strconv"

	"github.com/MDGSF/Blog/controllers/base"
	"github.com/MDGSF/Blog/modules/models"
	"github.com/astaxie/beego"
)

// IndexController main controller
type IndexController struct {
	base.Controller
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

	c.TplName = "front/index.html"

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

	c.Data["Posts"] = models.AllPostsReverseByTime[start:end]
	c.Data["YearMonthArchives"] = models.MonthPosts
	c.Data["TagsArchives"] = models.PostsTagsManyPost
	c.Data["IsHome"] = true

	c.SetPaginator(pageLimit, int64(pageCount))
}
