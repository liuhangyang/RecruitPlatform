package controllers

import (
	"github.com/astaxie/beego"
	"RecruitPlatform/models"
)
// Operations about News

type NewsController struct {
	beego.Controller
}

// @Title Getlist
// @Description get all news list
// @Success 200 {object}  models.News
// @Failure 400 get newlist error
// @router / [get]
func (n *NewsController) GetNewList(){
	list,err := models.GetNewList()
	if err != nil{
		n.Data["json"] = err.Error()
	}else{
		n.Data["json"] = list
	}
	n.ServeJSON()
}
