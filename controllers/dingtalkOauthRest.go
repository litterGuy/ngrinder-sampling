package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"ngrinder-sampling/utils"
)

type DingTalkOauthRest struct {
	BaseController
}

type NUser struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

func (d *DingTalkOauthRest) DingTalk() {
	code := d.GetString("code")
	if len(code) <= 0 {
		d.Data["errMsg"] = "code is empty"
		d.TplName = "404.html"
		d.StopRun()
	}

	accessToken, err := utils.GetAccessToken()
	if err != nil {
		d.Data["errMsg"] = err.Error()
		d.TplName = "404.html"
		d.StopRun()
	}

	openid, persistentCode, err := utils.GetPersistentCode(accessToken, code)
	if err != nil {
		d.Data["errMsg"] = err.Error()
		d.TplName = "404.html"
		d.StopRun()
	}

	ngrinderUrl := beego.AppConfig.String("ngrinder.serverurl")
	apiUrl := beego.AppConfig.String("ngrinder.api.getUserByUserId")
	ngrinderUrl += apiUrl
	req := httplib.Get(ngrinderUrl + "?userId=" + *openid)

	type Rst struct {
		Code   int    `json:"code"`
		ErrMsg string `json:"errMsg"`
		Data   NUser  `json:"data"`
	}
	rst := new(Rst)
	err = req.ToJSON(rst)
	if err != nil {
		d.Data["errMsg"] = err.Error()
		d.TplName = "404.html"
		d.StopRun()
	}
	if rst.Code == 1 {
		d.Data["errMsg"] = errors.New(rst.ErrMsg)
		d.TplName = "404.html"
		d.StopRun()
	}
	if rst.Code == 0 {
		d.SetSession(SESSION_NAME, *openid)
		//重定向
		d.Redirect("/v1/home/index", 200)
		d.StopRun()
	}

	snsToken, err := utils.GetSnsToken(openid, persistentCode, accessToken)
	if err != nil {
		d.Data["errMsg"] = errors.New(rst.ErrMsg)
		d.TplName = "404.html"
		d.StopRun()
	}
	dingTalkUser, err := utils.GetUserInfo(snsToken)
	if err != nil {
		d.Data["errMsg"] = errors.New(rst.ErrMsg)
		d.TplName = "404.html"
		d.StopRun()
	}

	serverUrl := beego.AppConfig.String("ngrinder.serverurl")
	saveUrl := beego.AppConfig.String("ngrinder.api.dingTalkSave")
	serverUrl += saveUrl

	req = httplib.Put(serverUrl)
	req.Param("userId", dingTalkUser.Openid)
	req.Param("userName", dingTalkUser.Nick)

	rst = new(Rst)
	err = req.ToJSON(rst)
	if err != nil {
		d.Data["errMsg"] = err.Error()
		d.TplName = "404.html"
		d.StopRun()
	}
	if rst.Code != 0 {
		d.Data["errMsg"] = errors.New(rst.ErrMsg)
		d.TplName = "404.html"
		d.StopRun()
	}
	d.SetSession(SESSION_NAME, *openid)
	//重定向
	d.Redirect("/v1/home/index", 200)
}
