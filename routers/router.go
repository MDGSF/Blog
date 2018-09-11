package routers

import (
	"github.com/MDGSF/Blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/pg", &controllers.IndexController{})

	beego.Router("/default", &controllers.MainController{})
	beego.Router("/layui", &controllers.LayuiController{})
	beego.Router("/bootstrap", &controllers.BootstrapController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/marked", &controllers.MarkedController{})
	beego.Router("/posts/*", &controllers.PostController{})
	beego.Router("/tags", &controllers.TagsController{})
	beego.Router("/yearmontharchives", &controllers.YearMonthController{})
}
