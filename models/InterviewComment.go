package models

import (
	//"github.com/astaxie/beego/orm"
	"time"

)
type InterviewComment struct{
	Id 					int `orm:"pk;auto"`
	Person_id			int
	Person_name			string `orm:"size(30)"`
	Release_head_url	string `orm:"size(100)"`
	Comment_parent		int
	Comment_created		time.Time `orm:"auto_now_add;type(datetime)"`
	Comment_modified 	time.Time `orm:"auto_now_add;type(datetime)"`
	UpvoteCount			int
	UserHasUpvoted		string `orm:"type(text)"`
	Interview 			*Interview `orm:"rel(fk)"`
	}