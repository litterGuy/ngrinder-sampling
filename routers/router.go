package routers

import (
	"github.com/astaxie/beego"
	"ngrinder-sampling/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/samp",
			beego.NSInclude(&controllers.SamplingController{})),
		beego.NSNamespace("/scenes",
			beego.NSInclude(&controllers.ScenesController{})),
		beego.NSNamespace("/script",
			beego.NSInclude(&controllers.ScriptController{})),
		beego.NSNamespace("/user",
			beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/home",
			beego.NSInclude(&controllers.HomeController{})),
	)

	beego.AddNamespace(ns)
}
