package base

import (
	"strings"
	"time"

	"github.com/MDGSF/Blog/modules/auth"
	"github.com/MDGSF/Blog/modules/models"
	"github.com/MDGSF/Blog/setting"
	"github.com/MDGSF/Blog/u"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

// Controller main controller
type Controller struct {
	beego.Controller
	i18n.Locale
	User    models.User
	IsLogin bool
}

// Prepare main controller Prepare
func (c *Controller) Prepare() {

	beego.Info("enter prepare")

	if c.Ctx.Request.Form == nil {
		c.Ctx.Request.ParseForm()
	}

	c.Data["PageStartTime"] = time.Now()

	c.StartSession()

	switch {
	case auth.GetUserFromSession(&c.User, c.CruSession):
		beego.Info("user in session =", c.User)
		c.IsLogin = true
	case auth.LoginUserFromRememberCookie(&c.User, c.Ctx):
		beego.Info("user in cookie =", c.User)
		c.IsLogin = true
	}

	beego.Info("prepare c.IsLogin =", c.IsLogin)

	if c.IsLogin {
		c.Data["User"] = &c.User
		c.Data["IsLogin"] = c.IsLogin

		if c.User.IsForbid {
			auth.LogoutUser(c.Ctx)
			//c.FlashRedirect() TODO
			return
		}
	}

	c.setLang()

	c.Data["AppName"] = setting.AppName
	c.Data["AppVersion"] = setting.AppVersion
	c.Data["AppAuthor"] = setting.AppAuthor
	c.Data["AppAuthorEmail"] = setting.AppAuthorEmail
	c.Data["AppAuthorGitHub"] = setting.AppAuthorGitHub
	c.Data["AppAuthorTwitter"] = setting.AppAuthorTwitter
	c.Data["AppAuthorFacebook"] = setting.AppAuthorFacebook
}

// Finish on router finished
func (c *Controller) Finish() {

}

func (c *Controller) setLangCookie(lang string) {
	c.Ctx.SetCookie("lang", lang, 60*60*24*365, "/", nil, nil, false)
}

// setLang sets site language.
func (c *Controller) setLang() bool {
	isNeedRedir := false
	hasCookie := false

	langs := setting.Langs

	// 1. check URL arguments.
	lang := c.GetString("lang")

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = c.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. check if isLogin then use user setting
	// if len(lang) == 0 && c.IsLogin {
	// 	lang = i18n.GetLangByIndex(c.User.Lang)
	// }

	// 4. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := c.Ctx.Input.Header("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 4. DefaucurLang language is English.
	if len(lang) == 0 {
		lang = "en-US"
		isNeedRedir = false
	}

	// Save language information in cookies.
	if !hasCookie {
		c.setLangCookie(lang)
	}

	// Set language properties.
	c.Data["Lang"] = lang
	c.Data["Langs"] = langs

	c.Lang = lang

	return isNeedRedir
}

// SetPaginator set paginator
func (c *Controller) SetPaginator(per int, nums int64) *u.Paginator {
	p := u.NewPaginator(c.Ctx.Request, per, nums)
	c.Data["paginator"] = p
	return p
}

// LoginUser ...
func (c *Controller) LoginUser(user *models.User, remember bool) string {
	loginRedirect := strings.TrimSpace(c.Ctx.GetCookie("login_to"))
	if u.IsMatchHost(loginRedirect) == false {
		loginRedirect = "/admin"
	} else {
		c.Ctx.SetCookie("login_to", "", -1, "/")
	}

	// login user
	auth.LoginUser(user, c.Ctx, remember)

	c.setLangCookie(i18n.GetLangByIndex(user.Lang))

	return loginRedirect
}
