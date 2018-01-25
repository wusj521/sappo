package controllers

import (
	"sappo/models"

	"github.com/astaxie/beego"
)

type DaispprController struct {
	beego.Controller
}

func (c *DaispprController) Get() {
	//登录检查
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsDaisppr"] = true
	c.TplName = "daisppr.html"

	var err error
	// 解析表单-从表单字段中获取内容
	prgpr := c.Ctx.GetCookie("__prgpr")
	c.Data["Sappr"], err = models.GetAllSapprs(prgpr)
	if err != nil {
		beego.Error(err)
	}

	c.Data["Tid"] = c.Ctx.Input.Param("0")

}

//审批PO
func (c *DaispprController) Post() {
	//登录检查
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	// 解析表单-从表单字段中获取内容
	tid := c.Input().Get("tid")
	banfn := c.Input().Get("banfn") //申请单号
	prgpr := c.Ctx.GetCookie("__prgpr")
	flag := "X" //c.Input().Get("flag")

	var err error
	//写入sappo表中flag审批标记为X
	err = models.ModifyDaisppr(tid, banfn, flag)
	if err != nil {
		beego.Error(err)
	}
	//调取审批RFC同时写入sappo中上传SAP标记
	// 解析表单-从表单字段中获取内容
	tid = c.Input().Get("tid")
	banfn = c.Input().Get("banfn")
	prgpr = c.Ctx.GetCookie("__prgpr")
	uppo := "X"
	//审批前检查sappo本地表是否还有未审批的PO行项目，因为sappo表是以行项目存取的。
	ok := models.Checkpr(banfn, prgpr)
	if ok == true {
		//没有查到同po号的行项目，说明是最后一个审批PO行项目，故调取SAP RFC审批。
		err = models.PostSapPr(tid, banfn, prgpr, uppo)
		if err != nil {
			beego.Error(err)
		}
	}

	c.Redirect("/daisppr", 302)
}

//列出将要审批的PO行项目内容
func (c *DaispprController) View() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsDaisppr"] = true
	c.TplName = "daisppr_view.html"

	tid := c.Input().Get("tid")
	Sappr, err := models.GetDaisppr(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/daisppr", 302)
		return
	}

	c.Data["sappr"] = Sappr
	c.Data["Tid"] = tid
	f := Sappr.Flag
	c.Data["Flag"] = checkFlag(f)
	//SAP Po,在查看html时，代码中加入以下隐藏字段，方便此页面执行POST或其它动作。
	////func (c *DaispController) Post()中使用
	//c.Data["Ebeln"] = sappo.Ebeln

	////在daisp的view明细页获取表单中的物料号matnr并存入cookie;
	////在price图表中使用
	///maxAge := 0
	////matnr := c.Input().Get("matnr")
	//matnr := sappo.Matnr
	////c.Ctx.SetCookie("__matnr", sappo.Matnr, maxAge, "/")

}

/*
func checkFlag(f string) bool {
	if f == "X" {
		return false
	}
	return true

}*/
