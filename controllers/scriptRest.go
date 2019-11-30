package controllers

import (
	"github.com/astaxie/beego"
	"ngrinder-sampling/models"
	"strconv"
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
	//请求ngrinder，生成压测脚本和测试数据

}

//handle the result
func (s *ScriptController) responseRst(result map[string]interface{}) {
	s.Data["json"] = result
	s.ServeJSON()
	s.StopRun()
}
