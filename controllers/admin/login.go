package admin

import (
	"github.com/MDGSF/Blog/controllers/base"
	"github.com/MDGSF/Blog/modules/auth"
	"github.com/astaxie/beego"
)

// LoginController controller
type LoginController struct {
	base.Controller
}

// Get controller get
func (c *LoginController) Get() {
	beego.Info("LoginController get")

	c.Data["IsLoginPage"] = true
	c.TplName = "admin/gentelella-1.4.0/production/login.html"

	//loginRedirect := strings.TrimSpace(c.GetString("to"))

}

// Login controller Post
func (c *LoginController) Login() {
	beego.Info("LoginController Login", c.Ctx.Input.Params(), c.Ctx.Request.Form, c.Ctx.Request.PostForm)

	username := c.Ctx.Request.Form.Get("name")
	password := c.Ctx.Request.Form.Get("password")

	beego.Info("username, password =", username, password)

	if len(username) == 0 || len(password) == 0 {
		beego.Error("Invalid username or password")
		return
	}

	if !auth.IsUserExist(username) {
		strError := "username not exist in db."
		beego.Error(strError)
		c.TplName = "admin/basic/errormsg.html"
		c.Data["error"] = strError
		return
	}

	c.SetSession("loginuser", username)
	beego.Info("current session =", c.GetSession("loginuser"), c.CruSession)

	// 抓包看下就知道了。
	// 直接用 TplName，就是服务器把这个 index.html 的页面直接发送给了客户端。
	// 但是这里用了模板，所以直接发送的话，有的东西可能就没有处理。
	// c.TplName = "admin/gentelella-1.4.0/production/index.html"

	// 用 redirect 的话，客户端会收到 302，然后用 /admin 重新向服务器发送请求。
	// 这里会走整个完整的请求流程。
	c.Redirect("/admin", 302)
}

// Logout implemented user logout page.
func (c *LoginController) Logout() {
	auth.LogoutUser(c.Ctx)

	c.FlashWrite("HasLogout", "true")

	c.Redirect("/admin/login", 302)
}
