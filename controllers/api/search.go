package api

import (
	"strings"

	"github.com/MDGSF/Blog/models"
	"github.com/astaxie/beego"
)

// SearchController search controller
type SearchController struct {
	beego.Controller
}

// Post search controller post
func (c *SearchController) Post() {

	r := c.Ctx.Request
	r.ParseForm()
	beego.Info("SearchController get", c.Ctx.Input.Params(), r.Form, r.PostForm, r.URL.Path, r.URL.Scheme)

	searchContent := r.Form.Get("searchContent")
	if len(searchContent) == 0 {
		beego.Error("no search content")
		return
	}

	postsArr := make([]*models.TPost, 0)

	for k, v := range models.AllPosts {
		if strings.Contains(strings.ToLower(v.Title), strings.ToLower(searchContent)) {
			postsArr = append(postsArr, models.AllPosts[k])
		}
	}

	c.TplName = "front/listArticleTitle.html"
	c.Data["TitleName"] = searchContent
	c.Data["Articles"] = postsArr
}
