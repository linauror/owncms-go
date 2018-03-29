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

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/index", &controllers.IndexController{}, "get:Index")
	beego.Router("/post/:id/:slug", &controllers.PostController{}, "get:Info")
	beego.Router("/page/:slug", &controllers.PageController{}, "get:Info")
	beego.Router("/category/:slug", &controllers.CategoryController{}, "get:Index")
	beego.Router("/author/:username", &controllers.AuthorController{}, "get:Index")
	beego.Router("/search", &controllers.SearchController{}, "get:Index")
	beego.Router("/tag/:tag", &controllers.TagController{}, "get:Index")
	beego.AutoRouter(&controllers.UserController{})
}
