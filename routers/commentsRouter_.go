package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["RecruitPlatform/controllers:NewsController"] = append(beego.GlobalControllerRouter["RecruitPlatform/controllers:NewsController"],
		beego.ControllerComments{
			Method: "GetNewList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["RecruitPlatform/controllers:ReCommentController"] = append(beego.GlobalControllerRouter["RecruitPlatform/controllers:ReCommentController"],
		beego.ControllerComments{
			Method: "GetRecuitComment",
			Router: `/:Id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["RecruitPlatform/controllers:ReCommentController"] = append(beego.GlobalControllerRouter["RecruitPlatform/controllers:ReCommentController"],
		beego.ControllerComments{
			Method: "InsertRecruitComment",
			Router: `/:Id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["RecruitPlatform/controllers:RecruitController"] = append(beego.GlobalControllerRouter["RecruitPlatform/controllers:RecruitController"],
		beego.ControllerComments{
			Method: "GetRecruitMsgList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["RecruitPlatform/controllers:RecruitController"] = append(beego.GlobalControllerRouter["RecruitPlatform/controllers:RecruitController"],
		beego.ControllerComments{
			Method: "CommitRecruitNews",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["RecruitPlatform/controllers:RecruitController"] = append(beego.GlobalControllerRouter["RecruitPlatform/controllers:RecruitController"],
		beego.ControllerComments{
			Method: "GetRecruitDetail",
			Router: `/:Id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
