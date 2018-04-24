package routers

import (
	"RecruitPlatform/controllers"
	"github.com/astaxie/beego"
	//"RecruitPlatform/filters"
	//"github.com/astaxie/beego/plugins/cors"
	//"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/plugins/cors"
)
// @APIVersion 1.0.0
// @Title XiYouLinux CS Recruit API
// @Description  Recruitment platform API of XiYouLinux CS.
// @Contact yanglongfei@xiyoulinux.org
func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))


    beego.Router("/", &controllers.IndexController{},"GET:Login")
    beego.Router("/index",&controllers.IndexController{},"GET:Index")
    beego.Router("/recruit/update_comment/:id",&controllers.ReCommentController{},"POST:UpdateRecruitComment")
    ns :=
    	beego.NewNamespace("/recruit",
    		beego.NSNamespace("/newslist",
					beego.NSInclude(
						&controllers.NewsController{},
					),
    			),
    				beego.NSNamespace("/get_recruit_msg",
    					beego.NSInclude(
    					&controllers.RecruitController{},
    					),
					),
						beego.NSNamespace("/commit_recruit_news",
							beego.NSInclude(
								&controllers.RecruitController{},
							),
							),
								beego.NSNamespace("/detail",
									beego.NSInclude(
										&controllers.RecruitController{},
									),
									),
										beego.NSNamespace("/comments",
											beego.NSInclude(
												&controllers.ReCommentController{},
											),
												),
													beego.NSNamespace("/post_comment",
														beego.NSInclude(
															&controllers.ReCommentController{},
														),
															),
    															)
    	beego.AddNamespace(ns)
    }

