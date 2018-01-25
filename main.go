package main

import (
	"sappo/controllers"
	"sappo/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	orm.Debug = true                      //开启 ORM 调试模式
	orm.RunSyncdb("default", false, true) //自动建立新DB表

	//注册路由
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/category", &controllers.CategoryController{})
	//采购订单待审批
	beego.Router("/daisp", &controllers.DaispController{})
	beego.AutoRouter(&controllers.DaispController{})
	beego.Router("/sappo", &controllers.SappoController{})
	beego.Router("/price", &controllers.PriceController{})
	//采购申请待审批
	beego.Router("/daisppr", &controllers.DaispprController{})
	beego.AutoRouter(&controllers.DaispprController{})
	//user
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/loginadd", &controllers.LoginaddController{})
	beego.Router("/loginchange", &controllers.LoginchangeController{})

	beego.ErrorController(&controllers.ErrorController{})
	//beego.ErrorHandler("/404", &controllers.PageNotFound{})

	//启动beeblog
	beego.Run()
}
