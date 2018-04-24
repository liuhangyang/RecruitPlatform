package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"RecruitPlatform/models"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type ReCommentController struct {
	beego.Controller
}
// @Title GetRecruitComment
// @Description get  recruit detail
// @Param key path string true   "the id for recruitMsg for comment"
// @Success 200 {object}  models.RecruitComment
// @Failure 400 get recruit detail error
// @router /:Id [get]
func (n *ReCommentController) GetRecuitComment(){
	id, _ := n.GetInt(":Id")
	list, err := models.GetRecruitCommentById(id)

	fmt.Println(list)
	if err != nil {
		n.Data["json"] = err.Error()
	}else{
		n.Data["json"] = list
	}
	n.ServeJSON()
}
// @TitleInsertRecruitComment
// @Description add  comment
// @Param key path string true   "the id for recruitMsg for comment"
// @Success 200 {object}  models.RecruitComment
// @Failure 400 get recruit detail error
// @router /:Id [post]
func(n *ReCommentController) InsertRecruitComment(){
	maps := make(map[string]string)
	var dat map[string]interface{}
	id, _ := n.GetInt(":Id")
	fmt.Println("AddComment:", id)
	//maps := make(map[string]string)
	str := n.GetString("data")
	var comment models.RecruitComment
	if err := json.Unmarshal([]byte(str), &dat); err != nil {
		maps["result"]="False"
		maps["message"] ="数据格式错误"
	}
		//把TZ 时间转换为UTC
		created := dat["created"]
		str1 := created.(string)
		s := strings.Split(str1,".")[0]
		s1 := strings.Replace(s,"T"," ",1)
		t, _ := time.Parse("2006-01-02 15:04:05", s1)
		t =t.Add(time.Hour*8)
		// 把content转换为string
		content := dat["content"].(string)
		if parent := dat["parent"];parent == nil{
			comment.Comment_parent = 0
		}else{
			if sid,ok := parent.(string);ok{
				fmt.Println("SID",sid)
				id,_:= strconv.Atoi(sid)
				fmt.Println("ID:",id)
				comment.Comment_parent= id
				}else{
					//打印log
			}
		}
		comment.Person_id = user.Id
		comment.Person_name = user.Name
		comment.Release_head_url = user.Avatar_url
		comment.Comment_created = t
		comment.Comment_content = content
		err := models.InsertRecruitCommentById(&comment,id)
		if err == nil{
			maps["result"]="True"
			maps["message"] ="success"

			}else{
			maps["result"]="False"
			maps["message"] =err.Error()
		}
	n.Data["json"] = maps
	n.ServeJSON()
	}
func(n *ReCommentController) UpdateRecruitComment(){
	maps := make(map[string]string)
	var dat map[string]string
	str := n.GetString("data")
	if err := json.Unmarshal([]byte(str), &dat); err != nil {
		maps["result"]="False"
		maps["message"] ="数据格式错误"
	}
	content := dat["content"]
	fmt.Println("dat[id]:",dat["id"])
	id,_:= strconv.Atoi(dat["id"])
	fmt.Println("cID:",id)
	comment := models.RecruitComment{Id:id}
	comment.Comment_content = content
	comment.Comment_modified = time.Now()
	err := models.UpdateRecruitCommentById(&comment)
	if err == nil{
		maps["result"]="True"
		maps["message"] ="success"

	}else{
		maps["result"]="False"
		maps["message"] =err.Error()
	}
	n.Data["json"] = maps
	n.ServeJSON()
	}