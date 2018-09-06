package main

import (
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

	beego.Run()
}
