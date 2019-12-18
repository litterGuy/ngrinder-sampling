package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"ngrinder-sampling/models"
	"strconv"
	"time"
)

type ScriptController struct {
	BaseController
}

/**
调用ngrinder去生成pftest
*/
// @router	/create	[post]
func (s *ScriptController) Create() {
	id, err := s.GetInt64("id")
	if err != nil {
		s.result.ErrMsg = err.Error()
		s.result.Code = 1
		s.responseAjax()
	}

	//获取testPms
	testPms, err := models.TestPmsGetById(id)
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = "the scenes of id " + strconv.FormatInt(id, 10) + " not exist"
		s.responseAjax()
	}
	//获取requestPms
	requestPmsList, err := models.RequestPmsGetByTestPmsId(id)
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}
	//组装数据
	sencesRequestBean, err := models.BuildScenesBean(testPms, &requestPmsList)
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}
	scheduledTimeStr := s.GetString("scheduledTime")
	if len(scheduledTimeStr) > 0 {
		loc, _ := time.LoadLocation("Local")
		scheduledTime, err := time.ParseInLocation("2006-01-02 15:04:05", scheduledTimeStr, loc)
		if err != nil {
			logs.Error("script convert time error {}", err.Error())
		} else {
			sencesRequestBean.ScheduledTime = scheduledTime.Format("2006-01-02 15:04:05")
		}
	} else {
		sencesRequestBean.ScheduledTime = time.Now().Format("2006-01-02 15:04:05")
	}
	//请求ngrinder，生成压测脚本和测试数据
	ngrinderUrl := beego.AppConfig.String("ngrinder.serverurl")
	apiUrl := beego.AppConfig.String("ngrinder.api.create")
	req := httplib.Post(ngrinderUrl + apiUrl)
	req.Header("Content-Type", "application/json")
	req.JSONBody(sencesRequestBean)
	var js NsResponseBean
	rst, err := req.String()
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
	} else {
		err = json.Unmarshal([]byte(rst), &js)
		if err != nil {
			s.result.Code = 1
			s.result.ErrMsg = rst
		} else {
			s.result.Code = js.Code
		}
	}

	s.responseAjax()
}
