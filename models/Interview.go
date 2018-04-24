package models

import (
	//"github.com/astaxie/beego/orm"
	"time"
)
type Interview struct{
	Id				int `orm:"pk;auto"`
	Person_id			int
	Person_name			string `orm:"size(30)"`
	Interview_title		string `orm:"size(150)"`
	Interview_company	string `orm:"size(100)"`
	Interview_time		time.Time `orm:"auto_now_add;type(datetime)"`
	Interview_way		string `orm:"size(100)"`
	Interview_class     string `orm:"size(100)"`
	Interview_content	string `orm:"type(text);null"`
	InterviewComment    []*InterviewComment `orm:"reverse(many)"`
}

