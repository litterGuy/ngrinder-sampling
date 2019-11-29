package models

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"strconv"
	"strings"
	"time"
)

type SampReqBean struct {
	PftestId string            `json:"pftestId" valid:"Required"`
	Sampling []ApiSamplingBean `json:"sampling"`
}

type ApiSamplingBean struct {
	ApiId          int64             `json:"api_id" valid:"Required"`
	Rt             int               `json:"rt"`
	Agent          string            `json:"agent"`
	HttpReqBody    string            `json:"http_req_body"`
	HttpResHeaders map[string]string `json:"http_res_headers"`
	Func           string            `json:"func"`
	HttpReqHeaders []NVPair          `json:"http_req_headers"`
	HttpResStatus  int               `json:"http_res_status"`
	HttpResBody    string            `json:"http_res_body"`
	HttpReqUrl     string            `json:"http_req_url"`
	HttpReqMethod  string            `json:"http_req_method"`
	TimeStamp      int64             `json:"timestamp"`
	ExportContent  map[string]string `json:"export_content"`
	CheckResult    []SencesAssertion `json:"check_result"`
}

func GetSampResultBean(s *[]SampReqBean) (*[]SampResult, error) {
	if len(*s) <= 0 {
		return nil, errors.New("the reqlist is null")
	}
	var result []SampResult
	for _, req := range *s {
		if len(req.Sampling) <= 0 {
			continue
		}
		//因为ngrinder生成的id格式为"test_140"
		if len(req.PftestId) <= 0 || !strings.Contains(req.PftestId, "_") {
			logs.Error("the sampling has no pftestId, {}", req)
			continue
		}

		pftestId, err := strconv.ParseInt(strings.Split(req.PftestId, "_")[1], 10, 64)
		if err!=nil{
			logs.Error("the sampling get the pftestId error, {}", req)
		}
		for _, api := range req.Sampling {
			str, _ := json.Marshal(api)
			if pftestId <= 0 || api.ApiId <= 0 {
				logs.Error("the sampling is not id, error: {}", str)
				continue
			}
			tmp := SampResult{}
			tmp.PftestId = pftestId
			tmp.ReqId = api.ApiId
			tmp.Func = api.Func
			tmp.Rt = api.Rt
			tmp.ReqStatus = api.HttpResStatus
			tmp.ReqContent = string(str)

			tmp.CreateTime = time.Unix(api.TimeStamp, 0)
			times := strconv.FormatInt(api.TimeStamp, 10)
			if len(times) == 13 {
				tmp.CreateTime = time.Unix(api.TimeStamp/1000, 0)
			}

			result = append(result, tmp)
		}
	}
	return &result, nil
}
