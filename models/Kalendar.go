package models

import (
	//"github.com/astaxie/beego/orm"
	"time"
)
type Kalendar struct {
	Id 				int `orm:"pk;auto"`
	Person_id			int
	Person_name			string `orm:"size(20)"`
	Event_title			string `orm:"size(50)"`
	Event_starttime		time.Time `orm:"auto_now_add;type(datetime)"`
	Event_endtime		time.Time `orm:"auto_now_add;type(datetime)"`
	Event_allday		int8
	Event_bgcolor		string `orm:"size(10)"`
	Is_avaliable		int8
	}
