package models
import (
	//"github.com/astaxie/beego/orm"
	"time"
)
type ResumeMsg struct{
	Id  				int `orm:"pk;auto"`
	Person_id			int
	Person_name			string `orm:"size(30)"`
	Resume_name			string `orm:"size(100)"`
	Resume_desc			string `orm:"size(100)"`
	Resume_path			string `orm:"size(100)"`
	Resume_thumb 		string `orm:"size(100)"`
	Upload_time			time.Time               `orm:"auto_now_add;type(datetime)"`
	Is_gathered  		int8
	Tag					[]*Tag `orm:"rel(m2m)"`
}

