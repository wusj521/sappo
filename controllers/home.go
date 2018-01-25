package controllers

import (
	"sappo/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	//登录检查
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	c.Data["IsShouy"] = true
	c.TplName = "home.html"

	c.Data["IsLogin"] = checkAccount(c.Ctx)

	//c.Data["Prgco_X"], c.Data["Prgco_not"],
	//rel_list := [2]int
//读取采购订单行数量
	// 解析表单-从表单字段中获取内容
	prg := c.Ctx.GetCookie("__prg")//PO审批码
	flagx, flagnot, ebelncont := models.GetPrgcocount(prg)
	//rel_list1 := [10]int{50, 30, 20}
	c.Data["Flagx"] = flagx
	c.Data["Flagnot"] = flagnot
	c.Data["Ebelncont"] = ebelncont
//读取采购申请行数量
	prgpr := c.Ctx.GetCookie("__prgpr")//PR审批码
	flagxpr, flagnotpr := models.GetPrgcocountpr(prgpr)
	c.Data["Flagxpr"] = flagxpr
	c.Data["Flagnotpr"] = flagnotpr


	//读取物料可用天数
	flag := "" //未审批标记
	Maktxs, Zgys06901s, err := models.GetMatnrkday(flag, prg)
	if err != nil {
		beego.Error(err)
		c.Redirect("/daisp", 302)
		return
	}

	c.Data["Maktxs"] = Maktxs         //工厂加物料描述
	c.Data["Zgys06901s"] = Zgys06901s //库存可用天数

	/*	c.Data["Website"] = "beego.me"
		c.Data["Email"] = "astaxie@gmail.com"
		c.TplName = "index.tpl"

		c.Data["TrueCond"] = true
		c.Data["FalseCond"] = false

		type u struct {
			Name string
			Age  int
			Sex  string
		}
		User := &u{
			Name: "wusj",
			Age:  36,
			Sex:  "Male",
		}
		c.Data["User"] = User

		numb := []int{1, 2, 3, 4, 5, 6}
		c.Data["Numb"] = numb
	*/

}
