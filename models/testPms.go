package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//场景测试
type TestPms struct {
	Id                      int64
	Name                    string
	FileData                string
	IgnoreSampleCount       int
	TargetHosts             string
	UseRampUp               string
	RampUpType              string
	Threshold               string
	Duration                int64
	RunCount                int
	AgentCount              int
	VuserPerAgent           int
	Processes               int
	RampUpInitCount         int
	RampUpInitSleepTime     int
	RampUpStep              int
	RampUpIncrementInterval int
	Threads                 int
	SamplingInterval        int
	Param                   string
	UserId                  string
	CreateTime              time.Time
	UpdateTime              time.Time
}

func init() {
	orm.RegisterModel(new(TestPms))
}

func TestPmsSave(t *TestPms, o orm.Ormer) (int64, error) {
	return o.Insert(t)
}

func TestPmsUpdate(t *TestPms, o orm.Ormer) error {
	if _, err := o.Update(t); err != nil {
		return err
	}
	return nil
}

func TestPmsDelete(t *TestPms, o orm.Ormer) (int64, error) {
	return o.Delete(t)
}

func TestPmsGetById(id int64) (*TestPms, error) {
	testPms := TestPms{Id: id}
	err := orm.NewOrm().Read(&testPms)
	if err != nil {
		return nil, err
	}
	return &testPms, nil
}

func TestPmsGetPageByUserId(userId string, name string, page int, pageSize int) (*[]TestPms, int64) {
	o := orm.NewOrm()
	var testPmsList []TestPms
	qs := o.QueryTable(testPmsTableName())
	qs = qs.Filter("user_id", userId)
	if len(name) > 0 {
		qs = qs.Filter("name__contains", name)
	}
	total, _ := qs.Count()
	qs.OrderBy("-id").Limit(pageSize).Offset((page - 1) * pageSize).All(&testPmsList)
	return &testPmsList, total
}

func testPmsTableName() string {
	return "test_pms"
}
