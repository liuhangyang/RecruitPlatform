package main

import (
	_ "RecruitPlatform/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"RecruitPlatform/models"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	//"github.com/astaxie/beego/session"
	//"github.com/astaxie/beego/session"
)
func init(){
	beego.BConfig.WebConfig.Session.SessionOn = true
	orm.RegisterDataBase("default","mysql",beego.AppConfig.String("jdbc.username")+":"+beego.AppConfig.String("jdbc.password")+"@/recruit?charset=utf8",30)
	orm.RegisterModel(
		new(models.Interview),
		new(models.Tag),
		new(models.InterviewComment),
		new(models.RecruitMsg),
		new(models.RecruitComment),
		new(models.ResumeMsg),
		new(models.News))
		orm.RunSyncdb("default",false,true)
	   beego.AddTemplateExt("js")



}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}

