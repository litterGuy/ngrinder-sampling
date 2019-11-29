package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//场景测试
type TestPms struct {
	Id         int64
	Name       string
	FileData   string
	ScriptPath string
	PftestId   int64
	CreateTime time.Time
	UpdateTime time.Time
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
