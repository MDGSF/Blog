package main

import (
	"github.com/MDGSF/Blog/modules/models"
	_ "github.com/MDGSF/Blog/routers"
	"github.com/MDGSF/Blog/setting"
	"github.com/beego/i18n"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogFuncCall(true)
	beego.SetLevel(beego.LevelDebug)
	beego.Debug("main start")

	initialize()

	beego.Run()
}

func initialize() {
	setting.LoadConfig()

	initStaticDirectory()

	beego.AddFuncMap("i18n", i18n.Tr)

	models.LoadAllPostsDirectory()

	models.DBStart()

	routers.Init()
}

func initStaticDirectory() {
	// use http://127.0.0.1:8080/static/css/blog.css to access "static" directory.
	// beego.SetStaticPath("/static", "static") // this is set default.

	// use http://localhost:8080/down1/123.txt to access directory "download1/123.txt"
	beego.SetStaticPath("/down1", "download1")

	beego.SetStaticPath("/images", "mdgsf.github.io/images")

	for _, postDir := range setting.PostDirectory {
		beego.SetStaticPath(postDir, postDir)
	}
}
