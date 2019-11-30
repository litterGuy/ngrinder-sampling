package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ngrinder-sampling/models"
)

type SamplingController struct {
	beego.Controller
}

// @router	/getByPftestId	[get]
func (s *SamplingController) GetByPftestId() {
	result := make(map[string]interface{})
	result["code"] = 0
	pftestId, err := s.GetInt64("pftestId")
	if err != nil {
		result["errMsg"] = err
		result["code"] = 1
		s.responseRst(result)
	}

	sampList, err := models.SampResultGetByPftestId(pftestId)
	if err != nil {
		result["errMsg"] = err
		result["code"] = 1
		s.responseRst(result)
	}

	if len(*sampList) <= 0 {
		s.responseRst(result)
	}

	for i, samp := range *sampList {
		var apiSampBean models.ApiSamplingBean
		err := json.Unmarshal([]byte(samp.ReqContent), &apiSampBean)
		if err != nil {
			logs.Error("format samp result error {}", err.Error())
		}
		samp.ApiSamplingBean = apiSampBean
		(*sampList)[i] = samp
	}

	result["data"] = sampList
	s.responseRst(result)
}

// @router	/gather	[post]
func (s *SamplingController) Gather() {
	result := make(map[string]interface{})
	result["code"] = 0

	var sampReqBean []models.SampReqBean
	data := s.Ctx.Input.RequestBody
	//json数据封装到user对象中
	err := json.Unmarshal(data, &sampReqBean)
	if err != nil {
		result["code"] = 1
		result["errMsg"] = err.Error()
		s.responseRst(result)
	}
	//校验数据完整性(因为该接口调用是由脚本，所以不进行回滚操作，数据不完整记录日志)
	sampRestultList, err := models.GetSampResultBean(&sampReqBean)
	if err != nil {
		logs.Error(err.Error())
	}
	if len(*sampRestultList) <= 0 {
		result["code"] = 1
		result["errMsg"] = errors.New("there is no samp result")
		s.responseRst(result)
	}
	o := orm.NewOrm()
	num, err := models.SampResultSaveMulti(sampRestultList, o)
	if err != nil {
		logs.Error("sampling save error :{}", string(data))
		result["code"] = 1
		result["errMsg"] = err.Error()
		s.responseRst(result)
	}
	if num < int64(len(*sampRestultList)) {
		logs.Error("sampling save error :{}", string(data))
		result["code"] = 1
		result["errMsg"] = err.Error()
		s.responseRst(result)
	}
	s.responseRst(result)
}

// @router	/delete	[get]
func (s *SamplingController) Delete() {
	result := make(map[string]interface{})
	result["code"] = 0
	pftestId, err := s.GetInt64("pftestId")
	if err != nil {
		result["errMsg"] = err
		result["code"] = 1
		s.responseRst(result)
	}
	o := orm.NewOrm()
	_, err = models.SampResultDeleteByTestPmsId(pftestId, o)
	if err != nil {
		result["code"] = 1
		result["errMsg"] = err.Error()
		s.responseRst(result)
	}
	s.responseRst(result)
}

//handle the result
func (s *SamplingController) responseRst(result map[string]interface{}) {
	s.Data["json"] = result
	s.ServeJSON()
	s.StopRun()
}
