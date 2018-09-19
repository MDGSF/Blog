package setting

import (
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

var (
	AppName string

	RunMode string

	AppVersion        string
	AppAuthor         string
	AppAuthorEmail    string
	AppAuthorGitHub   string
	AppAuthorTwitter  string
	AppAuthorFacebook string

	PostDirectory      []string
	PostAuthor         string
	PostAbstractionLen int
)

var (
	HTTPAddr string
	HTTPPort int
	AppHost  string

	EnableAdmin bool
	AdminAddr   string
	AdminPort   string
)

var (
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassWord string
	DBName     string
)

var (
	LoginRememberDays  int
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

	// you can use this to load many config file, app1.conf, app2.conf, app3.conf ...
	// beego.LoadAppConfig("ini", "conf/app.conf")

	// read app.conf
	AppName = beego.AppConfig.DefaultString("appname", "blog")
	RunMode = beego.AppConfig.DefaultString("runmode", "dev")

	AppVersion = beego.AppConfig.DefaultString("AppVersion", "1.0.0.1")
	AppAuthor = beego.AppConfig.DefaultString("AppAuthor", "author")
	AppAuthorEmail = beego.AppConfig.DefaultString("AppAuthorEmail", "xxx@email.com")
	AppAuthorGitHub = beego.AppConfig.DefaultString("AppAuthorGitHub", "")
	AppAuthorTwitter = beego.AppConfig.DefaultString("AppAuthorTwitter", "")
	AppAuthorFacebook = beego.AppConfig.DefaultString("AppAuthorFacebook", "")

	PostDirectory = beego.AppConfig.Strings("PostDirectory")
	PostAuthor = beego.AppConfig.DefaultString("PostAuthor", "author")
	PostAbstractionLen = beego.AppConfig.DefaultInt("PostAbstractionLen", 100)

	HTTPAddr = beego.AppConfig.DefaultString("httpaddr", "127.0.0.1")
	HTTPPort = beego.AppConfig.DefaultInt("httpport", 8080)
	AppHost = beego.AppConfig.DefaultString("app_host", "127.0.0.1:8080")
	EnableAdmin = beego.AppConfig.DefaultBool("EnableAdmin", false)
	AdminAddr = beego.AppConfig.DefaultString("AdminAddr", "localhost")
	AdminPort = beego.AppConfig.DefaultString("AdminPort", "8088")

	DBHost = beego.AppConfig.DefaultString("orm::DBHost", "localhost")
	DBPort = beego.AppConfig.DefaultString("orm::DBPort", "3306")
	DBUser = beego.AppConfig.DefaultString("orm::DBUser", "root")
	DBPassWord = beego.AppConfig.DefaultString("orm::DBPassWord", "123456")
	DBName = beego.AppConfig.DefaultString("orm::DBName", "db_name")

	LoginRememberDays = beego.AppConfig.DefaultInt("login_remember_days", 7)
	CookieRememberName = beego.AppConfig.DefaultString("cookie_remember_name", "blog_remember_name")
	CookieUserName = beego.AppConfig.DefaultString("cookie_user_name", "blog_user_name")

	// set to beego setting
	beego.BConfig.Listen.HTTPAddr = HTTPAddr
	beego.BConfig.Listen.HTTPPort = HTTPPort

	// print all settings
	beego.Info("AppName =", AppName)
	beego.Info("RunMode =", RunMode)

	beego.Info("AppVersion =", AppVersion)
	beego.Info("AppAuthor =", AppAuthor)
	beego.Info("AppAuthorEmail =", AppAuthorEmail)
	beego.Info("AppAuthorGitHub =", AppAuthorGitHub)
	beego.Info("AppAuthorTwitter =", AppAuthorTwitter)
	beego.Info("AppAuthorFacebook =", AppAuthorFacebook)

	beego.Info("PostDirectory =", PostDirectory)
	beego.Info("PostAuthor =", PostAuthor)
	beego.Info("PostAbstractionLen =", PostAbstractionLen)

	beego.Info("HTTPAddr =", HTTPAddr)
	beego.Info("HTTPPort =", HTTPPort)
	beego.Info("EnableAdmin =", EnableAdmin)
	beego.Info("AdminAddr =", AdminAddr)
	beego.Info("AdminPort =", AdminPort)

	beego.Info("DBHost =", DBHost)
	beego.Info("DBPort =", DBPort)
	beego.Info("DBUser =", DBUser)
	beego.Info("DBPassWord =", DBPassWord)
	beego.Info("DBName =", DBName)

	beego.Info("LoginRememberDays =", LoginRememberDays)
	beego.Info("CookieRememberName =", CookieRememberName)
	beego.Info("CookieUserName =", CookieUserName)
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
