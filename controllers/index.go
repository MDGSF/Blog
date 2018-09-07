package controllers

import (
	"strconv"

	"github.com/MDGSF/Blog/models"
	"github.com/astaxie/beego"
)

// IndexController main controller
type IndexController struct {
	CommonController
}

// Get main controller get
func (c *IndexController) Get() {
	beego.Info("IndexController get", c.Ctx.Input.Params(), c.Ctx.Request.Form, c.Ctx.Request.PostForm)

	var pageCount int
	var pageLimit int
	var curPage int

	strPageCount := c.Ctx.Request.Form.Get("pageCount")
	if len(strPageCount) > 0 {
		clientPageCount, _ := strconv.Atoi(strPageCount)
		if clientPageCount != len(models.AllPosts) {
			beego.Error("invalid page count from client")
		}
		pageCount = len(models.AllPosts)
	} else {
		pageCount = len(models.AllPosts)
	}

	strPageLimit := c.Ctx.Request.Form.Get("pageLimit")
	if len(strPageLimit) > 0 {
		clientPageLimit, _ := strconv.Atoi(strPageLimit)
		if clientPageLimit < 0 {
			pageLimit = 10
		} else {
			pageLimit = clientPageLimit
		}
	} else {
		pageLimit = 10
	}

	strCurPage := c.Ctx.Request.Form.Get("curPage")
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

	c.TplName = "index.tpl"
	c.Data["PageCount"] = pageCount
	c.Data["PageLimit"] = pageLimit
	c.Data["CurPage"] = curPage

	start := curPage * pageLimit
	end := start + pageLimit
	if start > pageCount {
		beego.Error(pageCount, pageLimit, curPage, start, end)
		return
	}
	if end > pageCount {
		end = pageCount
	}

	beego.Info(pageCount, pageLimit, curPage, start, end)

	c.Data["Posts"] = models.AllPosts[start:end]
}
