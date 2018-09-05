package main

import (
	_ "github.com/MDGSF/Blog/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetLevel(beego.LevelDebug)
	beego.Debug("main start")

	beego.Run()
}
