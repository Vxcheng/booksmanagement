// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"booksmanagementSys/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//add by xjy
	beego.Router("/index", &controllers.UserController{})
	beego.Router("/register", &controllers.UserController{}, "post:Register")
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	//add by xjy end

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/books",
			beego.NSInclude(
				&controllers.BooksController{},
			),
		),
		//add by Chengshipeng
		//注册登录日志
		beego.NSNamespace("/user",
			beego.NSNamespace("/registor",
				beego.NSInclude(
					&controllers.RegistorController{},
				),
			),
			beego.NSNamespace("/login",
				beego.NSInclude(
					&controllers.LoginController{},
				),
			),

			beego.NSNamespace("/logout",
				beego.NSInclude(
					&controllers.LogoutController{},
				),
			),
		),

		beego.NSNamespace("/logs",
			beego.NSInclude(
				&controllers.LogsController{},
			),
		),
		//add by Chengshipeng end
		//thanks
	)
	beego.AddNamespace(ns)
}
