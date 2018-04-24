package models


import (
	//"github.com/astaxie/beego/orm"
	"time"
	//"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/orm"

)
type News struct{
	Id 					int `orm:"pk;auto`
	Person_id			int
	Person_name			string `orm:"size(20)"`
	News_title			string `orm:"size(100)"`
	News_time			time.Time `orm:"auto_now_add;type(datetime)"`
	News_url			string `orm:"size(200)"`
}

func GetNewList()(maps []orm.Params,err error){
	o := orm.NewOrm()
	num,err := o.Raw("SELECT news_title,news_time,news_url from news").Values(&maps)
	if err != nil || num == 0{
		return nil,err
	}
	return

}
func  InsertNewList(new *News)(error){
		o := orm.NewOrm()
		_,err := o.Insert(new)
		return err
		}
