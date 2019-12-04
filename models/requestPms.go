package models

import "github.com/astaxie/beego/orm"

//请求api
type RequestPms struct {
	Id           int64
	TestPmsId    int64
	ApiName      string
	Sort         int
	FunName      string
	Type         int
	Method       string
	Timeout      int64
	Url          string
	Headers      string
	ContentType  string
	Body         string
	Params       string
	OutParams    string
	Assertion    string
	WaitTime     int
	WaitVuserNum int
}

func init() {
	orm.RegisterModel(new(RequestPms))
}

func tableName() string {
	return "request_pms"
}

func RequestPmsSave(t *RequestPms, o orm.Ormer) (int64, error) {
	return o.Insert(t)
}

func RequestPmsUpdate(t *RequestPms, o orm.Ormer) error {
	if _, err := o.Update(t); err != nil {
		return err
	}
	return nil
}

func RequestPmsDelete(t *RequestPms, o orm.Ormer) (int64, error) {
	return o.Delete(t)
}

func RequestPmsDeleteByTestPmsId(testPmsId int64, o orm.Ormer) (int64, error) {
	return o.QueryTable(tableName()).Filter("test_pms_id", testPmsId).Delete()
}

func RequestPmsGetByTestPmsId(testPmsId int64) ([]RequestPms, error) {
	o := orm.NewOrm()
	var requestPmsList []RequestPms
	_, err := o.QueryTable(tableName()).Filter("test_pms_id", testPmsId).OrderBy("sort").All(&requestPmsList)
	if err != nil {
		return nil, err
	}
	return requestPmsList, nil
}
