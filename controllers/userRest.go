package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type UserController struct {
	BaseController
}

// @router	/login	[post]
func (u *UserController) Login() {
	userId := u.GetString("userId")
	password := u.GetString("password")
	if len(userId) <= 0 || len(password) <= 0 {
		u.result.Code = 1
		u.result.ErrMsg = "userId or password is not empty"
		u.responseAjax()
	}
	ngrinderUrl := beego.AppConfig.String("ngrinder.serverurl")
	apiUrl := beego.AppConfig.String("ngrinder.api.login")
	req := httplib.Post(ngrinderUrl + apiUrl)
	req.Param("userId", userId)
	req.Param("password", password)
	var js NsResponseBean
	rst, err := req.String()
	if err != nil {
		u.result.Code = 1
		u.result.ErrMsg = err.Error()
	} else {
		err = json.Unmarshal([]byte(rst), &js)
		if err != nil {
			u.result.Code = 1
			u.result.ErrMsg = rst
		} else {
			u.result.Code = js.Code
			if js.Code == 0 {
				u.SetSession(SESSION_NAME, userId)
			} else {
				u.result.ErrMsg = js.ErrMsg
			}
		}
	}
	u.responseAjax()
}

// @router	/logout	[get]
func (u *UserController) Logout() {
	u.DestroySession()
	u.Redirect("/", 302)
}
