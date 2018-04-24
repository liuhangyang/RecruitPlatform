package models
import (
	//"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego/orm"
)
type RecruitMsg struct{
	Id 					int `orm:"pk;auto"`
	Person_id			int
	Person_name			string `orm:"size(30)"`
	Recruit_title		string `orm:"size(100)"`
	Recruit_company		string `orm:"size(30)"`
	Recruit_releasetime time.Time `orm:"auto_now_add;type(datetime)"`
	Recruit_endtime		time.Time `orm:"auto_now_add;type(datetime)"`
	Recruit_content		string `orm:"type(text)"`
	RecruitComment		[]*RecruitComment `orm:"reverse(many)"`
}

func GetRecruitMsgList() (maps []orm.Params,err error){
	o := orm.NewOrm()
	num,err := o.Raw("SELECT id,person_name,recruit_title,recruit_releasetime,recruit_endtime from recruit_msg").Values(&maps)
	if err != nil || num == 0{
		return nil,err
	}

	return
}
func InsertRecruitList( recruit *RecruitMsg)(map[string]string,error){
	maps := make(map[string]string)
	o := orm.NewOrm()
	_,err := o.Insert(recruit)
	if err != nil{
		maps["result"] = "False"
		maps["message"] = err.Error()
		return  maps,err
		}
	maps["result"] = "True"
	maps["message"] = "success"
	return  maps,err
}
func GetRecruitDetailById(id int)(maps []orm.Params,err error ) {
	o := orm.NewOrm()
	num, err := o.Raw("SELECT person_name,recruit_title,recruit_releasetime,recruit_content  from recruit_msg where id = ?", id).Values(&maps)
	if err != nil || num == 0 {
		return nil, err
	}
	return
}
