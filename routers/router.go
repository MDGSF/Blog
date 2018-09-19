package routers

import (
	"github.com/MDGSF/Blog/controllers/admin"
	"github.com/MDGSF/Blog/controllers/api"
	"github.com/MDGSF/Blog/controllers/demo"

	"github.com/astaxie/beego"
)

// Init init routers
func Init() {

	initDemo()

	initFront()

	initAdmin()
}

func initAdmin() {
	beego.Router("/admin", &admin.AdminController{})

	adminLogin := &admin.LoginController{}
	beego.Router("/admin/login", adminLogin, "get:Get;post:Login")
	beego.Router("/admin/logout", adminLogin, "get:Logout")

	beego.Router("/admin/register", &admin.RegisterController{})
	beego.Router("/admin/resetpwd", &admin.ResetPasswordController{})
	beego.Router("/admin/forgotpwd", &admin.ForgotPasswordController{})
}

func initFront() {
	beego.Router("/", &api.IndexController{})
	beego.Router("/pg", &api.IndexController{})
	beego.Router("/posts/*", &api.PostController{})
	beego.Router("/tags", &api.TagsController{})
	beego.Router("/yearmontharchives", &api.YearMonthController{})
	beego.Router("/search", &api.SearchController{})
	beego.Router("/about", &api.AboutController{})
	beego.Router("/classification", &api.ClassificationController{})
}

func initDemo() {
	beego.Router("/default", &demo.MainController{})
	beego.Router("/layui", &demo.LayuiController{})
	beego.Router("/bootstrap", &demo.BootstrapController{})
	beego.Router("/login", &demo.LoginController{})
	beego.Router("/upload", &demo.UploadController{})
	beego.Router("/marked", &demo.MarkedController{})
}
