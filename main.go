package main

import (
	_ "github.com/MDGSF/Blog/models"
	_ "github.com/MDGSF/Blog/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetLevel(beego.LevelDebug)
	beego.Debug("main start")

	appName := beego.AppConfig.String("appname")
	httpPort := beego.AppConfig.String("httpport")
	runMode := beego.AppConfig.String("runmode")
	beego.Info("appName =", appName)
	beego.Info("httpPort =", httpPort)
	beego.Info("runMode =", runMode)

	// use http://127.0.0.1:8080/static/css/blog.css to access "static" directory.

	// use http://localhost:8080/down1/123.txt to access directory "download1/123.txt"
	beego.SetStaticPath("/down1", "download1")

	beego.SetStaticPath("/images", "mdgsf.github.io/images")

	postDirs := beego.AppConfig.Strings("PostDirectory")
	beego.Info("postDirs =", postDirs)
	for _, postDir := range postDirs {
		beego.SetStaticPath(postDir, postDir)
	}

	beego.Run()
}
