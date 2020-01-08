package controllers

import (
	"github.com/astaxie/beego"
	"net/url"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["goto"] = url.QueryEscape("https://oapi.dingtalk.com/connect/oauth2/sns_authorize?appid=ding9d9231bfd078cf0d&response_type=code&scope=snsapi_login&state=STATE&redirect_uri=http://127.0.0.1:8090/")
	c.TplName = "login.html"
}
