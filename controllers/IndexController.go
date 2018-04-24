package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"fmt"
)

var  flag = 0
var resp OauthToken
var user User
type IndexController struct {
	beego.Controller
}
type OauthToken struct{
	Token 	string `json:"access_token"`
	Type    string `json:"token_type"`
	ExpiresIn string `json:"expires_in"`
}
type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Avatar_url string `json:"avatar_url"`
}

func (c *IndexController) Login(){
	flag = 1
	urlMap := make(map[string]string)
	urlMap["response_type"]= beego.AppConfig.String("Response_type")
	urlMap["client_id"] = beego.AppConfig.String("ClientId")
	urlMap["state"] = beego.AppConfig.String("State")
	urlMap["redirect_uri"]=beego.AppConfig.String("Redirect_uri")
	urlMap["scope"]="all"

	var url = beego.AppConfig.String("GetCodeURL") + "?"
	for key,value := range urlMap{
		url += key +"="+value+ "&"
	}
	url =url[:len(url)-1]
	c.Ctx.Redirect(302,url)
		}
func (c *IndexController) Index() {
	//	fmt.Println(c.GetString("code"))
	if flag != 0 {
		req := httplib.Post(beego.AppConfig.String("GetTokenURL"))
		req.Header("Content-Type", "application/x-www-form-urlencoded")
		req.Param("code", c.GetString("code"))
		fmt.Println("code:", c.GetString("code"))
		req.Param("grant_type", "authorization_code")
		req.Param("client_id", beego.AppConfig.String("ClientId"))
		req.Param("client_secret", beego.AppConfig.String("ClientSecret"))
		req.Param("redirect_uri", beego.AppConfig.String("Redirect_uri"))

		req.ToJSON(&resp)
		fmt.Println("token:" + resp.Token + " " + "type:" + resp.Type + " " + "expires_in")
		flag = 0
	}
	req_info := httplib.Get(beego.AppConfig.String("GetUserInfoURL") + resp.Token)
	fmt.Println(req_info.String())
	req_info.ToJSON(&user)

	c.Data["avatar"] = user.Avatar_url

	c.Data["user"] = user.Name
	c.Layout = "index.html"
	c.TplName = "news.html"
}
