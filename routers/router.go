package routers

import (
	"github.com/MDGSF/Blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/layui", &controllers.LayuiController{})
	beego.Router("/bootstrap", &controllers.BootstrapController{})
}
