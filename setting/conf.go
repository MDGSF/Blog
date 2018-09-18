package setting

import (
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

var (
	AppName  string
	HTTPPort string
	RunMode  string

	EnableAdmin bool
	AdminAddr   string
	AdminPort   string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassWord string
	DBName     string

	AppVersion        string
	AppAuthor         string
	AppAuthorEmail    string
	AppAuthorGitHub   string
	AppAuthorTwitter  string
	AppAuthorFacebook string

	PostDirectory      []string
	PostAuthor         string
	PostAbstractionLen int

	LoginRememberDays int

	CookieRememberName string
	CookieUserName     string
)

var (
	Langs []string
)

// LoadConfig load conf/app.conf
func LoadConfig() {

	settingGlobalVariables()

	settingLocales()
}

func settingGlobalVariables() {
	AppName = beego.AppConfig.DefaultString("appname", "blog")
	HTTPPort = beego.AppConfig.DefaultString("httpport", "8080")
	RunMode = beego.AppConfig.DefaultString("runmode", "dev")

	EnableAdmin = beego.AppConfig.DefaultBool("EnableAdmin", false)
	AdminAddr = beego.AppConfig.DefaultString("AdminAddr", "localhost")
	AdminPort = beego.AppConfig.DefaultString("AdminPort", "8088")

	DBHost = beego.AppConfig.DefaultString("DBHost", "localhost")
	DBPort = beego.AppConfig.DefaultString("DBPort", "3306")
	DBUser = beego.AppConfig.DefaultString("DBUser", "root")
	DBPassWord = beego.AppConfig.DefaultString("DBPassWord", "123456")
	DBName = beego.AppConfig.DefaultString("DBName", "db_name")

	AppVersion = beego.AppConfig.DefaultString("AppVersion", "1.0.0.1")
	AppAuthor = beego.AppConfig.DefaultString("AppAuthor", "author")
	AppAuthorEmail = beego.AppConfig.DefaultString("AppAuthorEmail", "xxx@email.com")
	AppAuthorGitHub = beego.AppConfig.DefaultString("AppAuthorGitHub", "")
	AppAuthorTwitter = beego.AppConfig.DefaultString("AppAuthorTwitter", "")
	AppAuthorFacebook = beego.AppConfig.DefaultString("AppAuthorFacebook", "")

	PostDirectory = beego.AppConfig.Strings("PostDirectory")
	PostAuthor = beego.AppConfig.DefaultString("PostAuthor", "author")
	PostAbstractionLen = beego.AppConfig.DefaultInt("PostAbstractionLen", 100)

	LoginRememberDays = beego.AppConfig.DefaultInt("login_remember_days", 7)

	CookieRememberName = beego.AppConfig.DefaultString("cookie_remember_name", "blog_remember_name")
	CookieUserName = beego.AppConfig.DefaultString("cookie_user_name", "blog_user_name")

	beego.Info("AppName =", AppName)
	beego.Info("HTTPPort =", HTTPPort)
	beego.Info("RunMode =", RunMode)

	beego.Info("EnableAdmin =", EnableAdmin)
	beego.Info("AdminAddr =", AdminAddr)
	beego.Info("AdminPort =", AdminPort)

	beego.Info("AppVersion =", AppVersion)
	beego.Info("AppAuthor =", AppAuthor)
	beego.Info("AppAuthorEmail =", AppAuthorEmail)

	beego.Info("PostDirectory =", PostDirectory)
	beego.Info("PostAuthor =", PostAuthor)
	beego.Info("PostAbstractionLen =", PostAbstractionLen)
}

func settingLocales() {
	// load locales with locale_LANG.ini files
	langs := "en-US|zh-CN"
	for _, lang := range strings.Split(langs, "|") {
		lang = strings.TrimSpace(lang)
		files := []string{"conf/" + "locale_" + lang + ".ini"}
		if fh, err := os.Open(files[0]); err == nil {
			fh.Close()
		} else {
			files = nil
		}
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini", files...); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			os.Exit(2)
		}
	}
	Langs = i18n.ListLangs()
}
