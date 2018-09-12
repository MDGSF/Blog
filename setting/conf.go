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

	PostDirectory      []string
	PostAuthor         string
	PostAbstractionLen int
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

	PostDirectory = beego.AppConfig.Strings("PostDirectory")
	PostAuthor = beego.AppConfig.DefaultString("PostAuthor", "author")
	PostAbstractionLen = beego.AppConfig.DefaultInt("PostAbstractionLen", 100)

	beego.Info("AppName =", AppName)
	beego.Info("HTTPPort =", HTTPPort)
	beego.Info("RunMode =", RunMode)

	beego.Info("EnableAdmin =", EnableAdmin)
	beego.Info("AdminAddr =", AdminAddr)
	beego.Info("AdminPort =", AdminPort)

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
