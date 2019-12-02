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
	beego.Controller
}

/**
调用ngrinder去生成pftest
*/
// @router	/create	[post]
func (s *ScriptController) Create() {
	result := make(map[string]interface{})
	result["code"] = 0
	id, err := s.GetInt64("id")
	if err != nil {
		result["errMsg"] = err
		result["code"] = 1
		s.responseRst(result)
	}

	//获取testPms
	testPms, err := models.TestPmsGetById(id)
	if err != nil {
		result["code"] = 1
		result["errMsg"] = "the scenes of id " + strconv.FormatInt(id, 10) + " not exist"
		s.responseRst(result)
	}
	//获取requestPms
	requestPmsList, err := models.RequestPmsGetByTestPmsId(id)
	if err != nil {
		result["code"] = 1
		result["errMsg"] = err
		s.responseRst(result)
	}
	//组装数据
	sencesRequestBean, err := models.BuildScenesBean(testPms, &requestPmsList)
	if err != nil {
		result["code"] = 1
		result["errMsg"] = err
		s.responseRst(result)
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
	}
	//请求ngrinder，生成压测脚本和测试数据
	ngrinderUrl := beego.AppConfig.String("ngrinder.serverurl")
	req := httplib.Post(ngrinderUrl)
	req.Header("Content-Type", "application/json")
	req.JSONBody(sencesRequestBean)
	var js struct{ Code int }
	rst, err := req.String()
	if err != nil {
		result["code"] = 1
		result["errMsg"] = err
	} else {
		err = json.Unmarshal([]byte(rst), &js)
		if err != nil {
			result["code"] = 1
			result["errMsg"] = rst
		} else {
			result["code"] = js.Code
		}
	}

	s.responseRst(result)
}

//handle the result
func (s *ScriptController) responseRst(result map[string]interface{}) {
	s.Data["json"] = result
	s.ServeJSON()
	s.StopRun()
}
