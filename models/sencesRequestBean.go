package models

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/validation"
	"github.com/gogf/gf/util/gconv"
	"reflect"
	"time"
)

/**
用于接收创建场景脚本参数的结构体，
不与数据库对应
*/
type SencesRequestBean struct {
	Id                      int64              `json:"id,omitempty"`
	Name                    string             `json:"name,omitempty" valid:"Required;MaxSize(200)"`
	FileDataList            []SencesFileData   `json:"fileDataList,omitempty"` //转化成string存库 fileData
	IgnoreSampleCount       int                `json:"ignoreSampleCount,omitempty"`
	TargetHosts             string             `json:"targetHosts,omitempty"`
	UseRampUp               string             `json:"useRampUp,omitempty" valid:"Required"`
	RampUpType              string             `json:"rampUpType,omitempty"`
	Threshold               string             `json:"threshold,omitempty" valid:"Required"`
	Duration                int64              `json:"duration,omitempty"`
	RunCount                int                `json:"runCount,omitempty"`
	AgentCount              int                `json:"agentCount,omitempty" valid:"Required"`
	VuserPerAgent           int                `json:"vuserPerAgent,omitempty" valid:"Required"`
	Processes               int                `json:"processes,omitempty" valid:"Required"`
	RampUpInitCount         int                `json:"rampUpInitCount,omitempty"`
	RampUpInitSleepTime     int                `json:"rampUpInitSleepTime,omitempty"`
	RampUpStep              int                `json:"rampUpStep,omitempty"`
	RampUpIncrementInterval int                `json:"rampUpIncrementInterval,omitempty"`
	Threads                 int                `json:"threads,omitempty" valid:"Required"`
	SamplingInterval        int                `json:"samplingInterval,omitempty" valid:"Required"`
	Param                   string             `json:"param,omitempty"`
	CreateTime              time.Time          `json:"createTime,omitempty"`
	UpdateTime              time.Time          `json:"updateTime,omitempty"`
	RequestPmsList          []APIRequestParams `json:"requestPmsList,omitempty" valid:"Required"`
	ScheduledTime           string             `json:"scheduledTime"`
}

type APIRequestParams struct {
	Id            int64             `json:"id,omitempty"`
	TestPmsId     int64             `json:"testPmsId,omitempty"`
	ApiName       string            `json:"apiName,omitempty" valid:"Required;MaxSize(100)"`
	Sort          int               `json:"sort" valid:"Required;"`
	FunName       string            `json:"funName,omitempty"`
	Type          int               `json:"type" valid:"Required;"`
	Method        string            `json:"method,omitempty" valid:"Required"`
	Timeout       int64             `json:"timeout,omitempty"`
	Url           string            `json:"url,omitempty" valid:"Required"`
	HeaderList    []NVPair          `json:"headerList,omitempty"` //转化成string存库 headers
	ContentType   string            `json:"contentType,omitempty"`
	Body          string            `json:"body,omitempty"`
	ParamList     []NVPair          `json:"paramList,omitempty"`     //转化成string存库 params
	OutParamsList []SencesOutParams `json:"outParamsList,omitempty"` //转化成string存库 outParams
	AssertionList []SencesAssertion `json:"assertionList,omitempty"` //转化成string存库 assertion
}

//数据文件结构体
type SencesFileData struct {
	Name       string   `json:"name,omitempty"`
	Path       string   `json:"path,omitempty"`
	HasHead    int      `json:"hasHead,omitempty"`
	ParamsList []NVPair `json:"paramsList,omitempty"`
}

type NVPair struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

//出参设置
type SencesOutParams struct {
	Name           string `json:"name,omitempty"`
	Source         int    `json:"source"`
	ResolveExpress string `json:"resolveExpress,omitempty"`
	Index          int    `json:"index"`
}

//检查点设置
type SencesAssertion struct {
	Name    string `json:"name,omitempty"`
	Type    int    `json:"type"`
	Factor  string `json:"factor,omitempty"`
	Content string `json:"content,omitempty"`
}

func GetTestPmsBean(s SencesRequestBean) (*TestPms, error) {
	var testPms TestPms
	if err := gconv.Struct(s, &testPms); err != nil {
		return nil, err
	}
	//fileDataList转化成字符串
	if len(s.FileDataList) > 0 {
		dataFile, dfErr := json.Marshal(s.FileDataList)
		if dfErr != nil {
			return nil, dfErr
		}
		testPms.FileData = string(dataFile)
	}
	return &testPms, nil
}

func GetRequestPmsBean(a SencesRequestBean) (*[]RequestPms, error) {
	result := make([]RequestPms, len(a.RequestPmsList))
	for i, s := range a.RequestPmsList {
		var requestPms RequestPms
		if err := gconv.Struct(s, &requestPms); err != nil {
			return nil, err
		}
		//headerList转化成headers
		if len(s.HeaderList) > 0 {
			headers, headErr := json.Marshal(s.HeaderList)
			if headErr != nil {
				return nil, headErr
			}
			requestPms.Headers = string(headers)
		}
		//paramList转化成params
		if len(s.ParamList) > 0 {
			params, pmsErr := json.Marshal(s.ParamList)
			if pmsErr != nil {
				return nil, pmsErr
			}
			requestPms.Params = string(params)
		}
		//outParamList转化成outParams
		if len(s.OutParamsList) > 0 {
			outParams, outPmsErr := json.Marshal(s.OutParamsList)
			if outPmsErr != nil {
				return nil, outPmsErr
			}
			requestPms.OutParams = string(outParams)
		}
		//assertionList转化成assertion
		if len(s.AssertionList) > 0 {
			assertions, asErr := json.Marshal(s.AssertionList)
			if asErr != nil {
				return nil, asErr
			}
			requestPms.Assertion = string(assertions)
		}
		result[i] = requestPms
	}

	return &result, nil
}

/**
校验请求参数是否完整
*/
func ValidSencesParams(s *SencesRequestBean) error {
	valid := validation.Validation{}
	b, err := valid.Valid(s)
	if err != nil {
		return err
	}
	if !b {
		// validation does not pass
		return errors.New(valid.Errors[0].Field + " " + valid.Errors[0].Error())
	}
	//检验嵌套结构体中的参数
	//fileDataList
	if len(s.FileDataList) > 0 {
		for _, fileData := range s.FileDataList {
			if len(fileData.ParamsList) <= 0 {
				return errors.New("fileData paramList can not be empty")
			}
		}
	}

	for _, requestPms := range s.RequestPmsList {
		//outParamList
		if len(requestPms.OutParamsList) > 0 {
			for _, outParam := range requestPms.OutParamsList {
				if len(outParam.Name) <= 0 {
					return errors.New(requestPms.ApiName + ":outparam`s name can not be empty")
				}
				if len(outParam.ResolveExpress) <= 0 {
					return errors.New(requestPms.ApiName + ":outparam`s resolveExpress can not be empty")
				}
			}
		}
		//assertionList
		if len(requestPms.AssertionList) > 0 {
			for _, assertion := range requestPms.AssertionList {
				if len(assertion.Name) <= 0 {
					return errors.New(requestPms.ApiName + ":assertion`s name can not be empty")
				}
				if len(assertion.Factor) <= 0 {
					return errors.New(requestPms.ApiName + ":assertion`s factor can not be empty")
				}
				if len(assertion.Content) <= 0 {
					return errors.New(requestPms.ApiName + ":assertion`s content can not be empty")
				}
			}
		}
	}
	return nil
}

func ValidUpdateSencesParams(s *SencesRequestBean) error {
	err := ValidSencesParams(s)
	if err != nil {
		return err
	}
	if s.Id <= 0 {
		return errors.New("testPms`s id can not be empty")
	}
	for _, requestPms := range s.RequestPmsList {
		if requestPms.Id <= 0 {
			return errors.New("requestPms`s id can not be empty")
		}
		if requestPms.TestPmsId <= 0 {
			return errors.New("requestPms`s testPmsId can not be empty")
		}
	}
	return nil
}

//组装数据
func BuildScenesBean(pms *TestPms, requestPmsList *[]RequestPms) (*SencesRequestBean, error) {
	var sencesRequestBean SencesRequestBean
	if err := gconv.Struct(pms, &sencesRequestBean); err != nil {
		return nil, err
	}
	//转成fileDataList
	var fileDataList []SencesFileData
	err := json.Unmarshal([]byte(pms.FileData), &fileDataList)
	if err != nil {
		return nil, err
	}
	sencesRequestBean.FileDataList = fileDataList

	if IsNil(requestPmsList) || len(*requestPmsList) <= 0 {
		return &sencesRequestBean, nil
	}
	//转化requestPmsList
	apiRequestParamsList := make([]APIRequestParams, len(*requestPmsList))
	for i, requestPms := range *requestPmsList {
		var apiRequestParams APIRequestParams
		if err := gconv.Struct(requestPms, &apiRequestParams); err != nil {
			return nil, err
		}
		//headers转化
		if len(requestPms.Headers) > 0 {
			var headerList []NVPair
			err := json.Unmarshal([]byte(requestPms.Headers), &headerList)
			if err != nil {
				return nil, err
			}
			apiRequestParams.HeaderList = headerList
		}
		//params转化
		if len(requestPms.Params) > 0 {
			var paramsList []NVPair
			err := json.Unmarshal([]byte(requestPms.Params), &paramsList)
			if err != nil {
				return nil, err
			}
			apiRequestParams.ParamList = paramsList
		}
		//outParams转化
		if len(requestPms.OutParams) > 0 {
			var outParamsList []SencesOutParams
			err := json.Unmarshal([]byte(requestPms.OutParams), &outParamsList)
			if err != nil {
				return nil, err
			}
			apiRequestParams.OutParamsList = outParamsList
		}
		//assertion转化
		if len(requestPms.Assertion) > 0 {
			var assertionList []SencesAssertion
			err := json.Unmarshal([]byte(requestPms.Assertion), &assertionList)
			if err != nil {
				return nil, err
			}
			apiRequestParams.AssertionList = assertionList
		}
		apiRequestParamsList[i] = apiRequestParams
	}

	sencesRequestBean.RequestPmsList = apiRequestParamsList
	return &sencesRequestBean, nil
}

func IsNil(i interface{}) bool {
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(i)
	return vi.IsNil()
}
