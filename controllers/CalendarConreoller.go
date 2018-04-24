package controllers

import (
	"github.com/astaxie/beego"
)

type CalendarContorller struct{
	beego.Controller
}
func ( c *CalendarContorller) CalendarIndex(){
	c.Layout = "index.html"
	c.TplName = "news.html"
}


