// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/linauror/owncms-go/controllers"
	"github.com/linauror/owncms-go/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/index", &controllers.IndexController{}, "get:Index")
	beego.Router("/post/:id/:slug", &controllers.PostController{}, "get:Info")
	beego.Router("/post/comment", &controllers.PostController{}, "post:Comment")
	beego.Router("/page/:slug", &controllers.PageController{}, "get:Info")
	beego.Router("/category/:slug", &controllers.CategoryController{}, "get:Index")
	beego.Router("/author/:username", &controllers.AuthorController{}, "get:Index")
	beego.Router("/search", &controllers.SearchController{}, "get:Index")
	beego.Router("/tag/:tag", &controllers.TagController{}, "get:Index")
	beego.AutoRouter(&controllers.UserController{})

	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/", &admincontrollers.IndexController{}, "get:Index"),
		beego.NSRouter("/login", &admincontrollers.IndexController{}, "get,post:Login"),
		beego.NSRouter("/siteconfig", &admincontrollers.SiteconfigController{}, "get:Index"),
		beego.NSRouter("/menu", &admincontrollers.MenuController{}, "get:Index"),
		beego.NSAutoRouter(&admincontrollers.CategoryController{}),
		beego.NSAutoRouter(&admincontrollers.PageController{}),
		beego.NSAutoRouter(&admincontrollers.PostController{}),
	)
	beego.AddNamespace(ns)
}
