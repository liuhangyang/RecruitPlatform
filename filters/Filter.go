package filters

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"fmt"
)
var IsLogin = func (ctx *context.Context)(){
	token, flag := ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
	fmt.Println("token:"+ token)
	fmt.Println(flag)

}