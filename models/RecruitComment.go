package models
import (
	//"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego/orm"
)
type RecruitComment struct {
	Id 					int `orm:"pk;auto"`
	Person_id			int
	Person_name			string `orm:"size(30)"`
	Release_head_url	string `orm:"size(200)"`
	Comment_parent		int
	Comment_content     string `orm:"type(text)"`
	Comment_created		time.Time `orm:"auto_now_add;type(datetime)"`
	Comment_modified 	time.Time `orm:"auto_now_add;type(datetime);default('')"`
	UpvoteCount			int
	UserHasUpvoted		string `orm:"type(text)"`
	RecruitMsg			*RecruitMsg `orm:"rel(fk)"`
	}

func GetRecruitCommentById(id int,)(mapsjson []map[string]interface{},err error){
	var comment []*RecruitComment
	o := orm.NewOrm()
	num,err := o.QueryTable("recruit_comment").Filter("recruit_msg_id",id).RelatedSel().All(&comment)
	if err != nil || num == 0{
		return mapsjson,err
	}
	 n:=int(num)
	for i:= 0; i <n ; i++{
		maps := make(map[string]interface{})
		maps["id"] = comment[i].Id
		maps["parent"] = comment[i].Comment_parent
		maps["created"] =comment[i].Comment_created.Format("2006-01-02 15:04:05")
		maps["modified"] =comment[i].Comment_modified.Format("2006-01-02 15:04:05")
		maps["content"] =comment[i].Comment_content
		maps["fullname"] =comment[i].Person_name
		maps["profile_picture_url"] = comment[i].Release_head_url
		maps["created_by_current_user"] = "True"
		maps["upvote_count"]=comment[i].UpvoteCount
		maps["user_has_upvoted"]=comment[i].UserHasUpvoted
		mapsjson=append(mapsjson, maps)
		}
	return mapsjson,err
}
func InsertRecruitCommentById(comment * RecruitComment,id int)(error){
	o := orm.NewOrm()
	_,err:=o.Raw("insert recruit_comment(person_id,person_name,release_head_url,comment_parent,comment_created,recruit_msg_id,comment_content) values(?,?,?,?,?,?,?)",comment.Person_id,comment.Person_name,comment.Release_head_url,comment.Comment_parent,comment.Comment_created,id,comment.Comment_content).Exec()
	return  err
}
func UpdateRecruitCommentById(comment *RecruitComment)(error){
	o := orm.NewOrm()
	_,err:=o.Update(comment,"Comment_content","Comment_modified")
	return  err
}