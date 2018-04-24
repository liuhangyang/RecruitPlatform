package controllers

import (
	"github.com/astaxie/beego"
	"RecruitPlatform/models"
	 "time"
	"fmt"
)

// Operations about RecruitMsg
type RecruitController struct {
	beego.Controller
}
// @Title GetRecruitList
// @Description get all recruit list
// @Success 200 {object}  models.RecruitMsg
// @Failure 400 get recruitmsg error
// @router / [get]
func (n *RecruitController) GetRecruitMsgList(){
	list,err := models.GetRecruitMsgList()
	urlMap := make(map[string]interface{})
	urlMap["data"]= list
	if err != nil{
		n.Data["json"] = err.Error()
	}else{
		n.Data["json"] = urlMap
	}
	n.ServeJSON()
}

// @Title CommitRecruitMsg
// @Description commit a recruit msg
// @Success 200 {object}  models.RecruitMsg
// @Failure 400 commit recruitmsg error
// @router / [post]
func (n *RecruitController) CommitRecruitNews(){
	title := n.Input().Get("title")
	company := n .Input().Get("company")
	endtime := n.Input().Get("endtime")
	endtime += " 00:00:00"
	content := n.Input().Get("detail")
	Releasetime := time.Now()


	t, _ := time.Parse("2006-01-02 15:04:05", endtime)
	recruit := models.RecruitMsg{Person_id:user.Id,Person_name:user.Name,Recruit_title:title,Recruit_company:company,Recruit_releasetime:Releasetime,Recruit_endtime:t,Recruit_content:content}
	maps,err:= models.InsertRecruitList(&recruit)
	if err ==nil{
		s := fmt.Sprintf("%s发布了一条新招聘动态",user.Name)
		new := models.News{Person_id:user.Id,Person_name:user.Name,News_title:s,News_time:Releasetime,News_url:"http://127.0.0.1/recruit/1"}
		err = models.InsertNewList(&new)
		if err != nil{

		}else{

		}
	}
	n.Data["json"] = maps
	n.ServeJSON()
	}
// @Title GetRecruitList
// @Description get  recruit detail
// @Param key path string true   "the id for recruitmsg"
// @Success 200 {object}  models.RecruitMsg
// @Failure 400 get recruit detail error
// @router /:Id [get]
func (n *RecruitController) GetRecruitDetail() {
	id, _ := n.GetInt(":Id")
	fmt.Println("ID:", id)
	list, err := models.GetRecruitDetailById(id)
	if err == nil{
		n.Data["recruit_title"] = list[0]["recruit_title"]
		n.Data["recruit_releasetime"] = list[0]["recruit_releasetime"]
		n.Data["recruit_content"] = list[0]["recruit_content"]
		n.Data["person_name"] = list[0]["recruit_title"]
	}
	n.Data["avatar"] = user.Avatar_url

	n.Data["user"] = user.Name
	n.Layout = "index.html"
	n.TplName = "recruit_detail.html"
}
