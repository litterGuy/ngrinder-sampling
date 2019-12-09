package controllers

import "github.com/astaxie/beego"

const (
	SESSION_NAME = "user_login_id"
	PAGESIZE     = 10
)

type BaseController struct {
	beego.Controller
	userId string
	result *NsResponseBean
}

func (b *BaseController) Prepare() {
	b.result = new(NsResponseBean)
	controller, action := b.GetControllerAndAction()

	if controller != "UserController" || action != "Login" {
		//校验是否登录
		userId := b.GetSession(SESSION_NAME)
		if userId != nil {
			b.userId = userId.(string)
		} else {
			b.Redirect("/", 302)
		}
	}
}

//handle the result
func (b *BaseController) responseAjax() {
	b.Data["json"] = b.result
	b.ServeJSON()
	b.StopRun()
}

type NsResponseBean struct {
	Code   int         `json:"code"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
	Count  int         `json:"count"`
	Msg    string      `json:"msg,omitempty"`
}
