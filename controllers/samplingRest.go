package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ngrinder-sampling/models"
)

type SamplingController struct {
	BaseController
}

// @router	/list	[get]
func (s *SamplingController) List() {
	page, _ := s.GetInt("page", 1)
	pageSize, _ := s.GetInt("limit", PAGESIZE)
	pftestId, err := s.GetInt64("pftestId")
	if err != nil {
		s.result.Code = 1
		s.result.Msg = "pftestId can not be empty"
		s.responseAjax()
	}
	list, total := models.SampResultGetPageByPftestId(pftestId, page, pageSize)
	s.result.Data = list
	s.result.Code = 0
	s.result.Count = int(total)
	s.responseAjax()
}

// @router	/getByPftestId	[get]
func (s *SamplingController) GetByPftestId() {
	pftestId, err := s.GetInt64("pftestId")
	if err != nil {
		s.result.ErrMsg = err.Error()
		s.result.Code = 1
		s.responseAjax()
	}

	sampList, err := models.SampResultGetByPftestId(pftestId)
	if err != nil {
		s.result.ErrMsg = err.Error()
		s.result.Code = 1
		s.responseAjax()
	}

	if len(*sampList) <= 0 {
		s.responseAjax()
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

	s.result.Data = sampList
	s.responseAjax()
}

// @router	/gather	[post]
func (s *SamplingController) Gather() {
	var sampReqBean []models.SampReqBean
	data := s.Ctx.Input.RequestBody
	//json数据封装到user对象中
	err := json.Unmarshal(data, &sampReqBean)
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}
	//校验数据完整性(因为该接口调用是由脚本，所以不进行回滚操作，数据不完整记录日志)
	sampRestultList, err := models.GetSampResultBean(&sampReqBean)
	if err != nil {
		logs.Error(err.Error())
	}
	if len(*sampRestultList) <= 0 {
		s.result.Code = 1
		s.result.ErrMsg = errors.New("there is no samp result").Error()
		s.responseAjax()
	}
	o := orm.NewOrm()
	num, err := models.SampResultSaveMulti(sampRestultList, o)
	if err != nil {
		logs.Error("sampling save error :{}", string(data))
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}
	if num < int64(len(*sampRestultList)) {
		logs.Error("sampling save error :{}", string(data))
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}
	s.responseAjax()
}

// @router	/delete	[get]
func (s *SamplingController) Delete() {
	pftestId, err := s.GetInt64("pftestId")
	if err != nil {
		s.result.ErrMsg = err.Error()
		s.result.Code = 1
		s.responseAjax()
	}
	o := orm.NewOrm()
	_, err = models.SampResultDeleteByTestPmsId(pftestId, o)
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}
	s.responseAjax()
}
