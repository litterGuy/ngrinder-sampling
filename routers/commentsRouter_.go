package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ngrinder-sampling/controllers:HomeController"] = append(beego.GlobalControllerRouter["ngrinder-sampling/controllers:HomeController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ngrinder-sampling/controllers:HomeController"] = append(beego.GlobalControllerRouter["ngrinder-sampling/controllers:HomeController"],
        beego.ControllerComments{
            Method: "ScenesList",
            Router: `/scenesList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ngrinder-sampling/controllers:SamplingController"] = append(beego.GlobalControllerRouter["ngrinder-sampling/controllers:SamplingController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ngrinder-sampling/controllers:SamplingController"] = append(beego.GlobalControllerRouter["ngrinder-sampling/controllers:SamplingController"],
        beego.ControllerComments{
            Method: "Gather",
            Router: `/gather`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ngrinder-sampling/controllers:SamplingController"] = append(beego.GlobalControllerRouter["ngrinder-sampling/controllers:SamplingController"],
        beego.ControllerComments{
            Method: "GetByPftestId",
            Router: `/getByPftestId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ngrinder-sampling/controllers:ScenesController"] = append(beego.GlobalControllerRouter["ngrinder-sampling/controllers:ScenesController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ngrinder-sampling/controllers:ScenesController"] = append(beego.GlobalControllerRouter["ngrinder-sampling/controllers:ScenesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ngrinder-sampling/controllers:ScenesController"] = append(beego.GlobalControllerRouter["ngrinder-sampling/controllers:ScenesController"],
        beego.ControllerComments{
            Method: "GetScenesById",
            Router: `/getScenesById`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ngrinder-sampling/controllers:ScenesController"] = append(beego.GlobalControllerRouter["ngrinder-sampling/controllers:ScenesController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ngrinder-sampling/controllers:ScenesController"] = append(beego.GlobalControllerRouter["ngrinder-sampling/controllers:ScenesController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ngrinder-sampling/controllers:ScriptController"] = append(beego.GlobalControllerRouter["ngrinder-sampling/controllers:ScriptController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ngrinder-sampling/controllers:UserController"] = append(beego.GlobalControllerRouter["ngrinder-sampling/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ngrinder-sampling/controllers:UserController"] = append(beego.GlobalControllerRouter["ngrinder-sampling/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
