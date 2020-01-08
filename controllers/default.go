package controllers

import (
	"github.com/astaxie/beego"
	"net/url"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	appId := beego.AppConfig.String("dingtalk.appid")
	redirectUrl := c.Ctx.Input.Site() + "/dingTalk"
	port := c.Ctx.Input.Port()
	if port != 80 && port != 443 {
		redirectUrl = c.Ctx.Input.Site() + ":" + strconv.Itoa(port) + "/dingTalk"
	}

	c.Data["goto"] = url.QueryEscape("https://oapi.dingtalk.com/connect/oauth2/sns_authorize?appid=" + appId + "&response_type=code&scope=snsapi_login&state=STATE&redirect_uri=" + redirectUrl)
	c.TplName = "login.html"
}
