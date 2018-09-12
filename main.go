package main

import (
	"github.com/MDGSF/Blog/models"
	_ "github.com/MDGSF/Blog/routers"
	"github.com/MDGSF/Blog/setting"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetLevel(beego.LevelDebug)
	beego.Debug("main start")

	initialize()

	beego.Run()
}

func initialize() {
	setting.LoadConfig()

	initStaticDirectory()

	models.LoadAllPostsDirectory()
}

func initStaticDirectory() {
	// use http://127.0.0.1:8080/static/css/blog.css to access "static" directory.

	// use http://localhost:8080/down1/123.txt to access directory "download1/123.txt"
	beego.SetStaticPath("/down1", "download1")

	beego.SetStaticPath("/images", "mdgsf.github.io/images")

	for _, postDir := range setting.PostDirectory {
		beego.SetStaticPath(postDir, postDir)
	}
}
