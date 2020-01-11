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
	//如果不配置dingtalk，则隐藏页面相关信息
	if len(appId) > 0 {
		redirectUrl := c.Ctx.Input.Site() + "/dingTalk"
		port := c.Ctx.Input.Port()
		if port != 80 && port != 443 {
			redirectUrl = c.Ctx.Input.Site() + ":" + strconv.Itoa(port) + "/dingTalk"
		}
		gotoUrl := "https://oapi.dingtalk.com/connect/oauth2/sns_authorize?appid=" + appId + "&response_type=code&scope=snsapi_login&state=STATE&redirect_uri=" + redirectUrl
		c.Data["goto"] = url.QueryEscape(gotoUrl)
		c.Data["goto_original"] = gotoUrl
	}

	c.TplName = "login.html"
}
