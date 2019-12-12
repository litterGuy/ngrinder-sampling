package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"ngrinder-sampling/utils"
)

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

func (b *BaseController) responseAjaxMore(more map[string]interface{}) {
	dst, _ := json.Marshal(b.result)
	src, _ := json.Marshal(more)
	result := utils.JSONMerger(dst, src)
	b.Data["json"] = result
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
