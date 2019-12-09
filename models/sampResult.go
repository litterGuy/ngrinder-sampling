package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//采样结果
type SampResult struct {
	Id              int64
	PftestId        int64
	ReqId           int64
	Rt              int
	CreateTime      time.Time
	ReqContent      string
	Func            string
	ReqStatus       int
	ApiSamplingBean ApiSamplingBean `orm:"-"`
}

func init() {
	orm.RegisterModel(new(SampResult))
}

func sampResultTableName() string {
	return "samp_result"
}

func SampResultSave(t *SampResult, o orm.Ormer) (int64, error) {
	return o.Insert(t)
}

func SampResultSaveMulti(t *[]SampResult, o orm.Ormer) (int64, error) {
	return o.InsertMulti(100, *t)
}

func SampResultUpdate(t *SampResult, o orm.Ormer) error {
	if _, err := o.Update(t); err != nil {
		return err
	}
	return nil
}

func SampResultDelete(t *SampResult, o orm.Ormer) (int64, error) {
	return o.Delete(t)
}

func SampResultDeleteByTestPmsId(testPmsId int64, o orm.Ormer) (int64, error) {
	return o.QueryTable(sampResultTableName()).Filter("pftest_id", testPmsId).Delete()
}

func SampResultGetByPftestId(pftestId int64) (*[]SampResult, error) {
	o := orm.NewOrm()
	var sampResultList []SampResult
	_, err := o.QueryTable(sampResultTableName()).Filter("pftest_id", pftestId).All(&sampResultList)
	if err != nil {
		return nil, err
	}
	return &sampResultList, nil
}

func SampResultGetPageByPftestId(pftestId int64, page int, pageSize int) (*[]SampResult, int64) {
	o := orm.NewOrm()
	var sampResultList []SampResult
	qs := o.QueryTable(sampResultTableName())
	qs = qs.Filter("pftest_id", pftestId)
	total, _ := qs.Count()
	qs.OrderBy("-id").Limit(pageSize).Offset((page - 1) * pageSize).All(&sampResultList)
	return &sampResultList, total
}
