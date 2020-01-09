package utils

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/httplib"
	"time"
)

var (
	urlcache cache.Cache
)

func init() {
	urlcache, _ = cache.NewCache("memory", `{"interval":20}`)
}

func GetAccessToken() (string, error) {
	accessToken := urlcache.Get("DINGTALK_ACCESS_TOKEN")
	if accessToken != nil {
		return cache.GetString(accessToken), nil
	}
	//请求接口 获取token
	appId := beego.AppConfig.String("dingtalk.appid")
	appSecret := beego.AppConfig.String("dingtalk.appsecret")
	req := httplib.Get("https://oapi.dingtalk.com/sns/gettoken?appid=" + appId + "&appsecret=" + appSecret)

	type Rst struct {
		Errcode     int    `json:"errcode"`
		Errmsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
	}
	rst := new(Rst)
	err := req.ToJSON(rst)
	if err != nil {
		return *new(string), err
	}
	if rst.Errcode != 0 {
		return *new(string), errors.New(rst.Errmsg)
	}
	urlcache.Put("DINGTALK_ACCESS_TOKEN", rst.AccessToken, 110*time.Minute)
	return rst.AccessToken, nil
}

func GetPersistentCode(accessToken, tmpAuthCode string) (*string, *string, error) {
	req := httplib.Post("https://oapi.dingtalk.com/sns/get_persistent_code?access_token=" + accessToken)
	pms := map[string]string{"tmp_auth_code": tmpAuthCode}
	req.JSONBody(pms)

	type Rst struct {
		Errcode        int    `json:"errcode"`
		Errmsg         string `json:"errmsg"`
		Openid         string `json:"openid"`
		PersistentCode string `json:"persistent_code"`
		Unionid        string `json:"unionid"`
	}

	rst := new(Rst)
	err := req.ToJSON(rst)
	if err != nil {
		return nil, nil, err
	}
	if rst.Errcode != 0 {
		return nil, nil, errors.New(rst.Errmsg)
	}
	return &rst.Openid, &rst.PersistentCode, nil
}

func GetSnsToken(openid, persistentCode *string, accessToken string) (string, error) {
	snsToken := urlcache.Get("DINGTALK_SNS_TOKEN_" + *openid)
	if snsToken != nil {
		return cache.GetString(snsToken), nil
	}

	type Rst struct {
		Errcode   int    `json:"errcode"`
		Errmsg    string `json:"errmsg"`
		ExpiresIn int    `json:"expires_in"`
		SnsToken  string `json:"sns_token"`
	}

	rst := new(Rst)
	req := httplib.Post("https://oapi.dingtalk.com/sns/get_sns_token?access_token=" + accessToken)
	pms := map[string]string{
		"openid":          *openid,
		"persistent_code": *persistentCode,
	}
	req.JSONBody(pms)
	err := req.ToJSON(rst)
	if err != nil {
		return *new(string), err
	}
	if rst.Errcode != 0 {
		return *new(string), errors.New(rst.Errmsg)
	}
	urlcache.Put("DINGTALK_SNS_TOKEN_"+*openid, rst.SnsToken, 110*time.Minute)
	return rst.SnsToken, nil
}

type DingTalkUser struct {
	Nick    string `json:"nick"`
	Openid  string `json:"openid"`
	Unionid string `json:"unionid"`
	DingId  string `json:"dingId"`
}

func GetUserInfo(snsToken string) (*DingTalkUser, error) {
	req := httplib.Get("https://oapi.dingtalk.com/sns/getuserinfo?sns_token=" + snsToken)

	type Rst struct {
		Errcode  int          `json:"errcode"`
		Errmsg   string       `json:"errmsg"`
		UserInfo DingTalkUser `json:"user_info"`
	}

	rst := new(Rst)
	err := req.ToJSON(rst)
	if err != nil {
		return nil, err
	}
	if rst.Errcode != 0 {
		return nil, errors.New(rst.Errmsg)
	}
	return &rst.UserInfo, nil
}
