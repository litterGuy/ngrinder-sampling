package controllers

import (
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
	js := new(NsResponseBean)
	err := req.ToJSON(js)
	if err != nil {
		u.result.Code = 1
		u.result.ErrMsg = err.Error()
	} else {
		u.result.Code = js.Code
		if js.Code == 0 {
			u.SetSession(SESSION_USER_ID, userId)
			nuser, ok := js.Data.(NUser)
			if ok {
				u.SetSession(SESSION_USER_NICK, nuser.UserName)
			} else {
				u.SetSession(SESSION_USER_NICK, userId)
			}
		} else {
			u.result.ErrMsg = js.ErrMsg
		}
	}
	u.responseAjax()
}

// @router	/logout	[get]
func (u *UserController) Logout() {
	u.DestroySession()
	u.Redirect("/", 302)
}
