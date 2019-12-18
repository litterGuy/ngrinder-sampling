package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ngrinder-sampling/models"
	"ngrinder-sampling/utils"
	"strconv"
	"time"
)

type ScenesController struct {
	BaseController
}

// @router	/list	[get]
func (s *ScenesController) List() {
	page, _ := s.GetInt("page", 1)
	pageSize, _ := s.GetInt("limit", PAGESIZE)
	scenesName := s.GetString("scenesName")

	list, total := models.TestPmsGetPageByUserId(s.userId, scenesName, page, pageSize)
	_ = utils.NewPaginator(s.Ctx.Request, pageSize, total)

	s.result.Data = list
	s.result.Code = 0
	s.result.Count = int(total)
	s.responseAjax()
}

// @router	/getScenesById	[get]
func (s *ScenesController) GetScenesById() {
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

	s.result.Data = sencesRequestBean
	s.responseAjax()
}

// @router	/valid	[post]
func (s *ScenesController) Valid() {
	var sencesRequestBean models.SencesRequestBean
	data := s.Ctx.Input.RequestBody
	//json数据封装到user对象中
	err := json.Unmarshal(data, &sencesRequestBean)
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}

	//对参数完整性进行校验
	if err := models.ValidSencesParams(&sencesRequestBean); err != nil {
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}
	if len(sencesRequestBean.UserId) <= 0 {
		sencesRequestBean.UserId = s.userId
	}
	//随意设置定时时间
	sencesRequestBean.ScheduledTime = time.Now().Format("2006-01-02 15:04:05")

	ngrinderUrl := beego.AppConfig.String("ngrinder.serverurl")
	apiUrl := beego.AppConfig.String("ngrinder.api.validScript")
	req := httplib.Post(ngrinderUrl + apiUrl)
	req.Header("Content-Type", "application/json")
	req.JSONBody(sencesRequestBean)
	rst, err := req.String()
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
	} else {
		s.result.Data = rst
	}
	s.responseAjax()
}

// @router /create	[post]
func (s *ScenesController) Create() {
	var sencesRequestBean models.SencesRequestBean
	data := s.Ctx.Input.RequestBody
	//json数据封装到user对象中
	err := json.Unmarshal(data, &sencesRequestBean)
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}

	//对参数完整性进行校验
	if err := models.ValidSencesParams(&sencesRequestBean); err != nil {
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}
	if len(sencesRequestBean.UserId) <= 0 {
		sencesRequestBean.UserId = s.userId
	}

	//开启事务
	o := orm.NewOrm()
	//获取testPms
	testPms, err := models.GetTestPmsBean(sencesRequestBean)
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}
	testPms.CreateTime = time.Now()
	testPms.UpdateTime = testPms.CreateTime

	o.Begin()
	id, dbErr := models.TestPmsSave(testPms, o)
	if dbErr != nil {
		s.result.Code = 1
		s.result.ErrMsg = dbErr.Error()
		s.responseAjax()
	}
	testPms.Id = id

	//获取requestPms
	requestPmsArray, err := models.GetRequestPmsBean(sencesRequestBean)
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = dbErr.Error()
		s.responseAjax()
	}
	for i, requestPms := range *requestPmsArray {
		requestPms.TestPmsId = id
		//设置生成脚本时的函数名
		requestPms.FunName = "test" + strconv.Itoa(i);
		pid, dbError := models.RequestPmsSave(&requestPms, o)
		if dbError != nil {
			//事务回滚
			o.Rollback()
			s.result.Code = 1
			s.result.ErrMsg = dbError.Error()
			s.responseAjax()
		}
		requestPms.Id = pid
	}
	o.Commit()

	s.result.Data = id
	s.responseAjax()
}

// @router	/update	[post]
func (s *ScenesController) Update() {
	var sencesRequestBean models.SencesRequestBean
	data := s.Ctx.Input.RequestBody
	//json数据封装到user对象中
	err := json.Unmarshal(data, &sencesRequestBean)
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}

	//对参数完整性进行校验
	if err := models.ValidUpdateSencesParams(&sencesRequestBean); err != nil {
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}
	if len(sencesRequestBean.UserId) <= 0 {
		sencesRequestBean.UserId = s.userId
	}
	//确认id是否存在
	originTestPms, err := models.TestPmsGetById(sencesRequestBean.Id)
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = "the scenes of id " + strconv.FormatInt(sencesRequestBean.Id, 10) + " not exist"
		s.responseAjax()
	}
	//开启事务
	o := orm.NewOrm()
	//获取testPms
	testPms, err := models.GetTestPmsBean(sencesRequestBean)
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}
	testPms.UpdateTime = time.Now()
	testPms.CreateTime = originTestPms.CreateTime

	o.Begin()
	dbErr := models.TestPmsUpdate(testPms, o)
	if dbErr != nil {
		s.result.Code = 1
		s.result.ErrMsg = dbErr.Error()
		s.responseAjax()
	}

	//获取requestPms
	requestPmsArray, err := models.GetRequestPmsBean(sencesRequestBean)
	if err != nil {
		s.result.Code = 1
		s.result.ErrMsg = dbErr.Error()
		s.responseAjax()
	}
	//删除requestPms，重新添加
	_, err = models.RequestPmsDeleteByTestPmsId(testPms.Id, o)
	if err != nil {
		o.Rollback()
		s.result.Code = 1
		s.result.ErrMsg = err.Error()
		s.responseAjax()
	}
	for i, requestPms := range *requestPmsArray {
		//设置生成脚本时的函数名
		requestPms.FunName = "test" + strconv.Itoa(i);
		requestPms.Id = 0
		requestPms.TestPmsId = testPms.Id
		_, dbError := models.RequestPmsSave(&requestPms, o)
		if dbError != nil {
			//事务回滚
			o.Rollback()
			s.result.Code = 1
			s.result.ErrMsg = dbError.Error()
			s.responseAjax()
		}
	}
	o.Commit()

	s.responseAjax()
}

// @router	/delete	[get]
func (s *ScenesController) Delete() {
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
	//开启事务，删除requestPms、testPms
	o := orm.NewOrm()
	o.Begin()
	_, err = models.RequestPmsDeleteByTestPmsId(id, o)
	if err != nil {
		o.Rollback()
		s.result.ErrMsg = err.Error()
		s.result.Code = 1
		s.responseAjax()
	}
	_, err = models.TestPmsDelete(testPms, o)
	if err != nil {
		o.Rollback()
		s.result.ErrMsg = err.Error()
		s.result.Code = 1
		s.responseAjax()
	}
	o.Commit()

	//增加请求，删除svn上的脚本和压测数据源文件
	requestBean, err := models.BuildScenesBean(testPms, nil)
	if err != nil {
		s.result.ErrMsg = err.Error()
		s.result.Code = 1
		s.responseAjax()
	}
	var path string
	if len(requestBean.FileDataList) > 0 {
		for index, fileData := range requestBean.FileDataList {
			if index != 0 {
				path += ","
			}
			path += fileData.Path
		}
	}
	ngrinderUrl := beego.AppConfig.String("ngrinder.serverurl")
	apiUrl := beego.AppConfig.String("ngrinder.api.delete")
	ngrinderUrl += apiUrl + "?id=" + strconv.FormatInt(id, 10)
	if len(path) > 0 {
		ngrinderUrl += "&path=" + path
	}
	ngrinderUrl += "&userId=" + s.userId
	req := httplib.Get(ngrinderUrl)
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
			s.result.ErrMsg = js.ErrMsg
		}
	}
	if s.result.Code == 1 {
		logs.Error("delete script error:{}, ", s.result.ErrMsg)
		//脚本删除失败，不影响场景数据的删除
		s.result.Code = 0
	}

	s.responseAjax()
}
