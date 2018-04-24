package models

import (
	//"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id 					int `orm:"pk;auto"`
	Tag_name 			string `orm:"size(20)"`
	ResumeMsg 			[]*ResumeMsg `orm:"reverse(many)"`
	}
