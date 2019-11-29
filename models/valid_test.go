package models

import (
	"github.com/astaxie/beego/validation"
	"testing"
	"time"
)

func TestValid(t *testing.T) {
	var sences SencesRequestBean
	valid := validation.Validation{}
	b, err := valid.Valid(sences)
	if err != nil {
		print(err.Error())
	}
	if !b {
		// validation does not pass
		print(valid.Errors[0].Field + ":" + valid.Errors[0].Error())
	}
}

func TestTimes(t *testing.T) {
	times := 1574753897
	println(time.Unix(int64(times), 849e9).String())
	println(time.Now().String())
}
