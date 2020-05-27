package routers

import (
	"github.com/astaxie/beego"
	"myfirst/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/demo0", &controllers.DemoController{})
	beego.Router("/demo1", &controllers.DemoController{}, "Get:Demo1Get")
	beego.Router("/Register", &controllers.RegisterController{})
	beego.Router("/Login", &controllers.RegisterController{}, "Get:Login;Post:LoginPost")
	beego.Router("/Index", &controllers.IndexController{})
	beego.Router("/AddArticle",&controllers.IndexController{},"Get:ShowAddArticle;Post:AddArticle")
	beego.Router("/Content",&controllers.ContentController{})
	beego.Router("/Update",&controllers.ContentController{},"Get:ShowUpdate;Post:Update")
	beego.Router("/delete",&controllers.IndexController{},"Get:HandleDelete")
	beego.Router("/AddType",&controllers.IndexController{},"Get:ShowAddType;Post:AddType")
	beego.Router("/deleteType",&controllers.IndexController{},"Get:DeleteType")
	beego.Router("/Logout",&controllers.IndexController{},"Get:Logout")
}
