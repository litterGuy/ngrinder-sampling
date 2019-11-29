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
	)

	beego.AddNamespace(ns)
}
