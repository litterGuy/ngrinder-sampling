package controllers

import (
	"bufio"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"html/template"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
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

// @router	/samplingLog	[get]
func (h *HomeController) SamplingLog() {
	pftestId, err := h.GetInt64("pftestId")
	if err != nil {
		//跳转向错误页
	}
	h.Data["pftestId"] = pftestId
	h.TplName = "sampling_log.html"
}

// @router	/scenesCreate	[get]
func (h *HomeController) ScenesCreate() {
	config := getAgentConfig(h.userId)
	h.Data["agentConfig"] = config
	h.TplName = "scenes_create.html"
}

// @router	/preview	[get]
func (h *HomeController) Preview() {
	h.TplName = "preview.html"
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

// @router	/agentList	[get]
func (h *HomeController) AgentList() {
	page, _ := h.GetInt("page", 1)
	pageSize, _ := h.GetInt("limit", PAGESIZE)
	ip := h.GetString("ip")

	ngrinderUrl := beego.AppConfig.String("ngrinder.serverurl")
	apiUrl := beego.AppConfig.String("ngrinder.api.agentList")
	ngrinderUrl += apiUrl + "?page=" + strconv.Itoa(page-1) + "&limit=" + strconv.Itoa(pageSize)
	if len(ip) > 0 {
		ngrinderUrl += "&ip=" + ip
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
		} else {
			myMap := js.Data.(map[string]interface{})
			h.result.Count = int(myMap["total"].(float64))
			h.result.Data = myMap["list"]
		}
	}
	h.responseAjax()
}

// @router	/uploadDataFile	[post]
func (h *HomeController) UploadDataFile() {
	file, header, err := h.GetFile("uploadFile")
	oldPath := h.GetString("oldPath", "")
	if err != nil {
		h.result.Code = 1
		h.result.ErrMsg = err.Error()
		h.responseAjax()
	}
	ext := path.Ext(header.Filename)
	if !strings.Contains(ext, "csv") {
		h.result.Code = 1
		h.result.ErrMsg = "file type is not csv"
		h.responseAjax()
	}
	//解析文件，确认有几个参数
	lineNum := 0
	br := bufio.NewReader(file)
	for {
		line, _, c := br.ReadLine()
		if c != nil {
			break
		}
		str := string(line)
		lineNum = strings.Count(str, ",")
		break
	}
	h.result.Count = lineNum + 1
	//保存文件
	uploadDir := "static/upload/" + time.Now().Format("20060102")
	_ = os.MkdirAll(uploadDir, 777)
	fpath := uploadDir + header.Filename
	defer file.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = h.SaveToFile("uploadFile", fpath)
	defer os.Remove(fpath)
	if err != nil {
		h.result.Code = 1
		h.result.ErrMsg = err.Error()
		h.responseAjax()
	}
	//将资源上传到ngrinder服务
	ngrinderUrl := beego.AppConfig.String("ngrinder.serverurl")
	apiUrl := beego.AppConfig.String("ngrinder.api.uploadData")
	ngrinderUrl += apiUrl

	req := httplib.Post(ngrinderUrl)
	req.Param("userId", h.userId)
	req.Param("oldPath", oldPath)
	req.PostFile("uploadFile", fpath)
	var js NsResponseBean
	rst, err := req.String()
	if err != nil {
		h.result.Code = 1
		h.result.ErrMsg = err.Error()
		h.responseAjax()
	} else {
		err = json.Unmarshal([]byte(rst), &js)
		if err != nil {
			h.result.Code = 1
			h.result.ErrMsg = err.Error()
			h.responseAjax()
		} else {
			more := make(map[string]interface{})
			more["fileName"] = header.Filename
			h.result.Code = js.Code
			h.result.ErrMsg = js.ErrMsg
			h.result.Data = js.Data
			h.responseAjaxMore(more)
		}
	}
}

func getAgentConfig(userId string) map[string]interface{} {
	ngrinderUrl := beego.AppConfig.String("ngrinder.serverurl")
	apiUrl := beego.AppConfig.String("ngrinder.api.agentConfig")
	ngrinderUrl += apiUrl + "?userId=" + userId
	req := httplib.Get(ngrinderUrl)
	var js NsResponseBean
	rst, err := req.String()
	if err != nil {
		return nil
	} else {
		err = json.Unmarshal([]byte(rst), &js)
		if err != nil {
			return nil
		} else {
			myMap := js.Data.(map[string]interface{})
			//将js模板化
			myMap["vuserCalcScript"] = template.JS(myMap["vuserCalcScript"].(string))
			return myMap
		}
	}
}
