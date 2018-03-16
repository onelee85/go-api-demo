// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"user/controllers"
	"user/filter"

	"github.com/astaxie/beego"
)

func init() {
	//为需要token验证的的request注册filter
	beego.InsertFilter("/v1/user/*", beego.BeforeRouter, filter.OAuthFilter)

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/thirdlogin",
			beego.NSInclude(
				&controllers.ThirdloginController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
