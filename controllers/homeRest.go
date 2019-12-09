package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"strconv"
)

//http://layuimini-onepage.99php.cn/#/page/icon.html

type HomeController struct {
	BaseController
}

// @router	/index	[get]
func (h *HomeController) Index() {
	h.Data["userId"] = h.GetSession(SESSION_NAME)
	h.TplName = "index.html"
}

// @router	/scenesList	[get]
func (h *HomeController) ScenesList() {
	h.TplName = "scenes_list.html"
}

// @router	/reportList	[get]
func (h *HomeController) ReportList() {
	ngrinderUrl := beego.AppConfig.String("ngrinder.serverurl")
	h.Data["ngrinder_host"] = ngrinderUrl
	h.TplName = "report_list.html"
}

// @router	/reportAjax	[get]
func (h *HomeController) ReportAjax() {
	page, _ := h.GetInt("page", 1)
	pageSize, _ := h.GetInt("limit", PAGESIZE)
	name := h.GetString("name")
	queryFilter := h.GetString("queryFilter")

	ngrinderUrl := beego.AppConfig.String("ngrinder.serverurl")
	apiUrl := beego.AppConfig.String("ngrinder.api.perfList")
	ngrinderUrl += apiUrl + "?page=" + strconv.Itoa(page-1) + "&pageSize=" + strconv.Itoa(pageSize)
	if len(name) > 0 {
		ngrinderUrl += "&name=" + name
	}
	if len(queryFilter) > 0 {
		ngrinderUrl += "&queryFilter=" + queryFilter
	}
	ngrinderUrl += "&userId=" + h.userId
	req := httplib.Get(ngrinderUrl)
	var js NsResponseBean
	rst, err := req.String()
	if err != nil {
		h.result.Code = 1
		h.result.Msg = err.Error()
	} else {
		err = json.Unmarshal([]byte(rst), &js)
		if err != nil {
			h.result.Code = 1
			h.result.Msg = rst
		} else {
			myMap := js.Data.(map[string]interface{})
			h.result.Count = int(myMap["total"].(float64))
			h.result.Data = myMap["list"]
		}
	}
	h.responseAjax()
}

// @router	/reportDelete	[get]
func (h *HomeController) ReportDelete() {
	ids := h.GetString("ids")

	ngrinderUrl := beego.AppConfig.String("ngrinder.serverurl")
	apiUrl := beego.AppConfig.String("ngrinder.api.perfDelete")
	ngrinderUrl += apiUrl
	ngrinderUrl += "?userId=" + h.userId
	if len(ids) > 0 {
		ngrinderUrl += "&ids=" + ids
	}
	req := httplib.Get(ngrinderUrl)
	var js NsResponseBean
	rst, err := req.String()
	if err != nil {
		h.result.Code = 1
		h.result.Msg = err.Error()
	} else {
		err = json.Unmarshal([]byte(rst), &js)
		if err != nil {
			h.result.Code = 1
			h.result.Msg = rst
		}
	}
	h.responseAjax()
}
